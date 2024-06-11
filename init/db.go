package init

import (
	// "context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	sqlc"server/sql/database" // Adjust this import path as per your project structure

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// Config holds the configuration for database connection
type Config struct {
    Driver      string
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}


// LoadConfig loads database configuration from environment variables or another source
func LoadConfig() Config {
// DB_URL=postgres://postgres:root@localhost:5432/attempt2?sslmode=disable

    return Config{
        Driver:   getEnv("DB_DRIVER", "postgres"),
        Host:     getEnv("DB_HOST", "postgres"),
        Port:     getEnvInt("DB_PORT", 5432),
        User:     getEnv("DB_USER", "postgres"),
        Password: getEnv("DB_PASSWORD", "root"),
        DBName:   getEnv("DB_NAME", "attempt2"),
        SSLMode:  getEnv("DB_SSLMODE", "disable"),
    }
}

// getEnv retrieves environment variables or returns a default value
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

// getEnvInt retrieves integer environment variables or returns a default value
func getEnvInt(key string, defaultValue int) int {
    if value, exists := os.LookupEnv(key); exists {
        intValue, err := strconv.Atoi(value)
        if err == nil {
            return intValue
        }
    }
    return defaultValue
}

// Global variables to hold the database and queries
var (
    DB      *sql.DB
    Queries *sqlc.Queries
)

// ConnectDB initializes the database connection and sqlc Queries
func ConnectDB() error {
    config := LoadConfig()

    // driver://user:password@host:port/database
    dsn := fmt.Sprintf(
        "%s://%s:%s@%s:%d/%s?sslmode=%s",
        config.Driver, config.User, config.Password, config.Host, config.Port, config.DBName, config.SSLMode,
    )


    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return fmt.Errorf("could not open db: %w", err)
    }

    if err := db.Ping(); err != nil {
        return fmt.Errorf("could not connect to db: %w", err)
    }

    DB = db
    Queries = sqlc.New(DB) // Initialize sqlc Queries with the database connection

    log.Println("Successfully connected to the database")
    return nil
}

// DisconnectDB closes the database connection
func DisconnectDB() {
    if DB != nil {
        DB.Close()
        log.Println("Database connection closed")
    }
}
