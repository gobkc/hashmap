package hashmap

import "hash/fnv"

const (
	bucketsNode   = 100
	bucketLenSpan = 4
)

type Bucket struct {
	nodes []*Node
	Next  *Bucket
}

type Node struct {
	K       string
	V       any
	deleted bool
}

type Hash struct {
	bucketNum int64
	buckets   []*Bucket
}

func NewHash() *Hash {
	return &Hash{}
}

func (h *Hash) Get(key string) *Node {
	pos := h.getPosAndNodePos(key)
	curNode := h.getCurrentNode(key, pos.curBucket, pos.NodePos)
	return curNode
}

func Bind[T any](node *Node) T {
	if convert, ok := node.V.(T); ok {
		return convert
	}
	return *new(T)
}

func (h *Hash) Set(key string, value any) {
	pos := h.getPosAndNodePos(key)
	curNode := h.getCurrentNode(key, pos.curBucket, pos.NodePos)
	curNode.K = key
	curNode.V = value
	curNode.deleted = false
}

func (h *Hash) Delete(key string) {
	pos := h.getPosAndNodePos(key)
	curNode := h.getCurrentNode(key, pos.curBucket, pos.NodePos)
	curNode.K = ``
	curNode.V = nil
	curNode.deleted = true
}

func (h *Hash) Has(key string) bool {
	pos := h.getPosAndNodePos(key)
	curNode := h.getCurrentNode(key, pos.curBucket, pos.NodePos)
	if (curNode.K == `` && curNode.V == nil) || curNode.deleted == true {
		return false
	}
	return true
}

type getPosAndNodePosResp struct {
	Pos       int64
	NodePos   int
	curBucket *Bucket
}

func (h *Hash) getPosAndNodePos(key string) *getPosAndNodePosResp {
	pos := int64(getBucketPos(key))
	hash := getHash(key)
	if h.bucketNum == 0 {
		h.bucketNum += pos
		h.buckets = make([]*Bucket, h.bucketNum)
	} else if waitAppend := pos - h.bucketNum; waitAppend > 0 {
		newBuckets := make([]*Bucket, waitAppend)
		h.bucketNum += waitAppend
		h.buckets = append(h.buckets, newBuckets...)
	}
	nodePos := hash % bucketsNode
	curBucket := h.buckets[pos-1]
	if curBucket == nil {
		curBucket = &Bucket{
			nodes: make([]*Node, bucketsNode),
		}
		h.buckets[pos-1] = curBucket
	}
	newPos := &getPosAndNodePosResp{
		Pos:       pos,
		NodePos:   nodePos,
		curBucket: curBucket,
	}
	return newPos
}

func (h *Hash) getCurrentNode(key string, buckets *Bucket, nodePos int) *Node {
	currentNode := buckets.nodes[nodePos]
	if currentNode == nil {
		currentNode = &Node{}
		buckets.nodes[nodePos] = currentNode
		return currentNode
	}
	//hash conflict occurred
	if currentNode.K != key {
		if buckets.Next == nil {
			newBucket := &Bucket{
				nodes: make([]*Node, bucketsNode),
			}
			buckets.Next = newBucket
		}
		return h.getCurrentNode(key, buckets.Next, nodePos)
	}
	return currentNode
}

func getHash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func getBucketPos(key string) (pos int) {
	l := len(key)
	if l%bucketLenSpan == 0 {
		pos = l / bucketLenSpan
	} else {
		pos = l/bucketLenSpan + 1
	}
	return
}
