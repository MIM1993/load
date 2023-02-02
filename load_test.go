/*
@Time : 2023/2/1 上午10:48
@Author : MuYiMing
@File : load_conf_test
@Software : GoLand
*/

package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"testing"
)

func TestLoad(t *testing.T) {
	l := NewLoadController()

	err := l.LoadConfFile("nlp.config")
	if err != nil {
		fmt.Println(err)
		return
	}
	keys := []string{"hello", "不在线", "ni", "排水设备"}
	//sort.Strings(keys)
	fmt.Println(keys)
	node := l.SearchNode("waterLoggings", keys)
	if node == nil {
		panic("null")
	}
	fmt.Println(node.Sql)
}

func TestLoadFile(t *testing.T) {
	data, err := ioutil.ReadFile("/home/mim/workspace/gopath/src/test/load/nlp.config")
	if err != nil {
		panic(err)
	}
	var tmp map[string][]FieldDataModle = make(map[string][]FieldDataModle)
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		panic(err)

	}

	fmt.Println(tmp)
}

func TestSortString(t *testing.T) {
	arr := []string{"数量（二）", "数量少", "数量", "数数", "数量（一）"}
	sort.Strings(arr)
	fmt.Println(arr)
}
