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

	fmt.Printf("Initializing ring queue\n")
	queue, err := NewQueue(QUEUE_SIZE)
	if err != nil {
		return fmt.Errorf("failed to create queue: %w", err)
	}
	defer queue.Close()

	fmt.Printf("Creating an SQE for each chunk\n")
	offset := uint(0)
	chunks := []*C.char{}
	defer func() {
		for _, chunk := range chunks {
			C.free(unsafe.Pointer(chunk))
		}
	}()
	for i := uint(0); i < nChunks; i++ {
		chunk, err := queue.Enqueue(Request{File: file, Size: CHUNK_BYTE_SIZE})
		if err != nil {
			return fmt.Errorf("failed to enqueue entry: %w", err)
		}
		chunks = append(chunks, chunk)
		offset += CHUNK_BYTE_SIZE
	}

	/*
			printf("Marking ring queue as read for processing\n");
		  int submitted_requests = io_uring_submit(&ring);
		  if (status < 0) {
		    fprintf(stderr, "io_uring_submit: %s\n", strerror(-submitted_requests));
		    exit_code = 1;
		    goto CLEANUP;
		  } else if (submitted_requests != n_chunks) {
		    fprintf(
		      stderr,
		      "io_uring_submit: created %d requests but %d were submitted\n",
		      n_chunks, submitted_requests
		    );
		    exit_code = 1;
		    goto CLEANUP;
		  }
	*/

	return nil
}

type Queue struct {
	Capacity uint
	Size     uint
	Ring     C.struct_io_uring
}

func NewQueue(capacity uint) (Queue, error) {
	q := Queue{Capacity: capacity}
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
	File   *os.File
	Offset uint
	Size   uint
}

func (q *Queue) Enqueue(req Request) (chunk *C.char, err error) {
	if q.Size == q.Capacity {
		return nil, errors.New("reached queue max capacity")
	}

	// Get the next available submission queue entry
	sqe := C.io_uring_get_sqe(&q.Ring)
	if sqe == nil {
		return nil, errors.New("reached queue max capacity")
	}

	chunk = (*C.char)(C.malloc(C.ulong(C.sizeof_char * req.Size)))
	defer func() {
		if err != nil {
			C.free(unsafe.Pointer(chunk))
		}
	}()

	C.io_uring_prep_read(sqe, C.int(req.File.Fd()), unsafe.Pointer(chunk), C.uint(req.Size), C.ulonglong(req.Offset))
	q.Size++

	return (chunk), nil
}
