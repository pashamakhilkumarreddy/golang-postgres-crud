package main

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/pashamakhilkumarreddy/golang-postgres-crud/database"
	"github.com/pashamakhilkumarreddy/golang-postgres-crud/helpers"
)

func main() {
	fmt.Println("Welcome to Golang Postgres CRUD")
	var choice string
	db := database.ConnectToDB()
	for {
		fmt.Printf("\nPlease enter your choice\n")
		fmt.Printf("1. Insert data in Postgres DB \n")
		fmt.Printf("2. Read data from Postgres DB \n")
		fmt.Printf("3. Update data in Postgres DB \n")
		fmt.Printf("4. Delete data from Postgres DB \n")
		fmt.Printf("5. Exit \n")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var rollno int
			var name, course string
			fmt.Println("Please enter the Roll no")
			rollno, _ = strconv.Atoi(helpers.GetInput())
			fmt.Println("Please enter the Name")
			name = helpers.GetInput()
			fmt.Println("Please enter the course")
			course = helpers.GetInput()
			database.InsertIntoDB(db, rollno, name, course)
		case "2":
			var roll_no int
			var name, course string
			rows := database.ReadFromDB(db)
			fmt.Println("Rollno\tName\tCourse")
			fmt.Println("-----------------------")
			for rows.Next() {
				rows.Scan(&roll_no, &name, &course)
				fmt.Printf("%d\t%s\t%s\n", roll_no, name, course)
			}
			fmt.Println("-----------------------")
		case "3":
			fmt.Println("Please enter the rollno of the student")
			rollno, _ := strconv.Atoi(helpers.GetInput())
			fmt.Println("Please enter the name to update")
			name := helpers.GetInput()
			fmt.Println("Please enter the course to update")
			course := helpers.GetInput()
			database.UpdateData(db, rollno, name, course)
		case "4":
			fmt.Println("Please enter the Roll No of the student to delete")
			rollno, _ := strconv.Atoi(helpers.GetInput())
			database.DeleteRecord(db, rollno)
		case "5":
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}
}
