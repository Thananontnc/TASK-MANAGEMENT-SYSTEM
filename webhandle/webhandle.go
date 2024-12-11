package webhandle

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"task-management-system/sqlhandle"
)

// Register
func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		filepath := filepath.Join("web", "register.html")
		http.ServeFile(w, r, filepath)
	} else if r.Method == http.MethodPost {

		r.ParseForm()

		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		phone := r.FormValue("phone")

		if username == "" || email == "" || password == "" || phone == "" {
			http.Error(w, "All are required", http.StatusBadRequest)
			return
		}
		err := sqlhandle.InsertRegister(username, email, password, phone)
		if err != nil {
			log.Fatal("Fail Insert Register data to user db : ", err)
		}
		fmt.Printf("Username: %s , Email: %s , Password: %s , Phone: %s", username, email, password, phone)
		fmt.Println("\nInsert Sucessfully ! ")
	}
}

// Login
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		filepath := filepath.Join("web", "login.html")
		http.ServeFile(w, r, filepath)
	} else if r.Method == http.MethodPost {

		r.ParseForm()

		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		valid, err := sqlhandle.CheckPwd(username, email)
		if err != nil {
			fmt.Println("Invalid username or password")
		}
		if valid == password {
			http.Redirect(w, r, "/tasks", http.StatusSeeOther)
		}

	}
}

// Index TASK MANAGEMENT SYSTEM
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		filepath := filepath.Join("web", "index.html")
		http.ServeFile(w, r, filepath)
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		taskName := r.FormValue("task")

		err := sqlhandle.InsertTask(taskName)
		if err != nil {
			fmt.Println("Fail to insert task")
		}
		fmt.Println("Add task successfully")
	}
}
