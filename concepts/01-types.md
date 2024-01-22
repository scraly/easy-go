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
