# 11-Struct

Collection/group of fields.

# Define a struct

```go
type gopher struct {
	Name        string `json:"name"`
	Displayname string `json:"displayname"`
	URL         string `json:"url"`
}
```

# Define a struct which is public (usable in all the application)

```go
type Gopher struct {
	Name        string `json:"name"`
	Displayname string `json:"displayname"`
	URL         string `json:"url"`
}
```

# Define and initiate a struct

```go

type Gopher struct {
	Name        string `json:"name"`
	Displayname string `json:"displayname"`
	URL         string `json:"url"`
}

myGopher := Gopher{"5th-element", "5th Element", "https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png"}

fmt.Println("my Gopher:", myGopher)
fmt.Println("name:", myGopher.Name)
```

# Nested struct

Define:

```go
type gopher struct {
	Name        string 
	Displayname string 
	URL         string 
    Size        Size
}

type Size struct {
  Width int
  Height string
}
```

Set the size of a gopher:

```go
myGopher := gopher{}
myGopher.Size.Width = 300
myGopher.Size.Height = 800
```

## Anonymous struct

It's a strict without a name.

You can use anonymous struct if you will use the struct only one time for example.

```go
myCar := struct {
  Make string
  Model string
} {
  Make: "tesla",
  Model: "model 3",
}

```

## Struct method

Methods can be defined on a struct.

Methods are just functions that have a receiver. A receiver is a special parameter that syntactically goes before the name of the function.

a method is a function associated with a specific data type (struct). A method is called on a specific variable of that type.

```go
type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
  return r.width * r.height
}

var r = rect{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
// prints 50
```

A receiver is just a special kind of function parameter. Receivers are important because they will, as you'll learn in the exercises to come, allow us to define interfaces that our structs (and other types) can implement.