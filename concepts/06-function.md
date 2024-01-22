# 06-Functions

Functions in Go can take zero or more arguments.

when a variable is passed into a function, that function receives a copy of the variable. The function is unable to mutate the caller's original data.

## Declare a function (with two parameters)

```go
func myFunction(myFirstVar int, mySecondVar string) {
  //...
}
```

## Declare a function (with two parameters with the same type)

```go
func myFunction(hp, hit int) {
  //...
}

```

## Declare a function (with one parameter and return one integer)

```go
func myFunction(myFirstVar int) int {
  //...
  return 6
}
```

## Declare a function (with one parameter and return two strings)

```go
func myFunction(myVar string) (string, string) {
  //...
  x := "hello " + myVar
  y := "bye " + myVar
  return x, y
}
```

## Use a function and ignore the first returned value

A function can return one or several values a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: `_`.

```go
func getPoint() (x int, y int) {
  return 3, 4
}

// ignore y value
x, _ := getPoint()
```

## Declare a function without naming the returned values (naked return / blank return)

```go
func getCoords() (x, y int) {
  // x and y are initialized with zero values

  return // implicitly returns x and y
}
```

Is the same as:

```go
func getCoords() (int, int){
  var x int
  var y int
  return x, y
}
```

## Declare a function and return before the end of the function

```go
func myFunction(length, age, size int) (int, error) {
    //...
	if length == 0 {
		return 0, errors.New("length should not be equals to 0")
	}
    //...
	return length+size, nil
}
```