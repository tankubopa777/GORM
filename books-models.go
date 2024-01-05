package main

import (
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	PublisherID uint
	Publisher   Publisher
	Authors     []Author `gorm:"many2many:author_books;"`
  }

type Publisher struct {
	gorm.Model
	Details string
	Name    string
  }
  
  type Author struct {
	gorm.Model
	Name  string
	Books []Book `gorm:"many2many:author_books;"`
  }
  
  type AuthorBook struct {
	AuthorID uint
	Author   Author
	BookID   uint
	Book     Book
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

// create publisher
func createPublisher(db *gorm.DB, publisher *Publisher) error{
	result := db.Create(&publisher)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil
}

// createAuthor
func createAuthor(db *gorm.DB, author *Author) error{
	result := db.Create(author)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil
}

// createBookWithAuthor
func createBookWithAuthor(db *gorm.DB, book *Book) error{
	if err := db.Create(&book).Error; err != nil {
		return err
	}

	return nil
}

// getBookWithPublisher
func getBookWithPublisher(db *gorm.DB, bookID uint) (*Book, error) {
	var book Book

	result := db.Preload("Publisher").First(&book, bookID)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return &book, nil
}
