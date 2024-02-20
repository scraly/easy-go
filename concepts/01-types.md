# 01-Types

## Differents types exists in Go:

```
bool

string

int  int8  int16  int32  int64 // numbers
uint uint8 uint16 uint32 uint64 uintptr // positive numbers

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64 // decimal/floating numbers | 32: 32 bits of precision, 64: 64 bits of precision

complex64 complex128 // complex/imaginary numbers
```

The size (8, 16, 32, 64, 128, etc) indicates how many bits in memory will be used to store the variable
The default int and uint types are just aliases that refer to their respective 32 or 64 bit sizes depending on the environment of the user.

## Common types

    bool
    string
    int
    uint
    byte
    rune
    float64
    complex128 (never use in my life...)


##  types of constants which are restricted to certain values and how you accomplish that in Go

custom data type

```go
type SomeType string

const(
    SomeTypeA SomeType = "A"
    SomeTypeB SomeType = "B"
) 
```

Type definitions may be used to define different boolean, numeric, or string types and associate methods with them:

```go
type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT%+dh", tz)
}
```


 If the type definition specifies type parameters, the type name denotes a generic type. Generic types must be instantiated when they are used.

```
type List[T any] struct {
	next  *List[T]
	value T
}
```

In a type definition the given type cannot be a type parameter.

```
type T[P any] P    // illegal: P is a type parameter

func f[T any]() {
	type L T   // illegal: T is a type parameter declared by the enclosing function
}
```

A generic type may also have methods associated with it. In this case, the method receivers must declare the same number of type parameters as present in the generic type definition.

```
// The method Len returns the number of elements in the linked list l.
func (l *List[T]) Len() int  { … }
```
