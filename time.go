package main

import (
	"fmt"
	"time"
)

func timer() {
	timenow := time.Now()
	fmt.Println(timenow)

	fmt.Println(timenow.Format("01-02-2006")) // just the date mm dd yyyy

	fmt.Println(timenow.Format("02-01-2006 ")) // just the date with ddmmyyyy

	fmt.Println(timenow.Format("02-01-2006 15:04:05 ")) // just the date with dd mm yyyy  - hrs min sec

	fmt.Println(timenow.Format("02-01-2006 15:04:05 Monday")) // just the date with dd mm yyyy  - hrs min sec day

}
