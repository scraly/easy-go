# 10-Struct


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



TODO