package util

import (
	"fmt"
)

// to build key with same separator
// and also have same prefix and id order, we use this function
func BuildCommonKey(prefix string, id string) string {
	return fmt.Sprintf("%s:%s", prefix, id)
}
