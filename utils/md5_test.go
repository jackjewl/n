package utils

import (
	"fmt"
	"testing"
)

func TestMD5Encrypt(t *testing.T) {
	gotResult := MD5Encrypt("hello");
	fmt.Println(gotResult)

}
