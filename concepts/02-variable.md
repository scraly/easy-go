# 02 - Variable

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

```go
fmt.Println(myString)
```