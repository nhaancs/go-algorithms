package main

import "fmt"

func logAllPairsOfArr(arr []string) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Println(arr[i], arr[j])
		}
	}
}