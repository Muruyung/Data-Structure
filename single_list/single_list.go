package singlelist

import "fmt"

type Mahasiswa struct {
	nim  int
	nama string
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
	}
}

// PopTail function for delete last element
func PopTail() {
	var mahasiswa *Mahasiswa

	if Tail == nil {
		fmt.Println("No Data")
	} else if Head == Tail {
		Head = nil
		Tail = nil
	} else {
		mahasiswa = Head
		for mahasiswa.next != Tail {
			mahasiswa = mahasiswa.next
		}

		mahasiswa.next = nil
		Tail = mahasiswa
	}
}

// Print function for print all element
func Print() {
	mahasiswa := Head
	fmt.Println("Single List Print")
	for mahasiswa != nil {
		fmt.Println("=========================")
		fmt.Printf("NIM:  %d\n", mahasiswa.nim)
		fmt.Printf("NAMA: %s\n", mahasiswa.nama)
		mahasiswa = mahasiswa.next
	}
	fmt.Println("=========================")
}
