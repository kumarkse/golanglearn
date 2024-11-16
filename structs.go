package main

import "fmt"

func structbase() {
	type User struct {
		Name    string
		Roll_no int
		Email   string
		Age     int
	}

	abhishek := User{"kumar abhishek", 21, "abhi@123", 23}
	fmt.Println(abhishek)
	fmt.Printf("the details of abhishek are : %+v \n", abhishek) // using + before v
	fmt.Println(abhishek.Email)

}
