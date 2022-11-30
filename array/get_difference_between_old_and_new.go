package array

import "fmt"

type Option struct {
	Case
}

type Case string

const (
	CreateArray Case = "create"
	DeleteArray Case = "delete"
)

// GetDifference
// compare 2 array and get an array of difference value,
// this array will be used to add or delete the difference object.
//
// Example:
//
// x := []int{1, 2, 3, 4, 7, 8, 9, 10}
//
// y := []int{5, 6}
//
// result := GetDifference[int](Option{Case: array.DeleteArray}, x, y)
//
// result will contain the data that we need to delete in our case
func GetDifference[T any](option Option, oldData []T, newData []T) (data []T) {

	//if old data = nil or old data length = 0 and the new data length > 0
	//return new data
	if (oldData == nil || len(oldData) == 0) && len(newData) > 0 {
		return newData
	}
	//if new data = nil or new data length = 0 and the old data length > 0
	//return old data
	if (newData == nil || len(newData) == 0) && len(oldData) > 0 {
		return oldData
	}

	//set  parent and child based on option case
	//if create then we need to set parent as old and child as new
	//else if delete then parent as new and child as old
	var parent, child []T
	switch {
	case option.Case == CreateArray:
		parent, child = oldData, newData

	case option.Case == DeleteArray:
		parent, child = newData, oldData
	}

	//middle var for storing and compare
	existData := make(map[any]bool)

	// fil data
	for _, p := range parent {
		existData[p] = true
	}

	// comparing step
	for _, ch := range child {
		fmt.Println(existData[ch])
		if !existData[ch] {
			data = append(data, ch)
		}
	}
	fmt.Println(data)
	//return final data
	return
}
