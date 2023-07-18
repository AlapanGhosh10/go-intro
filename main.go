package main

import (
	"fmt"
	"intro/helper"
	"sync"
	"time"
)

type UserData struct {
	firstName string
	lastName string
	emailId string
	rollNo int
}

var wg = sync.WaitGroup{}

func main() {
	var entries = make([]UserData, 0)
	var fname, lname, email string
	// var i int = 0
	var isValidName, isValidEmail bool
	
	fname, lname, email = getUserInput()
	isValidName, isValidEmail = helper.IsValid(fname, lname, email)
	if isValidEmail && isValidName {
		var entry = UserData {
			firstName: fname,
			lastName: lname,
			emailId: email,
			rollNo: helper.GenerateRoll(),
		}
		entries = append(entries, entry)
		wg.Add(1)
		go greetUsers(entries[len(entries) - 1])
		// i++
	} else {
		if !isValidName {
			fmt.Println("Enter enter name correctly.")
		}
		if !isValidEmail {
			fmt.Println("Enter email correctly.")
		}
	}
	wg.Wait()
}

func greetUsers(entry UserData) {
	time.Sleep(10 * time.Second)
	var msg =  fmt.Sprintf("%v %v your roll number is %v.\n", entry.firstName, entry.lastName, entry.rollNo)
	fmt.Println("##########")
	fmt.Printf("Sending roll no:\n %v to email address %v\n", msg, entry.emailId)
	fmt.Println("##########")
	wg.Done()
}

func getUserInput() (string, string, string) {
	var fname, lname, email string
	fmt.Print("Enter First Name of student: ")
	fmt.Scan(&fname)
	fmt.Print("Enter Last Name of student: ")
	fmt.Scan(&lname)
	fmt.Print("Enter email of student: ")
	fmt.Scan(&email)
	return fname, lname, email
}