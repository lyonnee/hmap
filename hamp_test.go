package hmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHmap(t *testing.T) {
	m := New[string, int]()

	m.Store("1", 1)
	m.Store("2", 2)

	assert.Equal(t, 2, m.Len())

	res, ok := m.Load("1")
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, res)

	res, ok = m.Load("3")
	assert.Equal(t, false, ok)

	m.Delete("3")
	assert.Equal(t, 2, m.Len())

	m.Delete("1")
	res, ok = m.Load("1")
	assert.Equal(t, false, ok)
	assert.Equal(t, 1, m.Len())

	m.LoadOrStore("2", 22)
	assert.Equal(t, 1, m.Len())
	res, ok = m.Load("2")
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, res)

	m.Store("3", 3)
	m.Store("3", 33)
	assert.Equal(t, 2, m.Len())
	res, ok = m.Load("3")
	assert.Equal(t, true, ok)
	assert.Equal(t, 33, res)

	m.Clean()
	res, ok = m.Load("3")
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, m.Len())

}
