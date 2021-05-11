package main

import "fmt"

func main() {
	ht := NewHashTable(5)
	ht.Set("abc", 123)
	ht.Set("bcd", 456)
	ht.Set("efg", 789)
	ht.Set("klm", 999)

	fmt.Println(ht.data)
	fmt.Println(ht.Get("abc"))
	fmt.Println(ht.Get("bcd"))
	fmt.Println(ht.Get("efg"))
	fmt.Println(ht.Get("klm"))
	fmt.Println(ht.Keys())

	fmt.Println(FirstRepeatedNumber([]int{}))
	fmt.Println(FirstRepeatedNumber([]int{2, 5, 1, 2, 3, 5, 1, 2, 4}))
	fmt.Println(FirstRepeatedNumber([]int{2, 1, 1, 2, 3, 3, 1, 2, 4}))
	fmt.Println(FirstRepeatedNumber([]int{2, 3, 4, 5}))

	fmt.Println(FirstRepeatedNumber2([]int{2, 5, 5, 1, 2, 3, 5, 1, 2, 4}))
}