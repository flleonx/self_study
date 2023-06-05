package main

import (
	"fmt"
)

func infinite() {
	go func() {
		for {
			select {
			default:
				fmt.Println("DOING WORK")
			}
		}
	}()

}
