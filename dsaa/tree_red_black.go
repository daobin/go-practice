package main

/*
二叉树，又称二叉排序树、二叉搜索树

性质：
1、左子树上的所有节点均小于当前节点，右子树反之
2、左右子树也是二叉排序树
3、没有键值相等的节点
*/

/*
红黑树

时间复杂度：O(logn)

本质：是一种二叉搜索树，相对平衡树（使用二叉树的存储结构，以及B树的操作逻辑）

定义：
1、所有节点不是黑色的就是红色的
2、根节点是黑色的
3、所有叶子节点是黑色的空节点
4、从每个叶子节点到根节点的所有路径上不能有两个连接的红色节点
5、从任一节点到其下每个叶子节点的所有路径都包含相同数目的黑色节点（这就要求插入红色节点，以便减少颜色修正次数）

插入步骤：查找、插入、修正（利用变色、旋转等方法修正）

删除步骤：查找、删除、修正（利用变色、旋转等方法修正）
*/

type rbNode struct {
	parent *rbNode
	left   *rbNode
	right  *rbNode
}

type rbTree struct {
	root *rbNode
}

func (rb *rbTree) insert() {

}

func (rb *rbTree) delete() {
}

func (rb *rbTree) get() {

}

func main() {

}
