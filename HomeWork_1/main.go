package main

import (
	"fmt"
	"log"
)
		func main() {
			mathDivide()
			fmt.Println("hi!, lets doing the panic ")

		}

		func mathDivide() {
			defer func() {
				if err := recover(); err != nil {
					log.Println("panic detected:", err)
				}
			}()
			fmt.Println(divide(8, 0))
		}

		func divide(a, b int) int {
			return a / b
		}

//func main() {

//defer func() {
//	fmt.Println("defer function is working")
//}()
//	names := []string{

//	"snowboarding",
//	"football",
//	"hockey",
//}

//	fmt.Println("My favorite sport is:", names[len(names)])







