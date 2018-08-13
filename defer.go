/*
  The defer statement allows you to designate specified functions to be executed just before returning from the current function block
  In the example below, we are going to pretend that we are doing some database operations: we make a connection and then attempt some
  database work. Then we act as if the database operations have crashed and return immediately. This would have left the database
  connection still dangling, but since we had already called the database disconnect function with a defer, this gets called now, and
  ensures that the connection is closed.
*/

package main

import "fmt"

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() // function called here with defer
	fmt.Println("Doing some DB operations....")
	fmt.Println("Opps! some crash or network error ....")
	fmt.Println("Returning from function here!")
	return //terminate function execution
}

func main() {
	doDBOperations()
}
