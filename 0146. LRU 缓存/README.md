#### [146. LRU 缓存](https://leetcode-cn.com/problems/lru-cache/)

> 难度中等

请你设计并实现一个满足 [LRU (最近最少使用) 缓存](https://baike.baidu.com/item/LRU) 约束的数据结构。

实现 `LRUCache` 类：

- `LRUCache(int capacity)` 以 **正整数** 作为容量 `capacity` 初始化 LRU 缓存
- `int get(int key)` 如果关键字 `key` 存在于缓存中，则返回关键字的值，否则返回 `-1` 。
- `void put(int key, int value)` 如果关键字 `key` 已经存在，则变更其数据值 `value` ；如果不存在，则向缓存中插入该组 `key-value` 。如果插入操作导致关键字数量超过 `capacity` ，则应该 **逐出** 最久未使用的关键字。

函数 `get` 和 `put` 必须以 `O(1)` 的平均时间复杂度运行。

**示例：**

```
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
```

**提示：**

- `1 <= capacity <= 3000`
- `0 <= key <= 10000`
- `0 <= value <= 105`
- 最多调用 `2 * 105` 次 `get` 和 `put`

#### 解题思路

> 参考1：https://leetcode-cn.com/problems/lru-cache/solution/jian-dan-shi-li-xiang-xi-jiang-jie-lru-s-exsd/
>
> 参考2：https://leetcode-cn.com/problems/lru-cache/solution/lruhuan-cun-ji-zhi-by-leetcode-solution/

分析上面的操作过程，要让 `put` 和 `get` 方法的时间复杂度为 O(1)，我们可以总结出 `cache` 这个数据结构必要的条件：

1、显然 `cache` 中的元素必须有时序，以区分最近使用的和久未使用的数据，当容量满了之后要删除最久未使用的那个元素腾位置。

2、我们要在 `cache` 中快速找某个 `key` 是否已存在并得到对应的 `val`；

3、每次访问 `cache` 中的某个 `key`，需要将这个元素变为最近使用的，也就是说 `cache` 要支持在任意位置快速插入和删除元素。

那么，什么数据结构同时符合上述条件呢？哈希表查找快，但是数据无固定顺序；链表有顺序之分，插入删除快，但是查找慢。所以结合一下，形成一种新的数据结构：哈希链表 `LinkedHashMap`。

LRU 缓存算法的核心数据结构就是哈希链表，双向链表和哈希表的结合体。这个数据结构长这样：

[![img](https://labuladong.gitee.io/algo/images/LRU%e7%ae%97%e6%b3%95/4.jpg)](https://labuladong.gitee.io/algo/images/LRU算法/4.jpg)

借助这个结构，我们来逐一分析上面的 3 个条件：

1、如果我们每次默认从链表尾部添加元素，那么显然越靠尾部的元素就是最近使用的，越靠头部的元素就是最久未使用的。

2、对于某一个 `key`，我们可以通过哈希表快速定位到链表中的节点，从而取得对应 `val`。

3、链表显然是支持在任意位置快速插入和删除的，改改指针就行。只不过传统的链表无法按照索引快速访问某一个位置的元素，而这里借助哈希表，可以通过 `key` 快速映射到任意一个链表节点，然后进行插入和删除。

==**思考这样2个问题:**==
1.使用单向链表怎么样，力扣上有一题237就是只能访问被删除节点来完成从节点在链表中的删除 时间复杂度O(1)
乍一看是可以的，但在添加元素超出容量，需要删除靠近尾节点时，所需的时间并不是 O(1),因为还要遍历确定删除哪个节点
2.双向链表节点里只存value,不存key 不行吗
乍一想好像在链表节点中，key也没有用到。但事实上key是用到了，还是在关键位置，那就是删除最久不用元素时，不仅需要在链表中删除该节点，还需要在哈希表中删除该节点。

#### 代码

```go
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

```

