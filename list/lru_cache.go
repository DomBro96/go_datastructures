package list

import "unsafe"

// 利用链表，实现缓存中的 LRU(最近最少访问算法) 淘汰机制
// 链表实现 LRU 算法思路:
// 1. 查找某个数据时，从头遍历单缓存单链表，如果单链表中查询到了该数据，则从删除该位置节点，
//	  将其插入到缓存单链表的头部；
// 2. 如果单链表中没有该数据，若缓存未满，则将该数据插入到缓存链表头部。若满了则将链表中的
// 	  最后一个节点删除，将新数据节点插入到缓存单链表头部;

const (
	CACHE_NODE_SIZE = int(unsafe.Sizeof(CacheLinkedNode{}))	// 缓存节点大小
	MAX_CACHE_SIZE = 1024*CACHE_NODE_SIZE	// 默认的缓存大小
)

type(
	Cache struct {
		head *CacheLinkedNode	// 首节点
		size int				// 缓存大小
	}

	CacheLinkedNode struct {
		cacheVal int			// 缓存数据
		next *CacheLinkedNode	// next 指针
	}
)
// 初始化缓存, 给首节点一个默认值
func OpenCache() *Cache {
	cache := Cache{
		head: nil,
		size: 0,
	}
	return &cache
}

// 向缓存中链表插入数据
func (cache *Cache) InsertCacheNode(val int) {
	newNode := &CacheLinkedNode{
		cacheVal: val,
		next: nil,
	}
    node := cache.SearchCacheNode(val)
    // 若该数据节点存在
    if node != nil {
    	// 删除该节点，将该节点插入链表头部
		cache.DelCacheNode(node)
  	}else if cache.size >= MAX_CACHE_SIZE{
  		// 若该数据节点不存在，查看缓存大小，超过缓存删除链表尾部节点，将新节点插入链表头部
		cache.DelCacheTail()
	}
	newNode.next = cache.head
	cache.head = newNode
	cache.size = cache.size+CACHE_NODE_SIZE
}

// 在缓存队列中查找数据节点是否存在
func (cache *Cache) SearchCacheNode(val int) *CacheLinkedNode {
	pNode := cache.head
	for pNode != nil && pNode.cacheVal != val {
		pNode = pNode.next
	}
	return pNode
}

func (cache *Cache) DelCacheNode(delNode *CacheLinkedNode)   {
	pNode := cache.head
	for pNode.next != delNode {
		pNode = pNode.next
	}
	nextNode := delNode.next
	pNode.next = nextNode
}

// 删除链表尾部数据节点
func (cache *Cache) DelCacheTail()  {
	pNode := cache.head
	nextNode := pNode.next
	for nextNode.next != nil {
		pNode = nextNode
		nextNode = nextNode.next
	}
	pNode.next = nextNode.next
}