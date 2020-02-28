package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	database = "system"
	user     = "arjom"
	password = "12345678"
)

// Usuario do banco
type Usuario struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func openConn() string {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, database)
	return connectionString

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// LoginHandler deve executar o login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		carregaLogin(w, r)
	case r.Method == "POST":
		login(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :")
	}

}

func carregaLogin(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./public/login.html")
	t, _ := template.ParseFiles("./public/login.html")
	t.Execute(w, nil)

}

func login(w http.ResponseWriter, r *http.Request) {
	// Inicializando conexão
	db, err := sql.Open("mysql", openConn())
	checkError(err)
	defer db.Close()
	log.Print("login")

	// Pegando informações
	r.ParseForm()
	username := strings.Join(r.Form["username"], "")
	password := strings.Join(r.Form["password"], "")
	password = getMD5Hash(password)

	var u Usuario
	db.QueryRow(`
	select id, username, password 
	from usuarios 
	where username = ? and password = ?
	`, username, password).Scan(&u.ID, &u.Username, &u.Password)

	// Verificando se a consulta retornou valor
	if u.ID != 0 {
		log.Print("Usuário autenticado")
	} else {
		log.Print("Usuário não cadastrado")
	}

	// Carregando o login ao final da operação
	carregaLogin(w, r)
}

// RegisterHandler deve executar o login
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		carregaRegister(w, r)
	case r.Method == "POST":
		register(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :")
	}
}

func carregaRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/register.html")
}

func register(w http.ResponseWriter, r *http.Request) {
	// Inicializando conexão
	db, err := sql.Open("mysql", openConn())
	checkError(err)
	defer db.Close()
	log.Print("register")

	// Pegando informações
	r.ParseForm()
	log.Print("username:", r.Form["username"])
	log.Print("password:", r.Form["password"])

	username := strings.Join(r.Form["username"], "")
	password := strings.Join(r.Form["password"], "")
	password = getMD5Hash(password)

	// Inserindo no banco
	stmt, _ := db.Prepare("insert into usuarios(username, password) values(?,?)")
	_, insertErr := stmt.Exec(username, password)
	checkError(insertErr)
	log.Print("Inserido com sucesso")
}
