package _3_map

import (
	"fmt"
	"sort"
	"strconv"
)

func MapSort() {
	m := make(map[string]int)
	m["zhangsan"] = 55
	m["lisi"] = 20
	m["alice"] = 23
	m["wangwu"] = 14
	// 验证key：laoshi是否存在
	if value, ok := m["laoshi"]; ok {
		fmt.Println("key exists, value:", value)
	} else {
		fmt.Println("key not exists")
	}
	// 按key排序，取出key，遍历输出value
	fmt.Println("sort By Key:")
	keys := sortMapByKey(m)
	for _, key := range keys {
		fmt.Println("{"+ key + " " + strconv.Itoa(m[key]) +"}")
	}
	// 按value排序，不用map，用struct，实现sort接口，调用sort.Sort进行排序
	fmt.Println("sort By Value:")
	pairList := sortMapByVaule(m)
	for _, pair := range pairList {
		fmt.Println(pair)
	}
}

type Pair struct {
	Key string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func sortMapByVaule(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func sortMapByKey(m map[string]int) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

