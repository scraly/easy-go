# 03-Constant

## Types

Constants can be character, string, boolean, or numeric values.
They can not be more complex types like slices, maps and structs.

## Declare a constant

```go
const myConst = "Hello you"
```

/!\ Constants can't use the := short declaration syntax.

## Declare several constants in a block

```go
const (
	myConstString = "hello"
    myBool = true
)
```

## Declare computed constant

```go
const firstName = "Aurelie"
const lastName = "Vache"
const fullName = firstName + " " + lastName
```

## Display a constant

```go
fmt.Println(myConstString)
```

## declare a constant and force explicit type through a conversion

```go
const pi = float32(3.14)
```

## iota

creating incrementing constants in Go.

iota is a predeclared constant which can only be used in other constant declarations

 When the predeclared iota constant is used in a custom constant declaration, at compile time, within the custom constant declaration, its value will be reset to 0 at the first constant specification of each group of constants and will increase 1 constant specification by constant specification.

 the value of iota is n (starting from zero). So iota is only useful in group-style constant declarations. 

```go
package main

import (
    "fmt"
)

const (
    // rainbow colors
    Red int = iota
    Orange
    Yellow
    Green
    Blue
    Indigo
    Violet
    )

func main() {
    fmt.Println("Red:", Red, "Orange:", Orange, "Yellow:", Yellow, "Green:", Green, "Blue:", Blue, "Indigo:", Indigo, "Violet:", Violet)
}
```