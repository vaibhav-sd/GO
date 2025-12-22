package main

import "fmt"

func main() {
	fmt.Println("Structs in GO")
	vaibhav := User{"Vaibhav", "vaibhavdagwal@gmail.com", true, 24}
	fmt.Println(vaibhav)
	fmt.Printf("Vaibhav details are: %+v\n", vaibhav)

	vaibhav.GetStatus()
	vaibhav.NewMail()
	fmt.Printf("Vaibhav details are: %+v\n", vaibhav)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Status of User :", u.Status)
}

func (u *User) NewMail() {
	u.Email = "vaibhav@dagwal.com"
	fmt.Println("New mail : ", u.Email)
}
