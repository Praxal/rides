package main

import (
	"fmt"
	"net/http"
	"os"
	db "app/postgres"
)

func getData(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}
func getDrivers(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Connection.Query("SELECT name FROM drivers")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	data := ""
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(name)
		data += fmt.Sprintf("%s ", name)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, data)
}
func main() {
   	db.InitDB()
	defer db.Connection.Close() 
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/data", getData)
	http.HandleFunc("/drivers", getDrivers)

	serverEnv := os.Getenv("SERVER_ENV")

	if serverEnv == "DEV" {
		http.ListenAndServe(":8080", nil)
	} else if serverEnv == "PROD" {
		http.ListenAndServeTLS(
			":443",
			"/etc/letsencrypt/live/app.p4family.com/fullchain.pem",
			"/etc/letsencrypt/live/app.p4family.com/privkey.pem",
			nil,
		)
	}
}
