// 插入操作
package linklist

import (
	"errors"
)

func (this *LinkList) Insert(i uint, node *ListNode) error {
	if i < 0 || node == nil || i > this.length {
		return errors.New("节点为空或越界")
	}

	// 从head通过循环依次定位到索引i
	curNode := (*this).head

	for j := uint(0); j < i; j++ {
		curNode = curNode.GetNext()
	}

	node.next = curNode.GetNext()
	curNode.next = node
	this.length++
	return nil
}
