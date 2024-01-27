# 07-Array

Arrays are fixed sized list of data of the same type.
You can't change the number of elements of an array (append/add an element, delete an element).

## Declare an array of 6 integers

```go
var myInts [6]int
```

## Declare an array of integers and initialize the values

```go
numbers := [6]int{1, 3, 5, 9, 11, 13}
```

## Set value by index

```go
numbers[2] = 7
```

## Display the length/list of elments of an array

```go
length := len(numbers)
```