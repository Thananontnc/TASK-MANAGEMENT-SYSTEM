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
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Debug print to check the form values
		fmt.Println("Login attempt: username =", username, "email =", email, "password =", password)

		if username == "" || email == "" || password == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		validPassword, err := sqlhandle.CheckPwd(username, email)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Debug print to check the password match
		fmt.Println("Password from DB:", validPassword)

		if validPassword == password {
			http.Redirect(w, r, "/main-page", http.StatusSeeOther)
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Index TASK MANAGEMENT SYSTEM
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		filepath := filepath.Join("web", "index.html")
		http.ServeFile(w, r, filepath)
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		tasks, err := sqlhandle.GetTasks()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching tasks: %v", err), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("web/index.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, tasks)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		}
	}
}

// ADD NEW TASK

func AddTaskHandle(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method == http.MethodPost {
		// Get the task name from the form input
		taskName := r.FormValue("task")

		// Insert the task into the database
		if err := sqlhandle.InsertTask(taskName); err != nil {
			http.Error(w, fmt.Sprintf("Error adding task: %v", err), http.StatusInternalServerError)
			return
		}

		// Redirect to the main page after adding the task
		http.Redirect(w, r, "/main-page", http.StatusSeeOther)
	} else {
		// If the method is not POST, handle it with an error
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

type Task struct {
	TaskName string
	Status   string
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{TaskName: "Complete the project", Status: "In Progress"},
		{TaskName: "Write documentation", Status: "Completed"},
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("web/index.html") // Ensure correct path to your template
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Pass tasks data to the template
	err = tmpl.Execute(w, tasks)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
