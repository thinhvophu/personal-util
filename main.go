package main

import (
	"com.github.thinhvp/utility/credit"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify input file path")
	} else {
		filePath := os.Args[1]
		fmt.Println("---CONVERTING---")
		fmt.Println("Input file: " + filePath + ".")
		if strings.Contains(filePath, "citi") {
			result, err := credit.ConvertCiti(filePath)
			if err != nil {
				fmt.Println("---CONVERTING ERROR---")
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("---CONVERTING DONE---")
			fmt.Println(result)
		} else if strings.Contains(filePath, "vcb") {
			result, err := credit.ConvertVCB(filePath)
			if err != nil {
				fmt.Println("---CONVERTING ERROR---")
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("---CONVERTING DONE---")
			fmt.Println(result)
		} else {
			fmt.Println("---CONVERTING ERROR---")
			fmt.Println("Error: couldn't recognize bank name from file path.")
		}
	}
}
