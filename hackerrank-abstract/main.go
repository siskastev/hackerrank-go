package main

import "fmt"

type myBook struct {
	title, author string
	price         int
}

type Book interface {
	display()
}

func common(book Book) {
	book.display()
}

func (mb myBook) display() {
	fmt.Printf("Title: %v\nAuthor: %v\nPrice: %v\n", mb.title, mb.author, mb.price)
}

func main() {
	var (
		title, author string
		price         int
	)
	fmt.Scanln(&title)
	fmt.Scanln(&author)
	fmt.Scanln(&price)
	mb := myBook{title: title, author: author, price: price}
	//1
	common(mb)
	//2
	var book Book
	book = mb
	book.display()
}
