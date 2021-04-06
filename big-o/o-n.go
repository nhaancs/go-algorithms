package main

import "fmt"

func findNemo(arr []string) {
	for i:=0; i < len(arr); i++ {
		if arr[i] == "nemo" {
			fmt.Println("Found NEMO")
		}
	}
}