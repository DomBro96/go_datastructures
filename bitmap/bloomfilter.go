package bitmap

// BloomFilter is a simple bloom filter implementation
type BloomFilter struct {
	size   uint
	seeds  []int
	bitmap *BitMap
}

func NewBloomFilter(size uint) *BloomFilter {
	sl := size * 69 / 100 // 0.69 =~ ln(2)

	if sl < 1 {
		sl = 1
	}

	if sl > 30 {
		sl = 30
	}

	primes := sievePrime(200)[:sl]

	return &BloomFilter{
		size:   size,
		seeds:  primes,
		bitmap: NewBitMap(size),
	}
}

func (b *BloomFilter) Add(data []byte) {
	for _, s := range b.seeds {
		h := b.hash(data, s)
		b.bitmap.Set(h)
	}
}

func (b *BloomFilter) Has(data []byte) bool {
	has := true
	for _, s := range b.seeds {
		h := b.hash(data, s)
		has := b.bitmap.Get(h)
		if !has {
			break
		}
	}

	return has
}

func (b *BloomFilter) hash(data []byte, seed int) uint {
	r := uint(0)
	for _, b := range data {
		r = uint(seed)*r + uint(b)
	}

	return (b.size - 1) | r
}

// sievePrime sieve prime numbers within n

func sievePrime(n int) []int {
	if n <= 0 {
		return nil
	}

	res := make([]int, 0)

	isPrime := make([]bool, n+1)
	for i := 2; i < len(isPrime); i++ {
		isPrime[i] = true
	}

	for i := 2; i <= n; i++ {
		if !isPrime[i] {
			continue
		}

		res = append(res, i)

		for j := 2 * i; j <= n; j += i {
			isPrime[j] = false
		}

	}

	return res
}
