

## Range, iterate over elements of a Slice

Usage:

```go
for <INDEX>, <ELEMENT> := range <SLICE> {
}
```

```go
vendors := []string{"nintendo", "sega", "sony"}
for i, vendor := range vendors {
    fmt.Println(i, vendor)
}
```

## Range, iterate over elements of a Slice, display only the index

```go
for i := range attendees {
    fmt.Println(i)
}
```

## range over integers

go 1.22 new!

