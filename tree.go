/*
@Time : 2023/2/2 下午4:07
@Author : MuYiMing
@File : tree
@Software : GoLand
*/

package load

type Tree struct {
	root *TreeNode
	Name string
}

func NewTree(name string) *Tree {
	return &Tree{
		root: NewTreeNode("", nil),
		Name: name,
	}
}

func (t *Tree) AddTreeNode(tn *TreeNode) {
	t.root.addTreeNode(tn)
}

func (t *Tree) SearchTreeNode(keys []string) *TreeNode {
	node := t.root.searchTreeNode(keys)
	if node.Sql == "" {
		return nil
	}
	return node
}
