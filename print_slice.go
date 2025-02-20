package go_print_slice

import "fmt"

func PrintSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
	fmt.Println()
}

func PrintSlice2(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Print(slice[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func CreateSlice(n int) (result [][]int) {
	result = make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	return result
}
