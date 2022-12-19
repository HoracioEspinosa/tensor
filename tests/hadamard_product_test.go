package tests

import (
	"fmt"
	"testing"

	"github.com/horacio/tensor/models"
)

func TestHadamardProduct(t *testing.T) {
	fmt.Printf("Testing Reshape...\n\n")

	tensorOne := models.NewTensor([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
	}, []int{4, 4})

	tensorTwo := models.NewTensor([][]int{
		{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}, []int{4, 4})

	matrix, err := models.HadamardProduct(tensorOne, tensorTwo)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nResults: \n")
	fmt.Println(string("\033[36m  Result \033[37m"))
	fmt.Printf("New Matrix with selected indexes: %v \n", matrix)
}
