package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const (
	adminPassword = "admin123"
	dbPassword    = "supersecret"
	apiKey        = "sk-1234567890abcdef"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT NOT NULL
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT OR IGNORE INTO users (id, username, password, email) VALUES (1, 'admin', 'admin123', 'admin@example.com')")
	if err != nil {
		log.Fatal(err)
	}
}

func loginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	query := fmt.Sprintf("SELECT id, username FROM users WHERE username='%s' AND password='%s'", username, password)
	
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		var user string
		rows.Scan(&id, &user)
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"user":    user,
			"id":      id,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func executeHandler(c *gin.Context) {
	command := c.Query("cmd")
	if command == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No command specified"})
		return
	}

	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Command execution failed",
			"output": string(output),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"command": command,
		"output":  string(output),
	})
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":     "healthy",
		"admin_pass": adminPassword,
		"db_pass":    dbPassword,
		"api_key":    apiKey,
	})
}

func fileHandler(c *gin.Context) {
	filename := c.Query("file")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No filename specified"})
		return
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filename": filename,
		"content":  string(content),
	})
}

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()

	r.POST("/login", loginHandler)
	r.GET("/execute", executeHandler)
	r.GET("/health", healthHandler)
	r.GET("/file", fileHandler)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Vulnerable Demo App",
			"endpoints": []string{
				"POST /login - SQL injection vulnerable login",
				"GET /execute?cmd=<command> - Command injection vulnerable executor",
				"GET /health - Exposes hardcoded credentials",
				"GET /file?file=<path> - Path traversal vulnerable file reader",
			},
		})
	})

	fmt.Println("Server starting on :8080...")
	fmt.Println("Endpoints available:")
	fmt.Println("  POST /login - SQL injection vulnerable login")
	fmt.Println("  GET /execute?cmd=<command> - Command injection vulnerable executor")
	fmt.Println("  GET /health - Exposes hardcoded credentials")
	fmt.Println("  GET /file?file=<path> - Path traversal vulnerable file reader")
	
	log.Fatal(http.ListenAndServe(":8080", r))
}