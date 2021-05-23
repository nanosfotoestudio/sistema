package handler

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
    "net/http"
)

type City struct {
	Id         int
	Name       string
	Population int
}

func Handler(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "admin:adminadmin@tcp(covid.cqjzeynatuei.us-east-2.rds.amazonaws.com:3306)/test")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM cities")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}
    var array []City
	for res.Next() {

		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}
        array = append(array, city)
	}
	w.Write(array)
}