package removeDuplicates

import (
	"errors"
	"fmt"
)

func Remove(s []int) (string, error) {
	if s != nil {
		keyListbool := make(map[int]bool)
		listResults := []int{}
		for _, v := range s {
			if value := keyListbool[v]; !value {
				keyListbool[v] = true
				listResults = append(listResults, v)
			}
		}
		return fmt.Sprint(listResults), nil
	}
	return "", errors.New("empty slice")
}
