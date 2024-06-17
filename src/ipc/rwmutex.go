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

type ReaderWriterMutexWithoutStarvation struct {
	readers      int
	writersQueue int
	cond         *sync.Cond
	writing      bool
}

func NewReaderWriterMutexWithoutStarvation() *ReaderWriterMutexWithoutStarvation {
	return &ReaderWriterMutexWithoutStarvation{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (rw *ReaderWriterMutexWithoutStarvation) ReadLock() {
	rw.cond.L.Lock()
	for rw.writing || rw.writersQueue > 0 {
		rw.cond.Wait()
	}

	rw.readers++
	rw.cond.L.Unlock()
}

func (rw *ReaderWriterMutexWithoutStarvation) ReadUnlock() {
	rw.cond.L.Lock()
	rw.readers--

	if rw.readers == 0 {
		rw.cond.Broadcast()
	}

	rw.cond.L.Unlock()
}

func (rw *ReaderWriterMutexWithoutStarvation) WriteLock() {
	rw.cond.L.Lock()

	rw.writersQueue++
	for rw.writing || rw.readers > 0 {
		rw.cond.Wait()
	}

	rw.writersQueue--
	rw.writing = true
	rw.cond.L.Unlock()
}

func (rw *ReaderWriterMutexWithoutStarvation) WriteUnlock() {
	rw.cond.L.Lock()
	rw.writing = false
	rw.cond.Broadcast()
	rw.cond.L.Unlock()
}
