package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Type objects

// Name structs definied
type Name struct {
	Family   string
	Personal string
}

// Email struct
type Email struct {
	ID      int
	Kind    string
	Address string
}

// Interest Struct
type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAddress(p *Person, i int) string {
	return p.Email[i].Address
}

func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarline() {
	fmt.Println("********************")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

func SaveJSON(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func main() {
	person := Person{
		ID:        12,
		FirstName: "Yasin",
		LastName:  "Kaplan",
		UserName:  "Yaskobar_",
		Gender:    "Erkek",
		Name: Name{
			Family:   "Kaplan",
			Personal: "YasinBaran",
		},
		Email: []Email{
			Email{ID: 1, Kind: "Work", Address: "kaplantest@mail.com"},
			Email{ID: 2, Kind: "Personel", Address: "testkaplant@mail.com"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "go"},
			Interest{ID: 2, Name: "golang"},
			Interest{ID: 3, Name: "googlelang"},
		},
	}
	WriteMessage("Reading Operation Started")

	WriteMessage("Personal Fullname")
	WriteStarline()

	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarline()

	WriteMessage("\n")

	WriteMessage("PErsonal Email with Index")
	WriteStarline()
	resEmail := GetPersonEmailAddress(&person, 1)
	fmt.Println(resEmail)
	WriteStarline()

	WriteMessage("\n")

	WriteMessage("Personal Email Object with Index")
	WriteStarline()
	resEmail2 := GetPersonEmail(&person, 0)
	fmt.Println(resEmail2)
	WriteStarline()

	WriteMessage("Reading Operation Ended")
	WriteMessage("\n")

	WriteMessage("Writing Operation Started")
	SaveJSON("person.json", person)
	WriteMessage("Writing Operation Ended")

}
