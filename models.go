package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name string
	Author string
	Description string
	Price uint
}

func createBook(db *gorm.DB, book *Book) {
	result := db.Create(&book)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Book created successfully!")
}

func getBook(db *gorm.DB, id int) *Book {
	book := &Book{}

	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return book
}

func updateBook(db *gorm.DB, book *Book) {
	result := db.Save(&book)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Book updated successfully!")
}

func deleteBook(db *gorm.DB, book *Book) {
	result := db.Delete(&book)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Book deleted successfully!")
}
