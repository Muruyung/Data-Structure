package datastruct

import "fmt"

type singleListClient struct {
	head   *singleListData
	tail   *singleListData
	length int
}

type SingleListService interface {
	PushHead(name string, age int) singleListClient
	PushTail(name string, age int) singleListClient
	PopHead() singleListClient
	PopTail() singleListClient
	Print()
	Length() int
}

type singleListData struct {
	name string
	age  int
	next *singleListData
}

// NewSingleList create new single list service
func NewSingleList() SingleListService {
	return singleListClient{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// PushHead function for insert element from head
func (data singleListClient) PushHead(name string, age int) singleListClient {
	tmpData := &singleListData{
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

// PushTail function for insert elemnt from tail
func (data singleListClient) PushTail(name string, age int) singleListClient {
	tmpData := &singleListData{
		name: name,
		age:  age,
	}

	if data.tail == nil {
		data.head = tmpData
		data.tail = tmpData
	} else {
		data.tail.next = tmpData
		data.tail = tmpData
	}
	data.length++

	return data
}

// PopHead function for delete first element
func (data singleListClient) PopHead() singleListClient {
	if data.head == nil {
		fmt.Println("No Data")
	} else {
		data.length--
		if data.head == data.tail {
			data.head = nil
			data.tail = nil
		} else {
			data.head = data.head.next
		}
	}

	return data
}

// PopTail function for delete last element
func (data singleListClient) PopTail() singleListClient {
	var tmpData *singleListData

	if data.tail == nil {
		fmt.Println("No Data")
	} else {
		data.length--
		if data.head == data.tail {
			data.head = nil
			data.tail = nil
		} else {
			tmpData = data.head
			for tmpData.next != data.tail {
				tmpData = tmpData.next
			}

			tmpData.next = nil
			data.tail = tmpData
		}
	}

	return data
}

// Print function for print all element
func (data singleListClient) Print() {
	tmpData := data.head
	fmt.Println("Single List Print")
	for tmpData != nil {
		fmt.Println("=========================")
		fmt.Printf("NAME: %s\n", tmpData.name)
		fmt.Printf("AGE:  %d\n", tmpData.age)
		tmpData = tmpData.next
	}
	fmt.Println("=========================")
}

// Length function for get length of list
func (data singleListClient) Length() int {
	return data.length
}
