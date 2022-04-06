package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type information struct {
	orderid string
	name    string
	email   string
	pNumber string
}

func testScanfStuff() {
	// Declaring some variables
	var name string
	var alphabet_count int

	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Println("enter name:")
	fmt.Scanf("%s", &name)
	fmt.Println("enter alphabet count:")
	fmt.Scanf("%d", &alphabet_count)
	fmt.Println("input catpured.")

	// Printing the given texts
	fmt.Printf("The word %s containg %d number of alphabets.",
		name, alphabet_count)

}

func main() {

	// hello testScanfStuff()

	connStr := "user=postgres dbname=Orders password=Liverpool99  host=localhost sslmode=disable"
	//driver name part of "github.com/lib/pq"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	//check postgres to see if table exists
	var checkDatabase string
	db.QueryRow("SELECT to_regclass('public.orderapi')").Scan(&checkDatabase)
	if err != nil {
		fmt.Println(err)
	}
	//if table dose not exist then create one to use for this example
	if checkDatabase == "" {
		fmt.Println("Database Created")
		createSQL := "CREATE TABLE public.orderapi (pk SERIAL PRIMARY KEY,orderid character varying,name character varying,email character varying,pNumber character varying);"
		db.Query(createSQL)
	}

	//sql to insert employee information
	statement := "INSERT INTO public.orderapi(orderid, name, email, pNumber) VALUES($1, $2, $3, $4)"
	//prepare statement for sql
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	//call a instant of employee
	eName := information{}
	//allow 3 employee to be entered into database
	for i := 0; i < 3; i++ {
		fmt.Println("Order ID: ")
		//set fName of strut with text input
		//		fmt.Scanf("%s", &eName.orderid)
		fmt.Scanf("%s", &eName.orderid)

		//		fmt.Println("Your Name: ")
		//set fName of strut with text input
		fmt.Scanf("%s", &eName.name)

		//		fmt.Println("Your Email: ")
		fmt.Scanf("%s", &eName.email)

		//		fmt.Println("Your Phone Number: ")
		fmt.Scanf("%s", &eName.pNumber)

		fmt.Printf("The id %s name %s email %s number %s .",
			eName.orderid, eName.name, eName.email, eName.pNumber)

		//&eName.orderid = 1
		//&eName.name = "Liam"
		//&eName.email = "LiamEmail"
		//&eName.pNumber = 994

		//call prepared statement above
		fmt.Println("before insert:")
		stmt.QueryRow(eName.orderid, eName.name, eName.email, eName.pNumber)
		fmt.Println("after insert:")

	}
	//select employee first and last name
	rows, err := db.Query("Select orderid, name from public.orderapi")
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()

	fmt.Println("---------------------------------------------------------------------")
	//loop through all employee results
	for rows.Next() {
		//assign values to variables
		var orderid int
		var name string
		var email string
		var pNumber int
		err := rows.Scan(&orderid, &name, &email, &pNumber)
		if err != nil {
			fmt.Print(err)
		}
		//print results to console
		fmt.Printf("%s %s\n", orderid, name, email, pNumber)
	} //end of for loop
} //end of main function
