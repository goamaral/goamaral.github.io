package main

/*
#cgo pkg-config: liburing
#include <bridge.c>
*/
import "C"

import (
	"errors"
	"fmt"
	"math"
	"os"
	"unsafe"
)

const QUEUE_SIZE = 10
const CHUNK_BYTE_SIZE = 1024 // 1KB
const PATH = "example.txt"

func main() {

}

func run() error {
	path := "example.txt"

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file at %s: %w", path, err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}
	size := fileInfo.Size()
	nChunks := uint(math.Ceil(float64(size) / CHUNK_BYTE_SIZE))
	fmt.Printf("File info (size: %d, chunk_size: %d, n_chunks: %d)\n", size, CHUNK_BYTE_SIZE, nChunks)

	fmt.Println("Initializing ring queue")
	queue, err := NewQueue(QUEUE_SIZE)
	if err != nil {
		return fmt.Errorf("failed to create queue: %w", err)
	}
	defer queue.Close()

	fmt.Println("Enqueue a SQE for each chunk")
	offset := uint(0)
	chunks := []*C.char{}
	defer func() {
		for _, chunk := range chunks {
			C.free(unsafe.Pointer(chunk))
		}
	}()
	for i := uint(0); i < nChunks; i++ {
		chunk, err := queue.Enqueue(Request{Id: C.ulonglong(i), File: file, Size: CHUNK_BYTE_SIZE})
		if err != nil {
			return fmt.Errorf("failed to enqueue entry: %w", err)
		}
		chunks = append(chunks, chunk)
		offset += CHUNK_BYTE_SIZE
	}

	fmt.Println("Submitting queue")
	err = queue.Submit()
	if err != nil {
		return fmt.Errorf("failed to submit queue: %w", err)
	}

	fmt.Println("Wait for CQEs")
	for i := uint(0); i < nChunks; i++ {
		res, err := queue.WaitForResponse()
		if err != nil {
			return fmt.Errorf("failed to wait for response: %w", err)
		}
		if res.Err != nil {
			fmt.Fprintf(os.Stderr, "failed to read chunk (%d): %s", res.Id, res.Err)
			continue
		}
		fmt.Printf("---------- CHUNK %d ----------\n", res.Id)
		fmt.Println(C.GoString(chunks[res.Id]))
	}

	return nil
}

type Queue struct {
	Ring C.struct_io_uring
}

var ErrReachedMaxCapacity = errors.New("reached max capacity")

func NewQueue(capacity uint) (Queue, error) {
	q := Queue{}
	status := C.io_uring_queue_init(C.uint(capacity), &q.Ring, 0)
	if status < 0 {
		return Queue{}, errors.New(C.GoString(C.strerror(-status)))
	}
	return q, nil
}
func (q *Queue) Close() {
	C.io_uring_queue_exit(&q.Ring)
}

type Request struct {
	Id     C.ulonglong
	File   *os.File
	Offset uint
	Size   uint
}

func (q *Queue) Enqueue(req Request) (chunk *C.char, err error) {
	sqe := C.io_uring_get_sqe(&q.Ring)
	if sqe == nil {
		return nil, ErrReachedMaxCapacity
	}
	sqe.user_data = req.Id

	chunk = (*C.char)(C.malloc(C.ulong(C.sizeof_char * req.Size)))
	defer func() {
		if err != nil {
			C.free(unsafe.Pointer(chunk))
		}
	}()

	C.io_uring_prep_read(sqe, C.int(req.File.Fd()), unsafe.Pointer(chunk), C.uint(req.Size), C.ulonglong(req.Offset))

	return chunk, nil
}

func (q *Queue) Submit() error {
	submittedRequests := C.io_uring_submit(&q.Ring)
	if submittedRequests < 0 {
		return errors.New(C.GoString(C.strerror(-submittedRequests)))
	}
	return nil
}

type Response struct {
	Id  C.ulonglong
	Err error
}

func (q *Queue) WaitForResponse() (Response, error) {
	var cqe *C.struct_io_uring_cqe
	status := C.io_uring_wait_cqe(&q.Ring, &cqe)
	if status < 0 {
		return Response{}, errors.New(C.GoString(C.strerror(-status)))
	}
	res := Response{Id: cqe.user_data}
	if cqe.res < 0 {
		res.Err = errors.New(C.GoString(C.strerror(-cqe.res)))
	}
	C.io_uring_cqe_seen(&q.Ring, cqe)
	return res, nil
}
