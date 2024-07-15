package internal

type SubPageAllocator struct {
	PageSize    int
	SubPageSize int
	PageCount   int
	pages       [][]bool
}

func NewSubPageAllocator(pageSize, subPageSize, pageCount int) *SubPageAllocator {
	pages := make([][]bool, pageCount)
	subPages := pageSize / subPageSize
	for i := range pages {
		pages[i] = make([]bool, subPages)
	}
	return &SubPageAllocator{
		PageSize:    pageSize,
		SubPageSize: subPageSize,
		PageCount:   pageCount,
		pages:       pages,
	}
}

func (a *SubPageAllocator) Allocate(size int) (pageIndex int, startIndex int, ok bool) {
	subPages := (size + a.SubPageSize - 1) / a.SubPageSize

	for pIdx, page := range a.pages {
		freeCount := 0

		for spIdx, subPage := range page {
			if subPage {
				freeCount = 0
				continue
			}

			freeCount++
			if freeCount == subPages {
				for i := 0; i < subPages; i++ {
					a.pages[pIdx][spIdx-i] = true
				}
				return pIdx, spIdx - subPages + 1, true
			}
		}
	}

	return -1, -1, false
}

func (a *SubPageAllocator) Free(pageIndex, startIndex, size int) {
	subPages := (size + a.SubPageSize - 1) / a.SubPageSize
	for i := 0; i < subPages; i++ {
		a.pages[pageIndex][startIndex+i] = false
	}
}
