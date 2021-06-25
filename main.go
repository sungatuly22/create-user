package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	login     string
	name      string
	surname   string
	birthYear int
}

func (username User) Save(f *os.File) {
	data := username.login + " " + username.name + " " + username.surname + " " + strconv.Itoa(username.birthYear) + "\n"
	w := bufio.NewWriter(f)
	_, err := w.WriteString(data)
	errCheck(err)
	w.Flush()
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var name, surname, login string
	var birthYear int
	username := make(map[string]User)
	loginCheck := make(map[string]bool)

	f, err := os.Create("information.txt")
	errCheck(err)
	defer f.Close()

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
		username[login].Save(f)
		fmt.Printf("The User %s %s has been created\n", surname, name)
	}
}
