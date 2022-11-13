package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleListPushHead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		singleList := NewSingleList().PushHead("test1", 1)
		assert.Equal(t, 1, singleList.Length())
		assert.Equal(t, singleListData{
			name: "test1",
			age:  1,
			next: singleList.head.next,
		}, *singleList.head)
		singleList.Print()

		singleList = singleList.PushHead("test2", 2)
		assert.Equal(t, 2, singleList.Length())
		assert.Equal(t, singleListData{
			name: "test2",
			age:  2,
			next: singleList.head.next,
		}, *singleList.head)
		singleList.Print()
	})
}

func TestSingleListPushTail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		singleList := NewSingleList().PushTail("test1", 1)
		assert.Equal(t, 1, singleList.Length())
		assert.Equal(t, singleListData{
			name: "test1",
			age:  1,
			next: nil,
		}, *singleList.tail)
		singleList.Print()

		singleList = singleList.PushTail("test2", 2)
		assert.Equal(t, 2, singleList.Length())
		assert.Equal(t, singleListData{
			name: "test2",
			age:  2,
			next: nil,
		}, *singleList.tail)
		singleList.Print()
	})
}

func TestSingleListPopHead(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		singleList := NewSingleList().
			PushHead("test1", 1).
			PushHead("test2", 2).
			PushTail("test3", 3)

		assert.Equal(t, 3, singleList.Length())
		assert.Equal(t, singleListData{
			name: "test2",
			age:  2,
			next: singleList.head.next,
		}, *singleList.head)
		singleList.Print()

		singleList = singleList.PopHead()
		assert.Equal(t, 2, singleList.Length())
		assert.Equal(t, singleListData{
			name: "test1",
			age:  1,
			next: singleList.head.next,
		}, *singleList.head)
		singleList.Print()
	})
}

func TestSingleListPopTail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		singleList := NewSingleList().
			PushHead("test1", 1).
			PushHead("test2", 2).
			PushTail("test3", 3)

		assert.Equal(t, 3, singleList.Length())
		assert.Equal(t, singleListData{
			name: "test3",
			age:  3,
			next: nil,
		}, *singleList.tail)
		singleList.Print()

		singleList = singleList.PopTail()
		assert.Equal(t, 2, singleList.Length())
		assert.Equal(t, singleListData{
			name: "test1",
			age:  1,
			next: nil,
		}, *singleList.tail)
		singleList.Print()
	})
}
