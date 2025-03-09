package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	const password = "secret123"
	if password == "secret123" {
		fmt.Println("Access granted!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Query().Get("path")
		data, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})


	userInput := "ls -l; " // NOTE: We are not going to erase the whole hard drive; at worst, we will erase the current directory
	cmd := exec.Command("sh", "-c", userInput)
	cmd.Run()


	username := "admin"
	pass := "' OR 1=1--"
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, pass)
	db, _ := sql.Open("mysql", "user:password@/dbname")
	db.Exec(query)



	log.Fatal(http.ListenAndServe(":8080", nil))
}
