# 03-Constant

## Types

Constants can be character, string, boolean, or numeric values.
They can not be more complex types like slices, maps and structs.

## Declare a constant

```go
const myConst = "Hello you"
```

/!\ Constants can't use the := short declaration syntax.

## Declare several constants

```go
const (
	myString = "hello"
    myBool = true
)
```

## Declare computed constant

```go
const firstName = "Aurelie"
const lastName = "Vache"
const fullName = firstName + " " + lastName
```
