package testutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestName(t *testing.T) {
	index := 999
	name := "my_test"

	testName := TestName(t, name, index)
	assert.Equal(t, testName, fmt.Sprintf("%d-%s", index, name))
}
