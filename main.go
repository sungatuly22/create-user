package main

import "fmt"

type User struct {
	login     string
	name      string
	surname   string
	birthYear int
}

func main() {

	var name, surname, login string
	var birthYear int
	username := make(map[string]User)
	loginCheck := make(map[string]bool)

	for {
		fmt.Println("Hello, enter the user data with space: name, surname, year of birth")
		fmt.Scanf("%s %s %d", &name, &surname, &birthYear)
		fmt.Println("Create login for User:")
		for {
			fmt.Scanf("%s", &login)
			if loginCheck[login] {
				fmt.Println("This login is registered already. Please, enter the new login:")
				continue
			} else {
				loginCheck[login] = true
				break
			}
		}
		username[login] = User{login, surname, name, birthYear}
		fmt.Printf("The User %s %s has been created\n", surname, name)
	}
}
