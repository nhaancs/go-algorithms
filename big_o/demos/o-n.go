package demos

import "fmt"

func FindNemo(arr []string) {
	for i:=0; i < len(arr); i++ {
		if arr[i] == "nemo" {
			fmt.Println("Found NEMO")
		}
	}
}