package ipc

import "sync"

type ReaderWriterMutex struct {
	readers    int
	readMutex  *sync.Mutex
	writeMutex *sync.Mutex
}

func NewReaderWriterMutex() *ReaderWriterMutex {
	return &ReaderWriterMutex{
		readMutex:  &sync.Mutex{},
		writeMutex: &sync.Mutex{},
	}
}

func (rw *ReaderWriterMutex) ReadLock() {
	rw.readMutex.Lock()
	rw.readers++

	if rw.readers == 1 {
		rw.writeMutex.Lock()
	}

	rw.readMutex.Unlock()
}

func (rw *ReaderWriterMutex) ReadUnlock() {
	rw.readMutex.Lock()
	rw.readers--

	if rw.readers == 0 {
		rw.writeMutex.Unlock()
	}

	rw.readMutex.Unlock()
}

func (rw *ReaderWriterMutex) WriteLock() {
	rw.writeMutex.Lock()
}

func (rw *ReaderWriterMutex) WriteUnlock() {
	rw.writeMutex.Unlock()
}
