package main

import "fmt"

func main() {

	// greeter()

	// fmt.Println(adder(2, 5))
	// fmt.Println(proadder(2, 5, 22))
	// fmt.Println(proAdderMultipleReturn(2, 5, 22))

	abhi := User{"abhi", "abhi@123", 123, 12}
	abhi.GetAge()
	abhi.SetMail("abhi@223")
	fmt.Println(abhi.Email)//does not get updated so we use pointers

}

// func greeter() {
// 	fmt.Print("namaste")
// }

// // func greeter(how string) {
// // 	fmt.Print(how)
// // }
// // will give error

// func adder(val1, val2 int) int {
// 	return val1 + val2
// }

// func proadder(val ...int) int {
// 	total := 0
// 	for _, value := range val {
// 		total += value
// 	}
// 	return total
// }

// func proAdderMultipleReturn(val ...int) (int, int) {
// 	total, cnt := 0, 0
// 	for _, value := range val {
// 		total += value
// 		cnt++
// 	}
// 	return total, cnt
// }

type User struct {
	Username string
	Email    string
	RolNo    int
	Age      int
}

func (u User) GetAge() {
	fmt.Println(u.Age)
}

func (u User) SetMail(mail string) {
	u.Email = mail
	fmt.Print(u.Email)
}
//here  , the mail does not get manipulated , because it takes a copy of user