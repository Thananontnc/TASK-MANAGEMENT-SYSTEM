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
	sqlhandle.ConnectToDB("root", "Plai1412", "localhost", "user_db")

	// Handle HTML FILE
	http.HandleFunc("/register", webhandle.RegisterHandle)
	http.HandleFunc("/login", webhandle.LoginHandle)
	http.HandleFunc("/tasks", webhandle.AddTaskHandle)
	http.HandleFunc("/main-page", webhandle.IndexHandle)
	http.HandleFunc("/task", webhandle.TaskHandler)
	// Handle CSS FILE
	http.HandleFunc("/register-style", csshandle.RegisterHandleCSS)
	http.HandleFunc("/login-style", csshandle.LoginHandleCss)
	http.HandleFunc("/main-style", csshandle.IndexHandleCSS)

	defer sqlhandle.CloseDB()

	fmt.Println("Run at port http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
