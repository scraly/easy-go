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

## Concat two or more slices

Go 1.22 new!!

Starting from Go 1.22, concatenating slices is straightforward with the introduction of the slice.Concat() method. This method allows for the combination of multiple slices into a single new slice. The method signature is as follows:

```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []string{"James", "Wagner", "Christene", "Mike"}
	s2 := []string{"Paul", "Haaland", "Patrick"}
	s3 := []string{"Peter", "Mark", "Luke"}

	s4 := slices.Concat(s1, s2, s3)
	fmt.Println(s4)
}
```

Result:
[James Wagner Christene Mike Paul Haaland Patrick Peter Mark Luke]

or

```go

	num1 := []int{0, 2, 4}
	num2 := []int{6, 8}
	num3 := []int{10, 12, 14, 16}

	numbers := slices.Concat(num1, num2, num3)

	fmt.Println("numbers =", numbers
```

Result:
numbers = [0 2 4 6 8 10 12 14 16]