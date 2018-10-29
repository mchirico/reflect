package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type User struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Github string `json:"github" default:"a8m"`
}

func main() {
	u := &User{Name: "Ariel Mashraki"}
	// Elem returns the value that the pointer u points to.
	v := reflect.ValueOf(u).Elem()
	f := v.FieldByName("Github")
	// make sure that this field is defined, and can be changed.
	if !f.IsValid() || !f.CanSet() {
		return
	}
	if f.Kind() != reflect.String || f.String() != "" {
		return
	}
	f.SetString("a8m")
	fmt.Printf("Github username was changed to: %q\n", u.Github)

	// Make json and print
	json_data, err := json.MarshalIndent(u, " ", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data))

	// Take json

	var userBlob = []byte(`[
{
  "email": "",
  "name": "Mike Chirico",
  "age": 0,
  "github": "mchirico"
 },
{
  "email": "",
  "name": "Bob Anderson",
  "age": 0,
  "github": "banderson"
 }

]`)


	var users []User
	err = json.Unmarshal(userBlob, &users)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", users)




}
