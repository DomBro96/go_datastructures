package list

import (
	"fmt"
	"testing"
)

func TestCache_InsertCacheNode(t *testing.T) {
	cache := OpenCache()
	cache.InsertCacheNode(1)
	cache.InsertCacheNode(2)
	cache.InsertCacheNode(3)
	cache.InsertCacheNode(4)
	cache.InsertCacheNode(5)
	pNode := cache.head
	for pNode != nil {
		fmt.Println(pNode.cacheVal)
		pNode = pNode.next
	}
	cache.InsertCacheNode(1)
	fmt.Println()
	pNode = cache.head
	for pNode != nil {
		fmt.Println(pNode.cacheVal)
		pNode = pNode.next
	}
}
