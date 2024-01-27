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

## Declare a function with variadic arguments

"..." syntax -> spread operator

```go
func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for _, str := range strs {
        final += str
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}
```

## Anonymous function

Anonymous functions are useful when you want to pass a function as an argument to another function or when you want to return a function from another function. They are also useful when you want to define a function inline without having to name it. In the following sections, we will explore how to use anonymous functions in Golang.

Anonymous functions are useful when you want to define a function inline without having to name it.

```go
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Println("Listening on localhost:8080")

  log.Fatal(http.ListenAndServe(":8080", nil))
```

## Closure

Anonymous functions form the basis of closures in Golang.

This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
	

  We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
	

A Go closure is an anonymous nested function which retains bindings to variables defined outside the body of the closure.

Closures can hold a unique state of their own. The state then becomes isolated as we create new instances of the function.

```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    nextInt := intSeq()
	

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())	

    newInts := intSeq()
    fmt.Println(newInts())
}

//
1
2
3
1
```

The middleware are functions that execute during the lifecycle of a request to a server. The middleware is commonly used for logging, error handling, or compression of data.

In Go, middleware is often created with the help of closures.

```go
func main() {
  http.HandleFunc("/", timer(hello))
  log.Println("Listening on localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func timer(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    f(w, r)
    end := time.Now()
    fmt.Println("The request took", end.Sub(start))
  }
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
```

## Defer

 A defer statement defers the execution of a function until the surrounding function returns. 
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. 

TODO xx

```go
```