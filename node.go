/*
@Time : 2023/2/2 下午4:07
@Author : MuYiMing
@File : node
@Software : GoLand
*/

package load

type TreeNode struct {
	Sql     string
	MextMap map[string]*TreeNode
	//有序key数组
	OrderKeys []string
	index     int
}

func NewTreeNode(sql string, okeys []string) *TreeNode {
	return &TreeNode{
		Sql:       sql,
		MextMap:   make(map[string]*TreeNode),
		OrderKeys: okeys,
		index:     0,
	}
}

func (tn *TreeNode) addTreeNode(node *TreeNode) {
	if node.index > len(node.OrderKeys)-1 {
		return
	}
	key := node.OrderKeys[node.index]
	if _, ok := tn.MextMap[key]; !ok {
		var sql string
		var arr []string
		if node.index == len(node.OrderKeys)-1 {
			sql = node.Sql
		} else if node.index < len(node.OrderKeys)-1 {
			arr = node.OrderKeys[:node.index+1]
		}
		tn.MextMap[key] = NewTreeNode(sql, arr)
	}
	node.index++
	tn.MextMap[key].addTreeNode(node)
}

func (tn *TreeNode) searchTreeNode(keys []string) *TreeNode {
	var nowNode *TreeNode = tn
	for i := 0; i < len(keys); i++ {
		if v, ok := nowNode.MextMap[keys[i]]; ok {
			nowNode = v
		}
	}
	return nowNode
}
