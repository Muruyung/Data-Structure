package datastruct

import "fmt"

type queueClient struct {
	head   *queueData
	tail   *queueData
	length int
}

type QueueService interface {
	Push(name string, age int) queueClient
	Pop() queueClient
	Length() int
}

type queueData struct {
	name string
	age  int
	next *queueData
}

// NewQueue create new single list service
func NewQueue() QueueService {
	return queueClient{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// Push function for insert element into queue
func (data queueClient) Push(name string, age int) queueClient {
	tmpData := &queueData{
		name: name,
		age:  age,
	}

	if data.head == nil {
		data.head = tmpData
		data.tail = tmpData
	} else {
		tmpData.next = data.head
		data.head = tmpData
	}
	data.length++

	return data
}

// Pop function for print and delete last element
func (data queueClient) Pop() queueClient {
	var tmpData *queueData

	if data.tail == nil {
		fmt.Println("No Data")
	} else {
		if data.head == data.tail {
			data.head = nil
			tmpData = nil
		} else {
			tmpData = data.head
			for tmpData.next != data.tail {
				tmpData = tmpData.next
			}
			tmpData.next = nil
		}

		fmt.Println("Queue Print")
		fmt.Println("=========================")
		fmt.Printf("NAME: %s\n", data.tail.name)
		fmt.Printf("AGE:  %d\n", data.tail.age)
		fmt.Println("=========================")

		data.tail = tmpData
		data.length--
	}

	return data
}

// Length function for get length of queue
func (data queueClient) Length() int {
	return data.length
}
