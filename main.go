package main

import (
	"fmt"
	"os"

	"log"
	"time"

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

	
	
	currentBook := getBook(db, 1)
	
	currentBook.Name = "The Lord of the Rings"
	currentBook.Author = "J.R.R. Tolkien"

	updateBook(db, currentBook)
  }