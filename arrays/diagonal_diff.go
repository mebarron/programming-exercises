package main 

import (
	"fmt"
	"math/rand"
)

func diagonalDifference(arr [][]int32) int32 {
    diag_sum1 := 0 
    diag_sum2 := 0
    
    column := 0
    
    absolute_sum := 0
    
    for _, x := range arr {
        diag_sum1 += int(x[column])
        column++
    }
    
    column = len(arr) - 1
    
    for _, x := range arr { 
        diag_sum2 += int(x[column])
        column-- 
    }
    
    absolute_sum = diag_sum1 - diag_sum2
    
    if absolute_sum < 0 {
		absolute_sum -= absolute_sum + absolute_sum
	}
    
    return int32(absolute_sum)
}

func main() {
	var arr [][]int32 

	for i:=0;i<9;i++ {
		var insert []int32
		for i:=0;i<9;i++ {
			r := rand.Intn(8)
			if r >= 7 {
				r = -1
			}
	
			insert = append(insert, int32(r))
		}
		arr = append(arr, insert) 
	}

	for _, row := range arr {
		fmt.Println(row)
	}

	fmt.Println("\n")

	fmt.Println(diagonalDifference(arr))
}