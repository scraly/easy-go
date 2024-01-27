# 11-Interface

Interfaces are collections of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.

(In Go, an interface is a custom type that other types are able to implement)
( Interfaces are named collections of method signatures, and when other types implement all the required methods, they implicitly implement the interface.)

Go does not have classes. However, you can define methods on types. 

A type implements an interface by implementing its methods. 

/!\ Unlike in Java, there is no "implements" keyword.


In Go, an interface is a custom type that other types are able to implement

## Define an interface and its struct that implements it

```go
type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```

A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.


Tips: name the arguments and return data.

Bad practice:
```go
type Copier interface {
  Copy(string, string) int
}
```

Good practice:
```go
type Copier interface {
  Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
```


    Interfaces are not classes, they are slimmer.
    Interfaces don’t have constructors or deconstructors that require that data is created or destroyed.
    Interfaces aren’t hierarchical by nature, though there is syntactic sugar to create interfaces that happen to be supersets of other interfaces.
    Interfaces define function signatures, but not underlying behavior. Making an interface often won’t DRY up your code in regards to struct methods. For example, if five types satisfy the fmt.Stringer interface, they all need their own version of the String() function.
