package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {

	type User struct {
		Username string
		Email    string
	}

	type Data struct {
		Users []User
	}

	UserData := Data{
		Users: []User{
			{Username: "John Doe", Email: "john@gmail.com"},
			{Username: "Peter Parker", Email: "spidy@marvel.com"},
		},
	}

	// Handle requests to localhost:8000/
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == "POST" {
			userName := request.FormValue("username")
			userEmail := request.FormValue("email")
			UserData.Users = append(UserData.Users, User{userName, userEmail})
			htmlResponse := fmt.Sprintf("<li class='p-3 rounded bg-white'>%s - %s</li>", userName, userEmail)
			io.WriteString(response, htmlResponse)
		}
		if request.Method == "GET" {
			t, err := template.ParseFiles("index.html")
			if err != nil {
				log.Fatal(err)
			}

			err = t.Execute(response, UserData)
			if err != nil {
				log.Fatal(err)
			}
		}

	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
