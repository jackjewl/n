package utils

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	gotRsult, err := GenerateToken(1, "accoutn")

	fmt.Println(gotRsult, err)
}
