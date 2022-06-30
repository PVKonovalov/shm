package shm

// #include "shm.h"
import "C"

import (
	"fmt"
	"unsafe"
)

type Segment struct {
	Id     int
	Size   int64
	offset int64
}

// OpenSegment open existed segment in shared memory
func OpenSegment() (*Segment, error) {
	if sharedMemoryId, err := C.sysv_shm_open(); err == nil {
		if actualSize, err := C.sysv_shm_get_size(sharedMemoryId); err != nil {
			return nil, fmt.Errorf("failed to retrieve segment size: %v", err)
		} else {
			return &Segment{
				Id:   int(sharedMemoryId),
				Size: int64(actualSize),
			}, nil
		}

	} else {
		return nil, err
	}
}

// Attach segment and return its unsafe pointer
func (s *Segment) Attach() (unsafe.Pointer, error) {
	if addr, err := C.sysv_shm_attach(C.int(s.Id)); err == nil {
		return unsafe.Pointer(addr), nil
	} else {
		return nil, err
	}
}
