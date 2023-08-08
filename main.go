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
type Medico struct {
	ID     int    `json:"ID"`
	Nombre string `json:"Nombre"`
}
type Fecha struct {
	ID    int    `json:"id"`
	Fecha string `json:"fecha"`
}
type horarioo struct {
	Horainicio string `json:"hora_inicio"`
	Horasalida string `json:"hora_salida"`
}

type Citas struct {
	Descripcion  string `json:"Descripcion"`
	Especialidad string `json:"Especialidad"`
	Medico       string `json:"Medico"`
	Fecha        string `json:"Fecha"`
	Hora         string `json:"Hora"`
	Estado       string `json:"Estado"`
}

type Reserva struct {
	Paciente     string `json:"Paciente"`
	Descripcion  string `json:"Descripcion"`
	Especialidad string `json:"Especialidad"`
	Medico       string `json:"Medico"`
	Fecha        string `json:"Fecha"`
	Hora         string `json:"Hora"`
}

type Especialidad struct {
	ID     int    `json:"ID"`
	Nombre string `json:"Nombre"`
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
	r.HandleFunc("/list-options", handleListOptions)
	r.HandleFunc("/get-data", getdata)
	r.HandleFunc("/horario", horario)
	r.HandleFunc("/Reservar", ReservarHandler)
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

	rows, err := con2.Query("SELECT Descripcion, Especialidad,Nombre,Fecha,Hora,Estado from citas_creadas INNER JOIN medicos ON citas_creadas.id_medico = medicos.id")
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
		if err := rows.Scan(&item.Descripcion, &item.Especialidad, &item.Medico, &item.Fecha, &item.Hora, &item.Estado); err != nil {
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

	// Set the appropriate content type and send the+ JSON data to the frontend
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handleListOptions(w http.ResponseWriter, r *http.Request) {
	db := DBconn()

	rows, err := db.Query("SELECT Id_especialidad,Nombre FROM especialidades")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	especialidades := []Especialidad{}
	for rows.Next() {
		var esp Especialidad
		err := rows.Scan(&esp.ID, &esp.Nombre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		especialidades = append(especialidades, esp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(especialidades)
}

func getdata(w http.ResponseWriter, r *http.Request) {
	conexion2 := DBconn()
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	selectedOption := data["option"]
	// Check if the username is already taken
	query := "SELECT id,Nombre FROM medicos WHERE id_especialidad = ?"
	rows, err := conexion2.Query(query, selectedOption)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	medicos := []Medico{}
	for rows.Next() {
		var med Medico
		err := rows.Scan(&med.ID, &med.Nombre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		medicos = append(medicos, med)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicos)

}

func horario(w http.ResponseWriter, r *http.Request) {
	conexion5 := DBconn()
	var datos map[string]string
	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	id_med := datos["id"]
	fecha_cita := datos["fecha"]
	// Check if the username is already taken
	query1 := "SELECT hora_inicio,hora_salida FROM horario WHERE not Id IN(select id_horario from dia where Dia = ?) AND id_medico = ?;"
	rows1, err := conexion5.Query(query1, fecha_cita, id_med)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows1.Close()

	hora := []horarioo{}
	for rows1.Next() {
		var hor horarioo
		err := rows1.Scan(&hor.Horainicio, &hor.Horasalida)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		hora = append(hora, hor)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hora)

}

func recibirFecha(w http.ResponseWriter, r *http.Request) {
	var fecha Fecha
	err := json.NewDecoder(r.Body).Decode(&fecha)
	if err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	// Puedes enviar una respuesta al cliente si es necesario
	respuesta := map[string]string{"mensaje": "Fecha recibida correctamente"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func ReservarHandler(w http.ResponseWriter, r *http.Request) {
	conexi := DBconn()
	var cit map[string]string
	err := json.NewDecoder(r.Body).Decode(&cit)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	des := cit["Descripcion"]
	medico := cit["Medico"]
	fec := cit["Fecha"]
	H := cit["Hora"]
	pacien := cit["Paciente"]
	espec := cit["Especialidad"]

	fmt.Println(des, medico, fec, H, pacien, espec)

	query2 := "SELECT id FROM horario WHERE hora_inicio = ? AND id_medico = ? "
	roww := conexi.QueryRow(query2, H, medico)
	var id_hora int
	err = roww.Scan(&id_hora)
	fmt.Println(id_hora)
	quer := "INSERT INTO dia(id_horario,Dia) VALUES (?,?)"
	_, err = conexi.Exec(quer, id_hora, fec)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	query := "SELECT id FROM users WHERE username = ?"
	row := conexi.QueryRow(query, pacien)

	query1 := "SELECT Nombre FROM especialidades  WHERE  Id_especialidad= ?"
	row1 := conexi.QueryRow(query1, espec)

	var estado = "Pendiente"

	var existingID int
	err = row.Scan(&existingID)

	var nombre1 string
	err = row1.Scan(&nombre1)
	fmt.Println(existingID, nombre1, estado)

	// Insert the new user into the database
	query = "INSERT INTO citas_creadas ( id_paciente, Descripcion, Especialidad, id_medico, Fecha, Hora, Estado) VALUES (?,?,?,?,?,?,?)"
	_, err = conexi.Exec(query, existingID, des, nombre1, medico, fec, H, estado)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Registration successful
	response := map[string]interface{}{
		"status": "cita reservada",
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
