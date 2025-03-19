package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res, err := Eval(scanner.Text())
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println(res)
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		os.Exit(1)
	}
}
