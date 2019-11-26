package main

import "fmt"

func checkMagazine(magazine []string, note []string) {
	magMap := make(map[string]int32)

	for _, value := range magazine {
		magMap[value]++
	}

	for _, noteValue := range note {
		if val, ok := magMap[noteValue]; ok {
			if val > 0 {
				magMap[noteValue]--
			} else {
				fmt.Println("No")
				return
			}
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func main() {}
