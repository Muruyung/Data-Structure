package doublelist

import "fmt"

type Mahasiswa struct {
	nim  int
	nama string
	prev *Mahasiswa
	next *Mahasiswa
}

var Head, Tail *Mahasiswa

// PushHead function for insert element from head
func PushHead(nim int, nama string) {
	mahasiswa := &Mahasiswa{
		nim:  nim,
		nama: nama,
	}

	if Head == nil {
		Head = mahasiswa
		Tail = mahasiswa
	} else {
		Head.prev = mahasiswa
		mahasiswa.next = Head
		Head = mahasiswa
	}
}

// PushTail function for insert elemnt from tail
func PushTail(nim int, nama string) {
	mahasiswa := &Mahasiswa{
		nim:  nim,
		nama: nama,
	}

	if Tail == nil {
		Head = mahasiswa
		Tail = mahasiswa
	} else {
		Tail.next = mahasiswa
		mahasiswa.prev = Tail
		Tail = mahasiswa
	}
}

// PopHead function for delete first element
func PopHead() {
	if Head == nil {
		fmt.Println("No Data")
	} else if Head == Tail {
		Head = nil
		Tail = nil
	} else {
		Head = Head.next
		Head.prev = nil
	}
}

// PopTail function for delete last element
func PopTail() {
	if Tail == nil {
		fmt.Println("No Data")
	} else if Head == Tail {
		Head = nil
		Tail = nil
	} else {
		Tail = Tail.prev
		Tail.next = nil
	}
}

// PrintFromHead function for print all element from head
func PrintFromHead() {
	mahasiswa := Head
	fmt.Println("Double List Print from Head")
	for mahasiswa != nil {
		fmt.Println("=========================")
		fmt.Printf("NIM:  %d\n", mahasiswa.nim)
		fmt.Printf("NAMA: %s\n", mahasiswa.nama)
		mahasiswa = mahasiswa.next
	}
	fmt.Println("=========================")
}

// PrintFromTail function for print all element from tail
func PrintFromTail() {
	mahasiswa := Tail
	fmt.Println("Double List Print from Tail")
	for mahasiswa != nil {
		fmt.Println("=========================")
		fmt.Printf("NIM:  %d\n", mahasiswa.nim)
		fmt.Printf("NAMA: %s\n", mahasiswa.nama)
		mahasiswa = mahasiswa.prev
	}
	fmt.Println("=========================")
}
