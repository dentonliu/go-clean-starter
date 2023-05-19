package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginatedList(t *testing.T) {
	list := NewPaginatedList(4, 0, 50)
	assert.Equal(t, 40, list.Offset())
	assert.Equal(t, 20, list.Limit())
}
