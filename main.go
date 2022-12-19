package main

import (
	"fmt"
	"github.com/horacio/tensor/models"
)

func main() {
	fmt.Println("\033[36mExercises \033[37m")

	// Index Select
	CallIndexSelect([][]int{{1, 2, 3, 4}}, []int{2, 2}, []int{0, 0, 2}, 0)      // {1,1,3}
	CallIndexSelect([][]int{{1, 2}, {3, 4}}, []int{2, 2}, []int{0}, 0)          // {1,2}
	CallIndexSelect([][]int{{1, 2}, {3, 4}}, []int{2, 2}, []int{0, 0}, 0)       // {[1, 2], [1, 2]}
	CallIndexSelect([][]int{{1, 2}, {3, 4}}, []int{2, 2}, []int{0, 0, 1, 1}, 0) // [[1, 2], [1, 2], [3, 4], [3, 4]]
	CallIndexSelect([][]int{{1, 2}, {3, 4}}, []int{2, 2}, []int{0}, 1)          // [[1], [3]]
	CallIndexSelect([][]int{{1, 2}, {3, 4}}, []int{2, 2}, []int{0, 0}, 1)       // [[1, 1], [3, 3]]
	CallIndexSelect([][]int{{1, 2}, {3, 4}}, []int{2, 2}, []int{0, 0, 1, 1}, 1) // [[1, 1, 2, 2], [3, 3, 4, 4]]

	// Tensors
	tensorOne := models.NewTensor([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
	}, []int{4, 4})

	tensorTwo := models.NewTensor([][]int{
		{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}, []int{4, 4})

	// Reshape
	fmt.Printf("\n\033[36mPROBLEM:\tReshape(%v) \033[37m\n", tensorOne.Data)
	reshape, _ := models.Reshape(tensorOne, []int{2, 2, 4})
	fmt.Printf("RESULT:\t\t%v\n", reshape)

	// HadamardProduct
	fmt.Printf("\n\033[36mPROBLEM:\tHadamardProduct(%v : %v) \033[37m\n", tensorOne.Data, tensorTwo.Data)
	hadamardProduct, _ := models.HadamardProduct(tensorOne, tensorTwo)
	fmt.Printf("RESULT:\t\t%v\n", hadamardProduct)
}

func CallIndexSelect(data [][]int, sizes []int, indices []int, dimensions int) {
	fmt.Printf("\n\033[36mPROBLEM:\tIndexSelect(%v, %d, %v) \033[37m\n", data, dimensions, indices)
	result, err := models.IndexSelect(
		models.NewTensor(data, sizes), dimensions, indices,
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("RESULT:\t\t%v\n", result)
}
