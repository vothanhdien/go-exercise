package max2Numbers

import (
	"errors"
	"sort"
)

func Get(s []int) (int, error) {
	if s == nil {
		return 0, errors.New("empty slice")
	}

	sort.Ints(s)
	return s[len(s)-2], nil
}
