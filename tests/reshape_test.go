package tests

import (
	"fmt"
	"testing"

	"github.com/horacio/tensor/models"
)

func TestReshape(t *testing.T) {
	fmt.Printf("Testing Reshape...\n\n")

	tensor := models.NewTensor([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
	}, []int{4, 4})

	matrix, err := models.Reshape(tensor, []int{2, 2, 4})

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix with selected indexes: %v \n", matrix)
}
