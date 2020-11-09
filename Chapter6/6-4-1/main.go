package main

import "fmt"

func main() {
	var m2 map[string][]string = map[string][]string{
		"火锅店":{"牛肉","羊肉","蔬菜拼盘"},
		"披萨店":{"超级至尊披萨","鸡翅","奶油蘑菇汤","香草凤尾虾"},
	}

	fmt.Println(m2)

	for k,v :=range m2{
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("=======")
	}

}
