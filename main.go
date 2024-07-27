package main

import (
	"dsa-go/structures/arrays"
	"fmt"
)

func main() {
	// Initialize with dimensions only
	arr1 := arrays.NewMDArray(2, 3)

	// Initialize with dimensions and data
	arr2 := arrays.NewMDArray(2, 3,
		[]interface{}{1, 4, 2},
		[]interface{}{3, 6, 8},
	)

	// Print the arrays
	fmt.Println(arr1)
	fmt.Println(arr2)

	err := arr1.AddItem([]interface{}{4, 2, 0})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = arr2.AddItem([]interface{}{6, 9, 0}, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(arr1)
	fmt.Println(arr2)

	item, err := arr2.GetItem(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(item)

	err = arr2.RemoveItem(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(arr2)

}
