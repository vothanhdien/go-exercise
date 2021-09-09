package findMaxLengthElement

import (
	"errors"
	"fmt"
	"sort"
)

func Find(s []string) (string, error) {
	if s != nil {
		sort.Slice(s, func(i, j int) bool { return len(s[i]) < len(s[j]) })
		for i, v := range s {
			if len(v) == len(s[len(s)-1]) {
				return fmt.Sprint(s[i:]), nil
				break
			}
		}
	}
	return "", errors.New("empty slice")
}
