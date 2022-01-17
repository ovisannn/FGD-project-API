package reslicing

import "errors"

func DeleteItemFromSlice(dataArray []string, item string) ([]string, error) {
	var eraseIndex int
	for index, itemInside := range dataArray {
		if itemInside == item {
			eraseIndex = index
			break
		}
	}
	if eraseIndex == 0 {
		return []string{}, errors.New("item not found")
	}
	dataArray[eraseIndex] = dataArray[len(dataArray)-1]
	dataArray[len(dataArray)-1] = ""
	dataArray = dataArray[:len(dataArray)-1]

	return dataArray, nil
}
