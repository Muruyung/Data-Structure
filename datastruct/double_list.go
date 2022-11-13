package datastruct

import "fmt"

type doubleListClient struct {
	head   *doubleListData
	tail   *doubleListData
	length int
}

type DoubleListService interface {
	PushHead(name string, age int) doubleListClient
	PushTail(name string, age int) doubleListClient
	PopHead() doubleListClient
	PopTail() doubleListClient
	PrintFromHead()
	PrintFromTail()
}

type doubleListData struct {
	name string
	age  int
	prev *doubleListData
	next *doubleListData
}

func NewDoubleList() DoubleListService {
	return doubleListClient{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// PushHead function for insert element from head
func (data doubleListClient) PushHead(name string, age int) doubleListClient {
	tmpData := &doubleListData{
		age:  age,
		name: name,
	}

	if data.head == nil {
		data.head = tmpData
		data.tail = tmpData
	} else {
		data.head.prev = tmpData
		tmpData.next = data.head
		data.head = tmpData
	}
	data.length++

	return data
}

// PushTail function for insert elemnt from tail
func (data doubleListClient) PushTail(name string, age int) doubleListClient {
	tmpData := &doubleListData{
		age:  age,
		name: name,
	}

	if data.tail == nil {
		data.head = tmpData
		data.tail = tmpData
	} else {
		data.tail.next = tmpData
		tmpData.prev = data.tail
		data.tail = tmpData
	}
	data.length++

	return data
}

// PopHead function for delete first element
func (data doubleListClient) PopHead() doubleListClient {
	if data.head == nil {
		fmt.Println("No Data")
	} else {
		if data.head == data.tail {
			data.head = nil
			data.tail = nil
		} else {
			data.head = data.head.next
			data.head.prev = nil
		}
		data.length--
	}

	return data
}

// PopTail function for delete last element
func (data doubleListClient) PopTail() doubleListClient {
	if data.tail == nil {
		fmt.Println("No Data")
	} else {
		if data.head == data.tail {
			data.head = nil
			data.tail = nil
		} else {
			data.tail = data.tail.prev
			data.tail.next = nil
		}
		data.length--
	}

	return data
}

// PrintFromHead function for print all element from head
func (data doubleListClient) PrintFromHead() {
	tmpData := data.head
	fmt.Println("Double List Print from Head")
	for tmpData != nil {
		fmt.Println("=========================")
		fmt.Printf("NAME: %s\n", tmpData.name)
		fmt.Printf("AGE:  %d\n", tmpData.age)
		tmpData = tmpData.next
	}
	fmt.Println("=========================")
}

// PrintFromTail function for print all element from tail
func (data doubleListClient) PrintFromTail() {
	tmpData := data.tail
	fmt.Println("Double List Print from Tail")
	for tmpData != nil {
		fmt.Println("=========================")
		fmt.Printf("NAME: %s\n", tmpData.name)
		fmt.Printf("AGE:  %d\n", tmpData.age)
		tmpData = tmpData.prev
	}
	fmt.Println("=========================")
}

// Length function for get data.length of list
func (data doubleListClient) Length() int {
	return data.length
}
