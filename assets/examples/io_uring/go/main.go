package main

import (
	"fmt"
	"os"
)

func main() {

}

func ReadFile(path string) (chan []byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file at %s: %w", path, err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	// submitChan <- &request{
	// 	code:   opCodeRead, // a constant to identify which syscall we are going to make
	// 	f:      f,          // the file descriptor
	// 	size:   fi.Size(),  // size of the file
	// 	readCb: cb,         // the callback to call when the read is done
	// }
	return nil, nil
}
