package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type Citas struct {
	Descripcion  string `json:"Descripcion"`
	Especialidad string `json:"Especialidad"`
	Medico       string `json:"Medico"`
	Fecha        string `json:"Fecha"`
	Hora         string `json:"Hora"`
}

func DBconn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go_auth"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1)/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db

}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/index.html")
		if err != nil {
			panic(err)
		}
		w.Write(file)
	})

	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/table", VerCitas)
	fmt.Println("server: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	conexion1 := DBconn()
	// Parse the JSON payload
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad2 Request", http.StatusBadRequest)
		return
	}

	query := "SELECT id, username, user_type, password FROM users WHERE username = ?"
	row := conexion1.QueryRow(query, user.Username)

	var dbUser User
	err = row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.UserType, &dbUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Verify the password
	if user.Password != dbUser.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Authentication
	response := map[string]interface{}{
		"status":    "success",
		"user_type": dbUser.UserType,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	conexion2 := DBconn()
	var user2 User
	err := json.NewDecoder(r.Body).Decode(&user2)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Check if the username is already taken
	query := "SELECT id FROM users WHERE username = ?"
	row := conexion2.QueryRow(query, user2.Username)

	var existingID int
	err = row.Scan(&existingID)
	if err == nil {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	} else if err != sql.ErrNoRows {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Insert the new user into the database
	query = "INSERT INTO users(full_name,email,username,user_type,password) VALUES (?,?,?,?,?)"
	_, err = conexion2.Exec(query, user2.Fullname, user2.Email, user2.Username, user2.UserType, user2.Password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Registration successful
	response := map[string]interface{}{
		"status": "success",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func VerCitas(w http.ResponseWriter, r *http.Request) {
	con2 := DBconn()
	// Retrieve data from the database
	rows, err := con2.Query("SELECT Descripcion,Especialidad,Medico,Fecha,Hora from citas_creadas")
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved data
	var items []Citas

	// Loop through the query results and populate the slice
	for rows.Next() {
		var item Citas
		if err := rows.Scan(&item.Descripcion, &item.Especialidad, &item.Medico, &item.Fecha, &item.Hora); err != nil {
			http.Error(w, "Failed to scan data", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	// Convert data to JSON
	jsonData, err := json.Marshal(items)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	// Set the appropriate content type and send the JSON data to the frontend
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
