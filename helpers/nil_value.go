package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"


	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name *string
	Age  *int
}

func main() {
	/*http.HandleFunc("/", handleForm)
	http.HandleFunc("/submit", handleSubmit)
	http.ListenAndServe(":8080", nil)*/

	var n1 *int32 = nil
	var n2 int32 = 45

	println("n1: ", Int32OrNilToString(n1))
	println("n2: ", Int32OrNilToString(&n2))
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	tmpl := `<html>
	<head><title>Submit User Form</title></head>
	<body>
		<form action="/submit" method="post">
			Name: <input type="text" name="name"><br>
			Age: <input type="text" name="age"><br>
			<input type="submit" value="Submit">
		</form>
	</body>
	</html>`
	fmt.Fprintf(w, tmpl)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.Form.Get("name")
	ageStr := r.Form.Get("age")

	var user User
	user.Name = stringPointerOrNil(name)
	user.Age = intPointerOrNil(ageStr)

	// Insert user into database
	err := insertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User submitted successfully!")
}

func stringPointerOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func intPointerOrNil(s string) *int {
	if s == "" {
		return nil
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &i
}

func IntOrNilToString(number *int) string {
	if number == nil {
		return ""
	}
	return strconv.Itoa(*number)
}


// Int32 or Nil -> string
func Int32OrNilToString(number *int32) string {
	if number == nil {
		return ""
	}
	return strconv.Itoa(int(*number))
}


func insertUser(user User) error {
	db, err := sql.Open("mysql", "root:<password>@tcp(127.0.0.1:3306)/test_nulls")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		return err
	}

	return nil
}
