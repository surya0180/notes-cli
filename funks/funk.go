package funks

import "fmt"

func transformArr(array *[]int, transformer func(int) int) []int {
	newArray := []int{}
	for _, num := range *array {
		newArray = append(newArray, transformer(num))
	}
	return newArray
}

func Execute() {
	array := []int{1, 2, 3, 4, 5}

	doubleTransformer := func(number int) int {
		return number * 2
	}
	tripleTransformer := func(number int) int {
		return number * 3
	}

	doubled := transformArr(&array, doubleTransformer)
	tripled := transformArr(&array, tripleTransformer)

	fmt.Println(doubled)
	fmt.Println(tripled)
}
