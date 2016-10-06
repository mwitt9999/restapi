package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ApiDocumentation)
	router.HandleFunc("/users", UsersShowAll)
	router.HandleFunc("/users/{userId}", UserShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func ApiDocumentation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our API Docs!")
}

func UsersShowAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users All")
}

func UserShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["userId"]

	if i, err := strconv.Atoi(userId); err == nil {
		fmt.Printf("i=%d, type: %T\n", i, i)

		h := md5.New()
		io.WriteString(h, "password123")

		pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

		user := Users{
			User{UserId: i, Username: "mwitt8178", Password: pwmd5, Name: "Matthew Witt"},
		}

		json.NewEncoder(w).Encode(user)
	}
}
