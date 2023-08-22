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
	DNI      string `json:"dni"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Tel      string `json:"telefono"`
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
	Id           string `json:"Id"`
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
	r.HandleFunc("/table1", VerCitas_doc)
	r.HandleFunc("/table2", VerCitas_doc2)
	r.HandleFunc("/list-options", handleListOptions)
	r.HandleFunc("/get-data", getdata)
	r.HandleFunc("/horario", horario)
	r.HandleFunc("/Reservar", ReservarHandler)
	r.HandleFunc("/usuario", VerUsuario)
	r.HandleFunc("/Eliminar", EliminarHandler)
	r.HandleFunc("/Actualizar", ActualizarHandler)
	r.HandleFunc("/Atender", AtenderHandler)
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
	query = "INSERT INTO users(full_name,DNI,email,username,telefono,user_type,password) VALUES (?,?,?,?,?,?,?)"
	_, err = conexion2.Exec(query, user2.Fullname, user2.DNI, user2.Email, user2.Username, user2.Tel, user2.UserType, user2.Password)
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
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	usuario := data["paciente"]

	query := "SELECT id FROM users WHERE username = ?"
	row1 := con2.QueryRow(query, usuario)

	var existingID int
	err = row1.Scan(&existingID)

	quer := "SELECT citas_creadas.Id,Descripcion, Especialidad,Nombre,Fecha,Hora,citas_creadas.Estado from citas_creadas INNER JOIN medicos ON citas_creadas.id_medico = medicos.id where id_paciente = ?"
	rows, err := con2.Query(quer, existingID)
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	var items []Citas

	for rows.Next() {
		var item Citas
		if err := rows.Scan(&item.Id, &item.Descripcion, &item.Especialidad, &item.Medico, &item.Fecha, &item.Hora, &item.Estado); err != nil {
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

	// Insert the new date into the database
	query = "INSERT INTO citas_creadas ( id_paciente, Descripcion, Especialidad, id_medico, Fecha, Hora, Estado) VALUES (?,?,?,?,?,?,?)"
	_, err = conexi.Exec(query, existingID, des, nombre1, medico, fec, H, estado)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Reservation successful
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

func VerUsuario(w http.ResponseWriter, r *http.Request) {
	con2 := DBconn()
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	usuario := data["paciente"]
	fmt.Println(usuario)

	query := "SELECT * FROM users WHERE username = ?"
	row1 := con2.QueryRow(query, usuario)

	var user User
	err = row1.Scan(&user.ID, &user.Fullname, &user.DNI, &user.Email, &user.Username, &user.Tel, &user.UserType, &user.Password)
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func EliminarHandler(w http.ResponseWriter, r *http.Request) {
	conexi := DBconn()
	var cit map[string]string
	err := json.NewDecoder(r.Body).Decode(&cit)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	Id_C := cit["Id"]
	fec := cit["Fecha"]
	H := cit["Hora"]

	fmt.Println(Id_C, fec, H)

	q1 := "SELECT id_medico FROM `citas_creadas` WHERE Id =? "
	roww2 := conexi.QueryRow(q1, Id_C)
	var id_medico int
	err = roww2.Scan(&id_medico)

	query2 := "SELECT id FROM horario WHERE hora_inicio = ? AND id_medico = ? "
	roww := conexi.QueryRow(query2, H, id_medico)
	var id_hora int
	err = roww.Scan(&id_hora)
	fmt.Println(id_hora)

	quer := "DELETE from dia  WHERE dia = ? and id_horario = ?"
	_, err = conexi.Exec(quer, fec, id_hora)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	quer3 := "delete from citas_creadas where Id = ?"
	_, err = conexi.Exec(quer3, Id_C)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"status": "cita eliminada",
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

func ActualizarHandler(w http.ResponseWriter, r *http.Request) {
	conexi := DBconn()
	var cit map[string]string
	err := json.NewDecoder(r.Body).Decode(&cit)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	nombre := cit["fullname"]
	correo := cit["email"]
	dni := cit["dni"]
	users := cit["username"]
	pass := cit["password"]
	telef := cit["telefono"]
	u := cit["user"]

	fmt.Println(nombre, correo, dni, users, pass, telef, u)

	quer := "UPDATE users SET  full_name = ?, email = ?, username = ?, DNI = ?, telefono = ?, password = ?  where username = ?"
	_, err = conexi.Exec(quer, nombre, correo, users, dni, telef, pass, u)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// perfil editado
	response := map[string]interface{}{
		"status": "Perfil editado",
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

func VerCitas_doc(w http.ResponseWriter, r *http.Request) {
	con2 := DBconn()
	// Retrieve data from the database
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	usuario := data["paciente"]
	fmt.Println(usuario)
	query := "SELECT id FROM medicos WHERE username = ?"
	row1 := con2.QueryRow(query, usuario)

	var existingID int
	err = row1.Scan(&existingID)
	fmt.Println(existingID)

	quer := "select citas_creadas.Id,citas_creadas.Descripcion,users.full_name,citas_creadas.Fecha,citas_creadas.Hora,citas_creadas.Estado from citas_creadas INNER JOIN users on citas_creadas.id_paciente = users.id where id_medico=?"
	rows, err := con2.Query(quer, existingID)
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var items []Citas

	for rows.Next() {
		var item Citas
		if err := rows.Scan(&item.Id, &item.Descripcion, &item.Especialidad, &item.Fecha, &item.Hora, &item.Estado); err != nil {
			http.Error(w, "Failed to scan data", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}
	fmt.Println(items)

	jsonData, err := json.Marshal(items)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	// Set the appropriate content type and send the+ JSON data to the frontend
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func AtenderHandler(w http.ResponseWriter, r *http.Request) {
	conexi := DBconn()
	var cit map[string]string
	err := json.NewDecoder(r.Body).Decode(&cit)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	id := cit["Id"]

	fmt.Println(id)

	quer := "UPDATE citas_creadas SET  Estado = 'Atendido'  where Id = ?"
	_, err = conexi.Exec(quer, id)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Registration successful
	response := map[string]interface{}{
		"status": "cita atendida",
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
func VerCitas_doc2(w http.ResponseWriter, r *http.Request) {
	con2 := DBconn()
	// Retrieve data from the database
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	usuario := data["paciente"]
	fmt.Println(usuario)
	query := "SELECT id FROM medicos WHERE username = ?"
	row1 := con2.QueryRow(query, usuario)

	var existingID int
	err = row1.Scan(&existingID)
	fmt.Println(existingID)

	quer := "select citas_creadas.Id,citas_creadas.Descripcion,users.full_name,citas_creadas.Fecha,citas_creadas.Hora,citas_creadas.Estado from citas_creadas INNER JOIN users on citas_creadas.id_paciente = users.id where id_medico=? and citas_creadas.Estado='Pendiente' "
	rows, err := con2.Query(quer, existingID)
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var items []Citas

	for rows.Next() {
		var item Citas
		if err := rows.Scan(&item.Id, &item.Descripcion, &item.Especialidad, &item.Fecha, &item.Hora, &item.Estado); err != nil {
			http.Error(w, "Failed to scan data", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}
	fmt.Println(items)

	jsonData, err := json.Marshal(items)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
