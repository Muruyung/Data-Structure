package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleListPushHead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		doubleList := NewDoubleList().PushHead("test1", 1)
		assert.Equal(t, 1, doubleList.Length())
		assert.Equal(t, doubleListData{
			name: "test1",
			age:  1,
			next: doubleList.head.next,
			prev: nil,
		}, *doubleList.head)
		doubleList.PrintFromHead()

		doubleList = doubleList.PushHead("test2", 2)
		assert.Equal(t, 2, doubleList.Length())
		assert.Equal(t, doubleListData{
			name: "test2",
			age:  2,
			next: doubleList.head.next,
			prev: nil,
		}, *doubleList.head)
		doubleList.PrintFromHead()
	})
}

func TestDoubleListPushTail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		doubleList := NewDoubleList().PushTail("test1", 1)
		assert.Equal(t, 1, doubleList.Length())
		assert.Equal(t, doubleListData{
			name: "test1",
			age:  1,
			next: nil,
			prev: doubleList.tail.prev,
		}, *doubleList.tail)
		doubleList.PrintFromTail()

		doubleList = doubleList.PushTail("test2", 2)
		assert.Equal(t, 2, doubleList.Length())
		assert.Equal(t, doubleListData{
			name: "test2",
			age:  2,
			next: nil,
			prev: doubleList.tail.prev,
		}, *doubleList.tail)
		doubleList.PrintFromTail()
	})
}

func TestDoubleListPopHead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		doubleList := NewDoubleList().
			PopHead().
			PushHead("test1", 1).
			PopHead()
		assert.Equal(t, 0, doubleList.Length())
		assert.Nil(t, doubleList.head)
		assert.Nil(t, doubleList.tail)

		doubleList = doubleList.
			PushHead("test1", 1).
			PushHead("test2", 2).
			PushTail("test3", 3)

		assert.Equal(t, 3, doubleList.Length())
		assert.Equal(t, doubleListData{
			name: "test2",
			age:  2,
			next: doubleList.head.next,
			prev: nil,
		}, *doubleList.head)
		doubleList.PrintFromHead()

		doubleList = doubleList.PopHead()
		assert.Equal(t, 2, doubleList.Length())
		assert.Equal(t, doubleListData{
			name: "test1",
			age:  1,
			next: doubleList.head.next,
			prev: nil,
		}, *doubleList.head)
		doubleList.PrintFromHead()
	})
}

func TestDoubleListPopTail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		doubleList := NewDoubleList().
			PopTail().
			PushHead("test1", 1).
			PopTail()
		assert.Equal(t, 0, doubleList.Length())
		assert.Nil(t, doubleList.head)
		assert.Nil(t, doubleList.tail)

		doubleList = doubleList.
			PushHead("test1", 1).
			PushHead("test2", 2).
			PushTail("test3", 3)

		assert.Equal(t, 3, doubleList.Length())
		assert.Equal(t, doubleListData{
			name: "test3",
			age:  3,
			next: nil,
			prev: doubleList.tail.prev,
		}, *doubleList.tail)
		doubleList.PrintFromHead()

		doubleList = doubleList.PopTail()
		assert.Equal(t, 2, doubleList.Length())
		assert.Equal(t, doubleListData{
			name: "test1",
			age:  1,
			next: nil,
			prev: doubleList.tail.prev,
		}, *doubleList.tail)
		doubleList.PrintFromHead()
	})
}
