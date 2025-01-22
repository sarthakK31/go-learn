package main

import "fmt"

func main(){
	if n:=9; n<0{					//here you can see that you can declare and init some things in the if else statement.
		fmt.Println("n<0")			//it will be valid in all the branches of the if else ladder.
	}else if n>10{
		fmt.Println("n>10")
	}else if n%3==0{
		fmt.Println("n is divisible by 3")
	}
}

/*
Note that you don’t need parentheses around conditions in Go, but that the braces are required.
There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
*/