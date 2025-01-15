package main

import "fmt"


// This will result in a compile-time error
//a, b, c := 1, 2.1, "hello"  // Invalid at package level
// Explicit declaration with different types
// var x int = 5
// var y float64 = 3.14
// var z string = "world"
// fmt.Println(x, y, z)

func main(){
	var a = "this is a string"
	fmt.Println(a)

	var b,c int= 5,6	//explicit declaration
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int			//without init value, it takes the zero value in the defined type
	fmt.Println(e)		//zero value for int is 0

	var f bool			//zero value for bool is false
	fmt.Println(f)

	var g string		//zero value for string is empty string
	fmt.Println(g)

	h:="hallelujah"		//here type has been inferred by go automatically. This shorthand only works inside functions
	fmt.Println(h)		//and not outside. so not for global vars. Its because it is a shorthand for local variable declarations
						//package level declarations(outside functions), go expects you to use var. 

    i,j,k := 1.1, "silmaril", true  //shorthand declaration with multiple data types
	fmt.Println(i,j,k)



}


/*
explicitly declare the type in the following cases:

    When you need a specific type (e.g., int32, float64) and want to override go's default inference.
    When declaring multiple variables of different types for clarity and to avoid ambiguity.
    When dealing with complex or custom types (e.g., structs, interfaces).
    When you declare a variable without initialization, especially at the package level.
    In function signatures and method declarations, where explicit types are required.

generally, go's type inference works well for most cases, but being explicit with types improves readability, 
especially in complex codebases.
*/


/*
scope resolution:
if you have a local and global var of the same name, only the most local one will be called if you just specify the name.
this is due to variable shadowing withing the respective scope. 
CASE-1:
a global and a local var of same name. 
local can be directly accessed by the name.
for global, you can either write <package_name>.<global_var_name> or you can write a separate func outside func main().
*******

package main

import "fmt"

// Global variable
var x = 10

func main() {
    // Local variable with the same name
    var x = 20
    
    // Local variable `x` shadows the global `x`
    fmt.Println("Local x:", x)  // Prints 20

    // Accessing the global variable using its package name
    fmt.Println("Global x:", main.x)  // Prints 10
	
    // Accessing the global `x`
    fmt.Println("Global x:", globalX())  // Prints 10
}

// Function to access the global `x`
func globalX() int {
    return x // Refers to the global `x`
}

one more thing to remember is that these vars can be of different types. you must use them carefully to avoid causing 
type mismatch.
*******
CASE-2:
There are many variables inside a function and a global var of same names. 
In this case also variable shadowing occurs with variables shadowed in their respective scope.


*/