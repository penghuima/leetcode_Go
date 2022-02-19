package leetcode
//面试出现的概率一定很大！！！！  哈希链表（哈希表+双向链表的新数据结构）
/*思考这样2个问题:
1.使用单向链表怎么样，力扣上有一题237就是只能访问被删除节点来完成从节点在链表中的删除 时间复杂度O(1)
乍一看是可以的，但在添加元素超出容量，需要删除靠近尾节点时，所需的时间并不是 O(1),因为还要遍历确定删除哪个节点
2.双向链表节点里只存value,不存key 不行吗
乍一想好像在链表节点中，key也没有用到。但事实上key是用到了，还是在关键位置，那就是删除最久不用元素时，不仅需要在链表中删除该节点，
还需要在哈希表中删除该节点。
*/
//双向链表
type duaLinkNode struct{
	key,value int              //键值
	pre ,next *duaLinkNode   //前后指针
}
//双向链表初始化
func initDuaLinkNode(key,value int) *duaLinkNode{
	return &duaLinkNode{
		key:key,
		value: value,
	}
}

type LRUCache struct {
	capacity int //容量
	size int //链表长度
	cache map[int]*duaLinkNode
	head,tail *duaLinkNode //虚拟 头  尾节点
}


func Constructor(capacity int) LRUCache {
	lru:=LRUCache{
		capacity: capacity,
		head: initDuaLinkNode(0,0),
		tail: initDuaLinkNode(0,0),
		cache: map[int]*duaLinkNode{},
		size: 0,
	}
	//双向链表的初始化 头尾虚拟节点互连
	lru.head.next=lru.tail
	lru.tail.pre=lru.head
	return lru
}
//靠近双链表头节点的表示最近使用的，靠近尾节点的数据是最久未使用的
// 方法1
func (this *LRUCache)removeNode(node *duaLinkNode){
	node.pre.next=node.next
	node.next.pre=node.pre
}
func (this *LRUCache)addToHead(node *duaLinkNode){
	node.pre=this.head
	node.next=this.head.next
	this.head.next.pre=node
	this.head.next=node
}
func (this *LRUCache)moveToHead(node *duaLinkNode){
	this.removeNode(node)
	this.addToHead(node)
}
//删除尾节点 并返回该删除的节点
func (this *LRUCache)removeTail() *duaLinkNode{
	node:=this.tail.pre
	this.removeNode(node)
	return node
}
func (this *LRUCache) Get(key int) int {
	if _,ok:=this.cache[key];!ok{ //不存在返回-1
		return -1
	}
	//存在 将数值返回 并将该节点移动到 头部 表示刚被访问
	node:=this.cache[key]
	//移动到头节点
	this.moveToHead(node)
	return node.value

}

func (this *LRUCache) Put(key int, value int)  {
	if _,ok:=this.cache[key];!ok{
		//键不存在 添加新的
		node:=initDuaLinkNode(key,value)
		this.cache[key]=node
		//节点添加到双向链表里
		this.addToHead(node)
		this.size++
		//如果元素数量超过容量了
		if this.size>this.capacity{
			//将尾节点删掉  从链表中删掉并从哈希表中删掉
			removed:=this.removeTail()
			delete(this.cache,removed.key) //之所以链表中需要同时存储 键-值，是因为删除的时候会用到 key
			this.size--
		}

	}else{
		//更新值
		node:=this.cache[key]
		node.value=value
		this.moveToHead(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
