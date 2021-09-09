package max2Numbers

import (
	"errors"
	"sort"
)

func Get(s []int) (int, error) {
	if s != nil {
		sort.Ints(s)
		return s[cap(s)-2], nil
	}
	return 0, errors.New("empty slice")
}
