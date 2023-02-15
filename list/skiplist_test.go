package list

import "testing"

func TestSkipList_Insert(t *testing.T) {
	skl := newSkipList(16)
	skl.insert(1)
	skl.insert(2)
	skl.insert(3)
	skl.insert(4)
	skl.insert(5)
}

func TestSkipList_Find(t *testing.T) {
	skl := newSkipList(16)
	skl.insert(1)
	skl.insert(2)
	skl.insert(3)
	skl.insert(4)
	skl.insert(5)
	skl.insert(100)
	skl.insert(500)
	skl.insert(32)
	skl.insert(9)
	skl.insert(10)
	t.Logf("find 3: %v", skl.find(3))
	t.Logf("find 10: %v", skl.find(10))
	t.Logf("find 100: %v", skl.find(100))
	t.Logf("skl max level: %d", skl.maxLevel)
}

func TestSkipList_Delete(t *testing.T) {
	skl := newSkipList(16)
	skl.insert(1)
	skl.insert(2)
	skl.insert(3)
	skl.insert(4)
	skl.insert(5)
	skl.insert(100)
	skl.insert(500)
	skl.insert(32)
	skl.insert(9)
	skl.insert(10)
	t.Logf("find 3: %v", skl.find(3))
	t.Logf("find 10: %v", skl.find(10))
	t.Logf("find 100: %v", skl.find(100))
	skl.delete(3)
	skl.delete(10)
	skl.delete(100)
	t.Logf("find 3: %v", skl.find(3))
	t.Logf("find 10: %v", skl.find(10))
	t.Logf("find 100: %v", skl.find(100))
}
