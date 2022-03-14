package database

import (
	"database/sql"
	"fmt"

	"github.com/pashamakhilkumarreddy/golang-postgres-crud/helpers"
)

func ConnectToDB() *sql.DB {
	connstring := "host=localhost port=5432 user=admin database=admin password=n9pzYyojp3zcU+QTMfeEaA== sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		helpers.LogError(err)
	}
	return db
}

func InsertIntoDB(db *sql.DB, rollno int, name, course string) {
	_, err := db.Exec("INSERT into student(roll_no, name, course) VALUES($1, $2, $3)", rollno, name, course)
	if err != nil {
		helpers.LogError(err)
	}
	fmt.Println("Successfully inserted the record into the student table")
}

func ReadFromDB(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		helpers.LogError(err)
	}
	fmt.Println("Successfully finished reading from the DB")
	return rows
}

func UpdateData(db *sql.DB, rollno int, name, course string) {
	_, err := db.Query("UPDATE student SET name=$1, course=$2 WHERE roll_no=$3", name, course, rollno)
	if err != nil {
		helpers.LogError(err)
	}
	fmt.Println("Successfully updated the record")
}

func DeleteRecord(db *sql.DB, rollno int) {
	res, err := db.Query("SELECT * FROM student WHERE roll_no=$1", rollno)
	if err != nil {
		helpers.LogError(err)
	}
	if res != nil {
		_, err := db.Query("DELETE FROM student WHERE roll_no=$1", rollno)
		if err != nil {
			helpers.LogError(err)
		}
	}
	fmt.Println("Successfully deleted the record")
}
