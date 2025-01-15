package main
import(
	"fmt"
	"math"
)

const k string = "rivendell"

func main(){
	fmt.Println(k)

	const n=500000000

	const d = 3e20/n   //3e20= 3*10^20   its scientific notation
	//here d is a 'typed-constant'. 3e20 is an 'untyped constant'. when we talk about numeric constant, we mean 3e20 here. 
	//its type is only decided when its being used as a part of an expression or in some context.
	//its a literal value that can be computed hence automatically interpreted as a numeric constant (with no type).

	fmt.Println(int64(d)) // A numeric const has no type unless explicitly given one by explicit conversion.
	fmt.Println(math.Sin(n)) //A number can be given a type by using it in a context like here sin expects float64
}

// constants can be used everywhere var can be used. 
// constants are offered for character, string, bool and numeric values.
// *const* declares the constant val

/*
Constant expressions perform arithmetic with arbitrary precision.
here this means that const expressions are not restricted by the usual limits of size. 
it is because the const expressions are calculated at the time of comilation and not runtime. so they are only limited 
by the physical size of the memory. So if the answer of the expression is very large, it will not be truncated due to the
data-type of the answer. its also not type restricted like if the expression is using int by=ut the answer is float then 
that is okay. for example:

const x = 1 << 100  // Result is a very large number which is 1267650600228229401496703205376

In Go, this is evaluated as a constant with arbitrary precision, meaning the compiler will compute and store the value 
correctly without worrying about integer overflow or truncation, which would be a concern in a typical variable-based 
calculation at runtime.

Using the term "arbitrary" in "arbitrary precision" emphasizes that the precision of constant expressions is not limited 
by a fixed boundary — it’s as precise as needed, within the system’s available resources. This allows for more accurate, 
larger-scale computations at compile-time without the usual limitations that would apply to variables during runtime.

*/