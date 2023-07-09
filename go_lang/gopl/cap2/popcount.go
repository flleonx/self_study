package main

import "fmt"

// pc[i] is the population count of i
var pc [256]byte

func init() {
  fmt.Println(pc)
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
  fmt.Println(pc)
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
  var count int
  for i := 0; i < 8; i ++ {
    count += int(pc[byte(x>>(i*8))])
  }

  return count
}
