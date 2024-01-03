package loader

import "fmt"

type Author struct {
	Email string
	Firstname string
	Lastname string
}

func (a Author) Print() {
	fmt.Printf("Firstname: %s\n", a.Firstname)
	fmt.Printf("Lastname: %s\n", a.Lastname)
	fmt.Printf("Email: %s\n", a.Email)
}
