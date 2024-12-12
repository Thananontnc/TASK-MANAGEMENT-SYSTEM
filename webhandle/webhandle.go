package webhandle

import (
	"fmt"
	"html/template"
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
		// Get tasks from database
		tasks, err := sqlhandle.GetTasks()
		if err != nil {
			http.Error(w, "Unable to load tasks", http.StatusInternalServerError)
			return
		}

		// Parse the HTML template
		tmpl, err := template.ParseFiles("web/index.html")
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
		}

		// Execute the template with the tasks data
		err = tmpl.Execute(w, tasks)
		if err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		// Handle task insertion
		r.ParseForm()
		taskName := r.FormValue("task")

		err := sqlhandle.InsertTask(taskName)
		if err != nil {
			http.Error(w, "Failed to insert task", http.StatusInternalServerError)
			return
		}
		// Redirect back to the task list page after adding
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
