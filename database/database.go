package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func Init() {
    // Load file .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, using environment variables")
    }

    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Test connection
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    log.Println("Database connection established")
}

