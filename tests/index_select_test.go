package tests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/horacio/tensor/models"
)

func TestIndexSelectOne(t *testing.T) {
	fmt.Println(string("\033[36mTesting Index Select \033[37m"))
	fmt.Println(string("\033[36mIndexSelect([1, 2, 3, 4], 0, [0, 0, 2]) => [1, 1, 3] \033[37m"))

	tensor := models.NewTensor([][]int{
		{1, 2, 3, 4},
	}, []int{2, 2})

	matrix, err := models.IndexSelect(tensor, 0, []int{0, 0, 2})

	if err != nil {
		panic(err)
	}

	result := [][]int{{1, 1, 3}}
	if !reflect.DeepEqual(matrix, result) {
		t.Fatal("Invalid result")
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix with selected indexes: %v ", matrix)
}

func TestIndexSelectTwo(t *testing.T) {
	fmt.Println(string("\033[36mTesting Index Select \033[37m"))
	fmt.Println(string("\033[36mIndexSelect([[1, 2], [3, 4]], 0, [0]) => [[1, 2]] \033[37m"))

	tensor := models.NewTensor([][]int{
		{1, 2}, {3, 4},
	}, []int{2, 2})

	matrix, err := models.IndexSelect(tensor, 0, []int{0})

	if err != nil {
		panic(err)
	}

	result := [][]int{{1, 2}}
	if !reflect.DeepEqual(matrix, result) {
		t.Fatal("Invalid result")
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix with selected indexes: %v ", matrix)
}

func TestIndexSelectThree(t *testing.T) {
	fmt.Println(string("\033[36mTesting Index Select \033[37m"))
	fmt.Println(string("\033[36mIndexSelect([[1, 2], [3, 4]], 0, [0, 0]) => [[1, 2], [1, 2]] \033[37m"))

	tensor := models.NewTensor([][]int{
		{1, 2}, {3, 4},
	}, []int{2, 2})

	matrix, err := models.IndexSelect(tensor, 0, []int{0, 0})

	if err != nil {
		panic(err)
	}

	result := [][]int{{1, 2}, {1, 2}}
	if !reflect.DeepEqual(matrix, result) {
		t.Fatal("Invalid result")
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix with selected indexes: %v ", matrix)
}

func TestIndexSelectFour(t *testing.T) {
	fmt.Println(string("\033[36mTesting Index Select \033[37m"))
	fmt.Println(string("\033[36mIndexSelect([[1, 2], [3, 4]], 0, [0, 0, 1, 1]) => [[1, 2], [1, 2], [3, 4], [3, 4]] \033[37m"))

	tensor := models.NewTensor([][]int{
		{1, 2}, {3, 4},
	}, []int{2, 2})

	matrix, err := models.IndexSelect(tensor, 0, []int{0, 0, 1, 1})

	if err != nil {
		panic(err)
	}

	result := [][]int{{1, 2}, {1, 2}, {3, 4}, {3, 4}}
	if !reflect.DeepEqual(matrix, result) {
		t.Fatal("Invalid result")
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix with selected indexes: %v ", matrix)
}

func TestIndexSelectFive(t *testing.T) {
	fmt.Println(string("\033[36mTesting Index Select \033[37m"))
	fmt.Println(string("\033[36mIndexSelect([[1, 2], [3, 4]], 1, [0])\t=> [[1], [3]] \033[37m"))

	tensor := models.NewTensor([][]int{
		{1, 2}, {3, 4},
	}, []int{2, 2})

	matrix, err := models.IndexSelect(tensor, 1, []int{0})

	if err != nil {
		panic(err)
	}

	result := [][]int{{1}, {3}}
	if !reflect.DeepEqual(matrix, result) {
		t.Fatal("Invalid result")
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix with selected indexes: %v ", matrix)
}

func TestIndexSelectSix(t *testing.T) {
	fmt.Println(string("\033[36mTesting Index Select \033[37m"))
	fmt.Println(string("\033[36mIndexSelect([[1, 2], [3, 4]], 1, [0, 0]) => [[1, 1], [3, 3]] \033[37m"))
	fmt.Printf("Testing Index Select...\n\n")

	tensor := models.NewTensor([][]int{
		{1, 2}, {3, 4},
	}, []int{2, 2})

	matrix, err := models.IndexSelect(tensor, 1, []int{0, 0})

	if err != nil {
		panic(err)
	}

	result := [][]int{{1, 1}, {3, 3}}
	if !reflect.DeepEqual(matrix, result) {
		t.Fatal("Invalid result")
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix with selected indexes: %v ", matrix)
}

// => IndexSelect([[1, 2], [3, 4]], 1, [0, 0, 1, 1]) => [[1, 1, 2, 2], [3, 3, 4, 4]]
func TestIndexSelectSeven(t *testing.T) {
	fmt.Println(string("\033[36mTesting Index Select \033[37m"))
	fmt.Println(string("\033[36mIndexSelect([[1, 2], [3, 4]], 1, [0, 0, 1, 1]) => [[1, 1, 2, 2], [3, 3, 4, 4]] \033[37m"))

	tensor := models.NewTensor([][]int{
		{1, 2}, {3, 4},
	}, []int{2, 2})

	matrix, err := models.IndexSelect(tensor, 1, []int{0, 0, 1, 1})

	if err != nil {
		panic(err)
	}

	result := [][]int{{1, 1, 2, 2}, {3, 3, 4, 4}}
	if !reflect.DeepEqual(matrix, result) {
		t.Fatal("Invalid result")
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix with selected indexes: %v ", matrix)
}
