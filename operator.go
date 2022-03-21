package main

import "fmt"

func main(){
	// const full_name string = "Golang"
	// fmt.Printf("Hello %s",  full_name)

	// var value = (5+2)/3
	// fmt.Println(value)

	var firstCondition bool = 3 < 5
	var secondCondition bool = "golang" == "Golang"
	var thirdConditions bool = 10 != 3.8
	var fourthConditions bool = 11 <= 11

	fmt.Println("first Condition", firstCondition)
	fmt.Println("second Condition", secondCondition)
	fmt.Println("third Conditions", thirdConditions)
	fmt.Println("fourth Conditions", fourthConditions)

	var right =true
	var wrong =false
	var wrongAndRight = wrong && right
	fmt.Printf("left && right \t(%t) \n", wrongAndRight)
}