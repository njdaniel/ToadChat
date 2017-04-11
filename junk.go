package main

import (
	r "github.com/dancannon/gorethink"
	"fmt"
)

type User struct {
	Id string `gorethink: "id,omitempyt"`
	Name string `gorethink: "name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
		Database: "rtsupport",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	//response, err := r.Table("user").
	//	Insert(user).
	//	RunWrite(session)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//user := User{
	//	Name: "Nick",
	//}

	response, _ := r.Table("user").
		Get("ae07ebda-db72-4697-ac5d-d1dc57b73c1e").
		Delete().
		RunWrite(session)

	fmt.Printf("%#v\n", response)
}
