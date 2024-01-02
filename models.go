package main

import (
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Description string `json:"description"`
	Price uint `json:"price"`
}

func createBook(db *gorm.DB, book *Book) error{
	result := db.Create(&book)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil
}

func getBooks(db *gorm.DB) []Book {
	var books []Book

	result := db.Find(&books)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return books
}

func getBook(db *gorm.DB, id string) Book {
	var book Book

	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return book
}

func updateBook(db *gorm.DB, book *Book) error{
	result := db.Save(&book)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil
}

func deleteBook(db *gorm.DB, id int) error{
	var book Book
	result := db.Delete(&book, id)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil
}

func searchBooks(db *gorm.DB, name string) []Book {
	var books []Book

	result := db.Where("name = ?",  name ).Order("price").Find(&books)

	if result.Error != nil {
		log.Fatal("search books error: ", result.Error)
	}

	return books
}
