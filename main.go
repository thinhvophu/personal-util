package main

import (
	"com.github.thinhvp/utility/credit"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Please specify input filepath")
	} else {
		filePath := os.Args[1]
		fmt.Println("Converting file: " + filePath + ". Please wait...")
		if strings.Contains(filePath, "citi") {
			result, err := credit.ConvertCiti(filePath)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(result)
		} else if strings.Contains(filePath, "vcb") {
			result, err := credit.ConvertVCB(filePath)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(result)
		} else {
			fmt.Println("Error: couldn't recognize bank name from file path.")
		}
	}
}
