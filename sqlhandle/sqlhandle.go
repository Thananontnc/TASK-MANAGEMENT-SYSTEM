package sqlhandle

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect TO Database
func ConnectToDB(username, password, hostname, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbname)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	// Verify the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Database connection established successfully!")
	return nil
}

// Close Database
func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database connection Closed")
	}
}

// Insert register data
func InsertRegister(username, email, password, phone string) error {
	query := "INSERT INTO users (username,email,password,phone) VALUES (?,?,?,?)"
	_, err := DB.Exec(query, username, email, password, phone)
	if err != nil {
		return fmt.Errorf("fail insert register data to users table: %v", err)
	}
	return nil
}

// Password Checking
func CheckPwd(username, email string) (string, error) {
	var password string
	query := "SELECT password FROM users WHERE username = ? AND email = ?"
	err := DB.QueryRow(query, username, email).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found with username '%s' and email '%s'", username, email)
		}
		return "", fmt.Errorf("error checking password: %v", err)
	}
	return password, nil
}

// Insert task to database
func InsertTask(taskName string) error {
	_, err := DB.Exec("INSERT INTO tasks (task_name,status) VALUES (?,?)", taskName, "Pending")
	if err != nil {
		log.Printf("Error inserting task: %v", err)
		return err
	}
	return nil
}

// GET TASKS FUNCTION
func GetTasks() ([]Task, error) {
	rows, err := DB.Query("SELECT id, task_name, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.TaskName, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Task represents a task in the task management system
type Task struct {
	ID       int
	TaskName string
	Status   string
}
