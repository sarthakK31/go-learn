package main

import(
	"fmt"
	"time"
)

func main(){
	i:=1
	fmt.Println("Write", i, "as")   //this is a basic switch.
	switch i{						  // you can see that the 'default' case is not a necessity unlike cpp	
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	fmt.Println(time.Now().Weekday())

	switch time.Now().Weekday(){			//time.Saturday is required and not simply Saturday can be put here. This is 
	case time.Saturday, time.Sunday:		//because the value returned by time.Now().Weekday() is of time.Weekday() type. 
		fmt.Println("Its the weekend!!!")	//Saturday and Sunday are defined constants in the time package.
	default:
		fmt.Println("Its a weekday.")
	}

	/*
	switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions 
	can be non-constants.
	*/
	switch{
	case t.Hour()>12:
		fmt.Println("Its after noon")
	default:
		fmt.Println("Its before noon")
	}

	/*
	A type switch compares types instead of values. You can use this to discover the type of an interface value. In this 
	example, the variable t will have the type corresponding to its clause.
	*/
}