package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

type User struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Image string `json:"image"`
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnUsers(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "username:passsword@tcp(127.0.0.1:3306)/belsql")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	users, err := db.Query("select * from voucher")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(users)
	var akhir []User
	for users.Next() {
        var user User
        // for each row, scan the result into our tag composite object
        err = users.Scan(&user.Id, &user.Name,&user.Image)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
				// and then print out the tag's Name attribute
		akhir = append(akhir,user)
    
	}
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(akhir)
}
func handleRequests() {
	fmt.Println("active in port 4000")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/coba", returnUsers)
    log.Fatal(http.ListenAndServe(":4000", nil))
}

func main() {

	

    handleRequests()
}