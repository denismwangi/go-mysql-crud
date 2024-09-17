package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const dbUser = "root"
const dbPassword = ""
const dbName = "goCrud"
const dbPort = "127.0.0.1:3306"


func ListMoviesHandler() []Movie{
	db, err := sql.Open("mysql", dbUser+"@tcp("+dbPort+")/"+dbName)

	if err != nil{
		fmt.Println("Error", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM movies")

	if err != nil{
		fmt.Println("Error", err.Error())
		return nil
	}

	movieItems := []Movie{}
    for results.Next(){
		var movieFound Movie
		err = results.Scan(&movieFound.ID,&movieFound.Title,&movieFound.Director)

		if err != nil{
			panic(err.Error())
		}
		if err != nil {
            panic(err.Error())
        }
		movieItems = append(movieItems,movieFound)
	}

	return movieItems

}

func CreateMovieHandler(movieItem Movie){
	db, err := sql.Open("mysql", dbUser+"@tcp("+dbPort+")/"+dbName)

	if err != nil {
		fmt.Println("Error",err.Error())
	}

	defer db.Close()

	insert ,err := db.Query("INSERT INTO movies (id,title,director) values(?,?,?)", movieItem.ID, movieItem.Title, movieItem.Director)

	if err != nil{
		fmt.Println("Error", err.Error())
	}

	defer insert.Close()
}

func GetMovieById(id string) *Movie{
	db, err := sql.Open("mysql", dbUser+"@tcp("+dbPort+")/"+dbName)
	movieItem := &Movie{}

	if err != nil {
		fmt.Println("Error",err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT *FROM movies WHERE id=",id)
	if err != nil {
		fmt.Println("Error",err.Error())
		return nil
	}

	if results.Next() {
		err := results.Scan(&movieItem.ID, &movieItem.Title, &movieItem.Director)
		if err != nil {
			return nil
		}
	}

	return movieItem
}