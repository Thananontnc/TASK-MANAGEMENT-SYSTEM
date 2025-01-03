package main

import (
	"fmt"
	"net/http"
	"task-management-system/csshandle"
	"task-management-system/sqlhandle"
	"task-management-system/webhandle"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Run DataBase
	sqlhandle.ConnectToDB("root", "root", "localhost", "user_db")

	// Handle HTML FILE
	http.HandleFunc("/register", webhandle.RegisterHandle)
	http.HandleFunc("/login", webhandle.LoginHandle)
	http.HandleFunc("/tasks", webhandle.IndexHandle)
	http.HandleFunc("/delete/", webhandle.DeleteTask) // Route for deleting a task
	http.HandleFunc("/complete/", webhandle.CompleteTask)

	// Handle CSS FILE
	http.HandleFunc("/register-style", csshandle.RegisterHandleCSS)
	http.HandleFunc("/login-style", csshandle.LoginHandleCss)
	http.HandleFunc("/tasks-style", csshandle.IndexHandleCSS)

	defer sqlhandle.CloseDB()

	fmt.Println("Run at port http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
