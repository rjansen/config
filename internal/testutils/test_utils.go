package testutils

import (
	"fmt"
	"testing"
)

func TestName(_ *testing.T, name string, index int) string {
	return fmt.Sprintf("%d-%s", index, name)
}
