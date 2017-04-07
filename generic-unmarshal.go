package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gocarina/gocsv"
)

type Book struct {
	Name string `csv:"name"`
	Isbn string `csv:"isbn"`
}

type User struct {
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}

type Counter interface {
	// this function should return a function that takes struct as argument
	getCounterFunc() interface{}
	getCount() int
}

type UserCounter struct {
	cnt int
}

func (uc *UserCounter) getCounterFunc() interface{} {
	f := func(u User) {
		// fmt.Println(">>", u.Name, u.Age)
		uc.cnt++
	}
	return f
}
func (uc *UserCounter) getCount() int {
	return uc.cnt
}

type BookCounter struct {
	bookSet map[Book]struct{}
}

func (bc *BookCounter) getCounterFunc() interface{} {
	bc.bookSet = make(map[Book]struct{})
	f := func(b Book) {
		fmt.Println(">>", b.Name, b.Isbn)
		if _, ok := bc.bookSet[b]; !ok {
			bc.bookSet[b] = struct{}{}
		}
	}
	return f
}
func (bc *BookCounter) getCount() int {
	return len(bc.bookSet)
}

func countRecords(r io.Reader, cntr Counter) (int, error) {
	err := gocsv.UnmarshalToCallback(r, cntr.getCounterFunc())
	if err != nil {
		return 0, err
	}
	return cntr.getCount(), nil
}

func countRecordsTheOldWay(r io.Reader) (int, error) {
	users := []User{}
	err := gocsv.Unmarshal(r, &users)
	if err != nil {
		return 0, err
	}
	return len(users), nil
}

func main() {

	booksData := `name,isbn
Book 1,100
Book 2,101
Book 3,103
Book 1,100
Book 2,101`

	usersData := `name,age
F1 L1,30
F2 L2,20
F3 L3,70`

	cnt, err := countRecords(bytes.NewBufferString(usersData), &UserCounter{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User count:", cnt)
	}

	cnt, err = countRecordsTheOldWay(bytes.NewBufferString(usersData))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User count (old way):", cnt)
	}

	cnt, err = countRecords(bytes.NewBufferString(booksData), &BookCounter{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Unique book count:", cnt)
	}

}
