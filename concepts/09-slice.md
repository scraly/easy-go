# 09-Slice

A slice is a dynamically-sized, flexible view of the elements of an array.
A slice is formed by specifying two indices, a low and high bound, separated by a colon: myArray[low : high]

A slice has both a length and a capacity. 
The length of a slice is the number of elements it contains. // len(mySlice) 
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. // cap(mySlice)  

The zero value of a slice is nil. 
A nil slice has a length and capacity of 0 and has no underlying array. 

Slices can includes other slices.

## Declare a slice of strings

```go
talks := []string{}
```

## Declare a slice of int and initialize the array

```go
numbers := []int{1, 3, 5, 7, 9}
```

## Declare an array, then create a slice from the existing array that contains second, third and fourth number

```go
primes := [6]int{1, 3, 5, 7, 9, 11}
var myNumbers []int = primes[1:4]
fmt.Println(myNumbers)
```

## Add one element (at the end) in an existing slice / Append to a slice

```go
talks = append(talks, category + title)
```

## Add three elements in an existing slice / Append to a slice several elements

```go
talks = talks(slice, title1, title2, title3)
```

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

## Range, iterate over elements of a Slice, display only the values

```go
vendors := []string{"nintendo", "sega", "sony"}
for _, vendor := range vendors {
    fmt.Printf("%d\n", vendor)
}
```

## Create dynamically-sized slice with a length equals to 5

```go
mySlice := make([]int, 5)
//[0 0 0 0 0]
//len:5, cap:5
```

## Create dynamically-sized slice with a length equals to 0 and a capacity equals to 5

```go
mySlice := make([]int, 0, 5)
//[]
//len: 0, cap:5
```

## Create a slice of slice

```go
rows := [][]int{}
```
