package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuario (:)
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e encaminha
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userByID(w, r, id)
	case r.Method == "GET":
		userAll(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry")
	}
}

func userByID(w http.ResponseWriter, r *http.Request, id int) {
	// dont do that
	db, err := sql.Open("mysql", "lucas:123456@tcp(localhost:33061)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&user.ID, &user.Nome)

	json, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func userAll(w http.ResponseWriter, r *http.Request) {
	// dont do that
	db, err := sql.Open("mysql", "lucas:123456@tcp(localhost:33061)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var users []Usuario
	for rows.Next() {
		var user Usuario
		rows.Scan(&user.ID, &user.Nome)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
