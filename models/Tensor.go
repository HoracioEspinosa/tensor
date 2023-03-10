package models

import (
	"errors"
	"fmt"
)

type Tensor struct {
	Range   int
	Sizes   []int
	Weights []int
	Data    [][]int
}

func NewTensor(data [][]int, sizes []int) Tensor {
	obj := Tensor{
		Sizes:   sizes,
		Range:   len(sizes),
		Weights: []int{},
		Data:    data,
	}

	obj.Process()

	return obj
}

func (t Tensor) Process() {
	sizes := t.Sizes
	weights := make([]int, len(sizes))
	weights[0] = 1
	for i := range sizes {
		if i > 0 {
			weights[i] = weights[i-1] * sizes[i-1]
		}
	}
	t.Weights = weights
}

func IndexSelect(tensor Tensor, dimension int, indexes []int) ([][]int, error) {
	if len(tensor.Data) <= 0 {
		return nil, errors.New("tensor cannot be empty")
	}

	if dimension < 0 {
		return nil, errors.New("the dimension cannot a negative number")
	}

	if len(indexes) <= 0 {
		return nil, errors.New("indexes cannot be empty")
	}

	var result [][]int
	switch dimension {
	case 0:
		if len(tensor.Data) > 1 {
			for _, i := range indexes {
				result = append(result, tensor.Data[i])
			}
		} else {
			result = append(result, GetArrayUsingIndexes(tensor.Data[0], indexes))
		}
	case 1:
		for _, v := range tensor.Data {
			var tempArray []int
			tempArray = append(tempArray, v...)
			result = append(result, GetArrayUsingIndexes(tempArray, indexes))
		}

	default:
		return nil, fmt.Errorf("the dimension %v is not implemented", dimension)
	}

	return result, nil
}

func Reshape(tensor Tensor, newShape []int) (Tensor, error) {
	if len(newShape) <= 0 {
		return Tensor{}, errors.New("new shape cannot be empty")
	}

	if len(tensor.Data) <= 0 {
		return Tensor{}, errors.New("tensor cannot be empty")
	}

	var numberOfVectors = newShape[0]
	var numberOfColsForEachVector = newShape[1]

	arrayOfArrays := make([][]int, numberOfVectors)
	var indexToContinue = 0
	var tensorSlice = tensor.Data[0]

	for i := 0; i < numberOfVectors; i++ {
		array := make([]int, numberOfColsForEachVector)
		for j := 0; j < numberOfColsForEachVector; j++ {
			if len(tensorSlice) > indexToContinue {
				array[j] = tensorSlice[indexToContinue]
				indexToContinue++
			}
		}
		arrayOfArrays[i] = array
	}
	tensor.Sizes = newShape
	tensor.Data = arrayOfArrays

	return tensor, nil
}

func HadamardProduct(tensorOne, tensorTwo Tensor) ([][]int, error) {
	matrixA := tensorOne.Data
	matrixB := tensorTwo.Data

	if (len(matrixA) != len(matrixB)) || (len(matrixA[0]) != len(matrixB[0])) {
		return nil, errors.New("the arrays do not have the same dimensions")
	}

	resultingMatrix := make([][]int, len(matrixA))

	for i := 0; i < len(matrixA); i++ {
		array := make([]int, len(matrixA[0]))
		for j := 0; j < len(matrixA[0]); j++ {
			array[j] = matrixA[i][j] * matrixB[i][j]
		}
		resultingMatrix[i] = array
	}

	return resultingMatrix, nil
}

func GetArrayUsingIndexes(array []int, indexes []int) []int {
	var tempArray []int
	for _, i := range indexes {
		tempArray = append(tempArray, array[i])
	}
	return tempArray
}
