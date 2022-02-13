package main

import "fmt"

func main() {
	mapEmpty := map[string]int{}
	fmt.Println(mapEmpty)

	mapMake := make(map[string]int)
	fmt.Println(mapMake)

	mapCap := make(map[string]int, 10)
	fmt.Println(mapCap)

	m := map[string]int{
		"John":    42,
		"Richard": 33,
	}
	age := m["John"]
	fmt.Println(age)

	age, ok := m["John"]
	fmt.Println(age, ok)

	_, ok = m["Richard"]
	fmt.Println(ok)

	m["Jane"] = 61
	fmt.Println(m)

	m["Jane"] = 27
	fmt.Println(m["Jane"])

	delete(m, "John")

	_, ok = m["John"]
	fmt.Println(ok)

	var mm map[string]int
	fmt.Println(mm == nil)

	fmt.Println(len(mm))

	v, ok := mm["John"]
	fmt.Println(v, ok)

	delete(mm, "Richard")

	//mm["Jane"] = 61

	followers := []string{"John", "Richard", "John", "Jane", "Jane", "Alen"}
	unique := make([]string, 0, len(followers))

	uniqueMap := make(map[string]struct{})

	for _, v := range followers {
		if _, ok := uniqueMap[v]; ok {
			continue
		}
		unique = append(unique, v)
		uniqueMap[v] = struct{}{}
	}

	fmt.Println(unique)
}
