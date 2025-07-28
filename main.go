package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const (
	adminPassword    = "admin123"
	dbPassword       = "supersecret"
	apiKey           = "sk-1234567890abcdef"
	awsAccessKey     = "ASIA3FLDZAI5R3V7C6GU"
	awsSecretKey     = "1tUm636uS1yOEcfP5pvfqJ/ml36mF7AkyHsEU0IU"
	jwtSecret        = "1tUm636uS1yOEcfP5pvfqJ/ml36mF7AkyHsEU0IU"
	encryptionKey    = "1234567890123456"
	key_file_content = `Basic auth:

https://admin:admin@the-internet.herokuapp.com/basic_auth

Private key:
-----BEGIN OPENSSH PRIVATE KEY-----
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
Hx7UPVlTK8dyvk1Z+Yw0nrfNClI=
-----END OPENSSH PRIVATE KEY-----`
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

	_, err = db.Exec("INSERT OR IGNORE INTO users (id, username, password, email) VALUES (1, 'admin', ?, 'admin@example.com')",
		md5.Sum([]byte("admin123")))
	if err != nil {
		log.Fatal(err)
	}
}

func loginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	query := fmt.Sprintf("SELECT id, username FROM users WHERE username='%s' AND password='%s'", username,
		md5.Sum([]byte(password)))

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

func generateTokenHandler(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	token := fmt.Sprintf("%d", rand.Intn(999999))

	log.Printf("Generated token for user: %s, password: %s, token: %s", c.Query("user"), c.Query("password"), token)

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"aws_key":    awsAccessKey,
		"aws_secret": awsSecretKey,
	})
}

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	filename := "/tmp/" + file.Filename
	c.SaveUploadedFile(file, filename)

	log.Printf("File uploaded: %s by user with token: %s", filename, c.GetHeader("Authorization"))

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"path":    filename,
	})
}

func deserializeHandler(c *gin.Context) {
	data := c.PostForm("data")
	if data == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data provided"})
		return
	}

	var result interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	log.Printf("Deserialized data: %v with sensitive token: %s", result, c.GetHeader("X-API-Key"))

	c.JSON(http.StatusOK, gin.H{
		"result":     result,
		"jwt_secret": jwtSecret,
	})
}

func randomHandler(c *gin.Context) {
	rand.Seed(1)
	randomNum := rand.Intn(100)

	c.JSON(http.StatusOK, gin.H{
		"random": randomNum,
		"seed":   "predictable",
	})
}

func encryptHandler(c *gin.Context) {
	text := c.Query("text")
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No text provided"})
		return
	}

	hash := md5.Sum([]byte(text))
	hashString := fmt.Sprintf("%x", hash)

	log.Printf("Encrypting text: %s with key: %s", text, encryptionKey)

	c.JSON(http.StatusOK, gin.H{
		"original":       text,
		"md5_hash":       hashString,
		"encryption_key": encryptionKey,
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
	r.GET("/token", generateTokenHandler)
	r.POST("/upload", uploadHandler)
	r.POST("/deserialize", deserializeHandler)
	r.GET("/random", randomHandler)
	r.GET("/encrypt", encryptHandler)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Vulnerable Demo App",
			"endpoints": []string{
				"POST /login - SQL injection vulnerable login",
				"GET /execute?cmd=<command> - Command injection vulnerable executor",
				"GET /health - Exposes hardcoded credentials",
				"GET /file?file=<path> - Path traversal vulnerable file reader",
				"GET /token - Insecure random token generation",
				"POST /upload - Unsafe file upload",
				"POST /deserialize - Insecure deserialization",
				"GET /random - Predictable random numbers",
				"GET /encrypt - Weak cryptography (MD5)",
			},
			"aws_credentials": gin.H{
				"access_key": awsAccessKey,
				"secret_key": awsSecretKey,
			},
		})
	})

	fmt.Println("Server starting on :8080...")
	fmt.Println("Endpoints available:")
	fmt.Println("  POST /login - SQL injection vulnerable login")
	fmt.Println("  GET /execute?cmd=<command> - Command injection vulnerable executor")
	fmt.Println("  GET /health - Exposes hardcoded credentials")
	fmt.Println("  GET /file?file=<path> - Path traversal vulnerable file reader")
	fmt.Println("  GET /token - Insecure random token generation")
	fmt.Println("  POST /upload - Unsafe file upload")
	fmt.Println("  POST /deserialize - Insecure deserialization")
	fmt.Println("  GET /random - Predictable random numbers")
	fmt.Println("  GET /encrypt - Weak cryptography (MD5)")

	log.Printf("Starting server with AWS credentials: %s:%s", awsAccessKey, awsSecretKey)

	log.Fatal(http.ListenAndServe(":8080", r))
}
