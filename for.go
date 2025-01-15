package main

import "fmt"

func main(){

	i:=1

	//most basic with single condition
	for i<=3 {
		fmt.Println(i)
		i=i+1
	}

	//classic initial-condition-after for loop
	for j:=0; j<3; j++{       //here short-hand is necessary to init j. if we do not do it, then we have to init j outside
		fmt.Println(j)		  //like var j int and then use it like j=0 inside loop.
	}

	//this for loop is for doing anything n times by ranging over integer. this won't work because The range keyword is used 
	//to iterate over collections like arrays, slices, strings, and maps, but it doesn't work directly with an integer 
	//constant.
	//for i := range 3 {
    //    fmt.Println("range", i)
    //}

	//if you are still adamant about using range then you have to create an empty array and iterate over its indices.
	for i:= range(make([]int, 3)){ //range fn returs 2 vals- index and value stored there. but since here we  only gave 1 var 
		fmt.Println("range" , i)   //to store val, we only get index.
	}
	
	//for without condition
	for {							 //this loop will run infinitely until you break out of it explicitly.
		fmt.Println("loop infinito") 
		break
	}
	//continue to next statement
	for n:=range(make([]int,6)){
		if n%2==0{
			continue
		}
		fmt.Println(n)
	}
	
}

/*
for loop is the only loop in golang.
*/