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
		// Get tasks from the database
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
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
	}
}

// DeleteTask handles the deletion of a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL
	taskID := r.URL.Path[len("/delete/"):]

	// Call the function to delete the task from the database
	err := sqlhandle.DeleteTask(taskID)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	// Redirect back to the task list after deleting the task
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// CompleteTask handles marking a task as completed
func CompleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL
	taskID := r.URL.Path[len("/complete/"):]

	// Call the function to update the task status to "Completed"
	err := sqlhandle.CompleteTask(taskID)
	if err != nil {
		http.Error(w, "Failed to complete task", http.StatusInternalServerError)
		return
	}

	// Redirect back to the task list after marking the task as completed
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// AddTaskHandle handles the submission of a new task from the form
func AddTaskHandle(w http.ResponseWriter, r *http.Request) {
	// Only handle POST requests
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}

		// Get the task name from the form
		taskName := r.FormValue("task")

		// Call InsertTask to add the task to the database
		err = sqlhandle.InsertTask(taskName)
		if err != nil {
			http.Error(w, "Failed to insert task", http.StatusInternalServerError)
			return
		}

		// Redirect back to the task list page after adding the task
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
	} else {
		// If the method is not POST, show an error
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
