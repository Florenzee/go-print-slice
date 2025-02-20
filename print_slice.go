package go_print_slice

import "fmt"

func PrintSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i], " ")
	}

}

func PrintSlice2(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			fmt.Print(slice[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
