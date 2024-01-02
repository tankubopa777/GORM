package main

import (
	"fmt"
	"os"

	"log"
	"time"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
  )

  func main() {
	// Configure your PostgreSQL database details here
	  dsn := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold: time.Second, // Slow SQL threshold
		  LogLevel:      logger.Info, // Log level
		  Colorful:      true,        // Enable color
		},
	  )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
	  panic("failed to connect to database")
	}
	// Migrate the schema
	db.AutoMigrate(&Book{})
	fmt.Println("Database migration completed!")

	// Set up a new Fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(getBooks(db))
	})

	// Get one book
	app.Get("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		book := getBook(db, strconv.Itoa(id))
		return c.JSON(book)
	})

	// Create new book
	app.Post("/books", func(c *fiber.Ctx) error {
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}
		err := createBook(db, book)
		
		if err != nil {
			return c.Status(500).SendString("Internal server error")
		}

		return c.JSON(fiber.Map{
			"message": "New book created successfully!",
	})
	})

	// Update book
	app.Put("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		book.ID = uint(id)
		updateBook(db, book)

		return c.JSON(fiber.Map{
			"message": "Book updated successfully!",
		})
	})

	// Delete book
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		err = deleteBook(db, id)
		if err != nil {
			return c.Status(500).SendString("Internal server error")
		}
		return c.JSON(fiber.Map{
			"message": "Book deleted successfully!",
		})
	})
	

	app.Listen(":8080")
	
  }