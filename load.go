/*
@Time : 2023/2/1 上午10:47
@Author : MuYiMing
@File : loadConf
@Software : GoLand
*/

package load

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strings"
)

type LoadController struct {
	TreeMap map[string]*Tree
}

func NewLoadController() *LoadController {
	var lc = &LoadController{
		TreeMap: make(map[string]*Tree),
	}
	return lc
}

type ConfDataModle struct {
	WaterLoggings      []FieldDataModle `json:"waterLoggings"`
	PublicHealths      []FieldDataModle `json:"publicHealths"`
	ForestFires        []FieldDataModle `json:"forestFires"`
	GovernmentServices []FieldDataModle `json:"governmentServices"`
}

type FieldDataModle struct {
	Intention string `json:"intention"`
	Keywords  string `json:"keywords"`
	Sql       string `json:"sql"`
	Args      string `json:"args"`
}

func (lc *LoadController) LoadConfFile(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	var fields = make(map[string][]FieldDataModle)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		return err
	}

	for k, v := range fields {
		t := NewTree(k)
		for i := 0; i < len(v); i++ {
			arr := strings.Split(v[i].Keywords, ",")
			sort.Strings(arr)
			node := NewTreeNode(v[i].Sql, arr)
			t.AddTreeNode(node)
		}
		lc.TreeMap[t.Name] = t
	}
	return nil
}

func (lc *LoadController) SearchNode(fieldName string, keys []string) *TreeNode {
	if len(fieldName) == 0 || len(keys) == 0 {
		return nil
	}
	if v, ok := lc.TreeMap[fieldName]; ok {
		return v.SearchTreeNode(keys)
	}
	return nil
}
