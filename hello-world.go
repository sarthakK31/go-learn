package main

import "fmt"

func main(){
	fmt.Println("Hello World at long last!")
}

/* 
package main marks the entry point of the go program. its only in the entrypoint file and nowhere 
else like handlers utilities etc. main is a speacial package that contains the main function. the
go runtime looks for main.main function when trying to execute the program. without main function 
you cannot run the code. whatever file has package main at the top of its code, is usually the 
entrypoint of the code. 
'package main' this line actually means that on this page we are defining the package called main.
its not some outside package being imported. 
*/
/*
import fmt is that we are importing a package called fmt. its a really useful pacakage that provides 
all the fucntions for formatted I/O like fmt.Println and fmt.Printf . this statement exists in most 
go files because of the modularity of Go. In golang, you have to explicitly mention all the packages
you are going to use in that particular file. scoped usage is there. this is neater and better defined. 
Go does not automatically import even common packages. This probably increases the speed of execution.
*/
/*
fmt.Println : here we are using the fmt (stands for format) package. it has the function Println. It prints
text or other stuff in the console.
*/

/*
for running the go files, you write '''go run <filename>'''. If you write '''go build''' then it constructs 
a binary that you can use separately too. in go run, the binary is temporary and discarded after use. it is 
retained only by go build command. This binary contains everything and does not need the go runtime to run. so
it can be put on a different system and run without go installed.
*/