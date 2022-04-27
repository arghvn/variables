package main

import "fmt"

//In Golang, just as a package can not be called and not used.
// Here, too,
// if we create a variable, we must use it in our program.
// In this code Variables must be explicitly constructed and quantified as needed.

func main() {

	var a = "initial"

	fmt.Println(a)

	var b, c int = 1, 2

	fmt.Println(b, c)

	var d = true

	fmt.Println(d)

	// Variables declared without a corresponding
	//initialization are _zero-valued_.
	//For example, the zero value for an int is 0

	var e int

	fmt.Println(e)

	//The := syntax is shorthand for declaring and
	//initializing a variable  e.g. for
	//var f string = "short" in this case.

	f := "short"

	fmt.Println(f)

}

// output :
//initial
//1 2
//true
//0
//short

//more details :

//Variable a is constructed with the variable keyword without specifying the type of variable.
//Do this automatically if you do not specify the variable type Becomes.
//In the definition of variable f, we see that a variable is constructed and valued, without specifying the type of variable.
//In fact, the short definition of a variable is that there is no need to write the var keyword and it is enough to use =: instead of =.
//To define this variable in a simple way and by specifying the type of variable should From the code
// var f string = "short"   used .
//In case of not quantifying the variable, a constant value is considered according to the type of variable. For example,
// the value of the variable e would be 0, so if You specify a variable without a value, you must also specify its type.
//Variables b and c also refer to the syntax of defining multiple variables with a single use of the var keyword, and we can create multiple variables in a line.
//
//  var (
//	b int = 1
//	c int = 2
//  )
