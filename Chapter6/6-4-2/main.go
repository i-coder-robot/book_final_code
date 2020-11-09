package main

import "fmt"

func main() {
	var shops map[string]map[string]int = map[string]map[string]int {
		"火锅店":map[string]int{"牛肉":168,"羊肉":168,"蔬菜拼盘":98},
		"披萨店":map[string]int{"超级至尊披萨":138,"鸡翅":48,"奶油蘑菇汤":68,"香草凤尾虾":78},
	}
	fmt.Println(shops)
}
