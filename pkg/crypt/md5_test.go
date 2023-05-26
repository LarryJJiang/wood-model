package crypt

import (
	"fmt"
	"testing"
)

func TestGetMd5(t *testing.T) {
	fmt.Println(GetMd5("goblog"))
}
