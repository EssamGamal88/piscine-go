package main

import (
	"fmt"
	"os"
)

func main() {
	targets := []string{"01", "galaxy", "galaxy 01"}

	for _, arg := range os.Args[1:] {
		for _, target := range targets {
			if arg == target {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
