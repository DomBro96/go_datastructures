package bitmap

import "testing"

func TestSet(t *testing.T) {
	b := NewBitMap(1024)
	b.Set(1023)
	t.Log(b.Get(1023))

	b.Set(1)
	t.Log(b.Get(1))

	t.Log(b.Get(2))

	b.Set(100)
	t.Log(b.Get(100))

	b.Set(1026)
	t.Log(b.Get(1026))
}

func TestUnSet(t *testing.T) {
	b := NewBitMap(1024)
	b.Set(1023)
	t.Log(b.Get(1023))

	b.Set(1)
	t.Log(b.Get(1))

	b.Set(100)
	t.Log(b.Get(100))

	b.UnSet(1023)
	t.Log(b.Get(1023))

	b.UnSet(1)
	t.Log(b.Get(1))

	b.UnSet(1023)
	t.Log(b.Get(1023))

}
