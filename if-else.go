package main

import "fmt"

func main(){
	if n:=9; n<0{
		fmt.Println("n<0")
	}else if n>10{
		fmt.Println("n>10")
	}else if n%3==0{
		fmt.Println("n is divisible by 3")
	}
}