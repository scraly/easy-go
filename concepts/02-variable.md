# 02 - Variable

We can declare variables at package or at function level, depending of the scope you want.

```go
package main

import "fmt"

var name, city string

func main() {
	var i int
	fmt.Println(name, i, city)
}
```

When you declare a variable of a type, you can't change its type but you can change its value at any time.

## Declare a variable

```go
var myVar string
```

## Declare and initialize a variable 

```go
var myString string = "hello"
```

## Use the short variable declaration to declare a variable and initialize it

The `:=` operator infers the type of the new variable based on the value.
So with this operator, no need to declare the type of the variable.

```go
myString := "hello"
```

## Declare several variables in one line

We can declare one or several variables in a single line.

```go
var name, country string = "Aurelie", "France"
```

Note: If you define the <type> of the variables, it's only possible to define one type of variable par line.

## Declare several variables, of different types, in one line

We can declare one or several variables in a single line.

```go
name, age, isSeaker := "Aurelie", 39, true
```

Note that you don't have to specify the type of the variables.

## Declare several variables in a block

Multiple variable declarations can also be grouped together into a block (for greater readability).

```go
var (
    name string = "Aurelie"
    age int = 39
    country string
   )
```

## Display a variable

Note: you have to import `fmt` package:

```go
import "fmt"
```

Then in a function:

```go
fmt.Println(myString)
```