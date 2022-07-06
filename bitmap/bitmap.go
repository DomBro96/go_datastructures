package bitmap

type BitMap struct {
	size uint
	bits []uint64
}

func NewBitMap(size uint) *BitMap {
	return &BitMap{
		size: size,
		bits: make([]uint64, size/64+1),
	}
}

func (bm *BitMap) Set(n uint) {
	if n > bm.size {
		return
	}

	bm.bits[n/bm.size] |= (1 << (n % 64))
}

func (bm *BitMap) UnSet(n uint) {
	if n > bm.size {
		return
	}

	bm.bits[n/bm.size] &= (^(1 << (n % 64)))
}

func (bm *BitMap) Get(n uint) bool {
	if n > bm.size {
		return false
	}

	return bm.bits[n/bm.size]&(1<<(n%64)) != 0
}
