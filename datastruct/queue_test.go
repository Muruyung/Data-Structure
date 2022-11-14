package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueuePush(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queue := NewQueue().Push("test1", 1)
		assert.Equal(t, 1, queue.Length())
		assert.Equal(t, queueData{
			name: "test1",
			age:  1,
			next: queue.head.next,
		}, *queue.head)

		queue = queue.Push("test2", 2)
		assert.Equal(t, 2, queue.Length())
		assert.Equal(t, queueData{
			name: "test2",
			age:  2,
			next: queue.head.next,
		}, *queue.head)
	})
}

func TestQueuePop(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queue := NewQueue().
			Pop().
			Push("test1", 1).
			Pop()
		assert.Equal(t, 0, queue.Length())
		assert.Nil(t, queue.head)
		assert.Nil(t, queue.tail)

		queue = queue.
			Push("test1", 1).
			Push("test2", 2).
			Push("test3", 3)

		assert.Equal(t, 3, queue.Length())
		assert.Equal(t, queueData{
			name: "test1",
			age:  1,
			next: nil,
		}, *queue.tail)

		queue = queue.Pop()
		assert.Equal(t, 2, queue.Length())
		assert.Equal(t, queueData{
			name: "test2",
			age:  2,
			next: nil,
		}, *queue.tail)
	})
}
