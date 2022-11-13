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
	mahasiswa := &doubleListData{
		age:  age,
		name: name,
	}

	if data.head == nil {
		data.head = mahasiswa
		data.tail = mahasiswa
	} else {
		data.head.prev = mahasiswa
		mahasiswa.next = data.head
		data.head = mahasiswa
	}

	return data
}

// PushTail function for insert elemnt from tail
func (data doubleListClient) PushTail(name string, age int) doubleListClient {
	mahasiswa := &doubleListData{
		age:  age,
		name: name,
	}

	if data.tail == nil {
		data.head = mahasiswa
		data.tail = mahasiswa
	} else {
		data.tail.next = mahasiswa
		mahasiswa.prev = data.tail
		data.tail = mahasiswa
	}

	return data
}

// PopHead function for delete first element
func (data doubleListClient) PopHead() doubleListClient {
	if data.head == nil {
		fmt.Println("No Data")
	} else if data.head == data.tail {
		data.head = nil
		data.tail = nil
	} else {
		data.head = data.head.next
		data.head.prev = nil
	}

	return data
}

// PopTail function for delete last element
func (data doubleListClient) PopTail() doubleListClient {
	if data.tail == nil {
		fmt.Println("No Data")
	} else if data.head == data.tail {
		data.head = nil
		data.tail = nil
	} else {
		data.tail = data.tail.prev
		data.tail.next = nil
	}

	return data
}

// PrintFromHead function for print all element from head
func (data doubleListClient) PrintFromHead() {
	mahasiswa := data.head
	fmt.Println("Double List Print from Head")
	for mahasiswa != nil {
		fmt.Println("=========================")
		fmt.Printf("NAME: %s\n", mahasiswa.name)
		fmt.Printf("AGE:  %d\n", mahasiswa.age)
		mahasiswa = mahasiswa.next
	}
	fmt.Println("=========================")
}

// PrintFromTail function for print all element from tail
func (data doubleListClient) PrintFromTail() {
	mahasiswa := data.tail
	fmt.Println("Double List Print from Tail")
	for mahasiswa != nil {
		fmt.Println("=========================")
		fmt.Printf("NAME: %s\n", mahasiswa.name)
		fmt.Printf("AGE:  %d\n", mahasiswa.age)
		mahasiswa = mahasiswa.prev
	}
	fmt.Println("=========================")
}

// Length function for get data.length of list
func (data doubleListClient) Length() int {
	return data.length
}
