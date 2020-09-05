package main

import "fmt"

func main() {

name()
choice()
}

func name() {
	fmt.Println("Enter first name")
	var firstname string
	fmt.Scanln(&firstname)

	fmt.Println("Enter last name")
	var lastname string
	fmt.Scanln(&lastname)
	
	for i:=1; i<=100; i++{
		if i % 2 == 0 {
			fmt.Print(firstname)
			fmt.Println(":",i)
			fmt.Println()
		}else{
			fmt.Print(lastname)
			fmt.Println(":",i)
		}
	}

}

func choice() {
	fmt.Print("Do you want to try again?(Y/N)")
	var option string
	fmt.Scanln(&option)

	if option == "Y" || option == "y" {
		name()
	}else if option == "N" || option =="n"{
		fmt.Println("exiting..")
	}else {
		fmt.Println("wrong choice, try again")
		choice()
	}
}




   

