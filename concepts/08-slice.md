# 08-Slice

A slice is a dynamically-sized, flexible view of the elements of an array.
A slice is formed by specifying two indices, a low and high bound, separated by a colon: myArray[low : high]

## Declare a slice of strings

```go
talks := []string{}
```

## Declare an array, then create a slice that contains second, third and fourth number

```go
primes := [6]int{1, 3, 5, 7, 9, 11}
var myNumbers []int = primes[1:4]
fmt.Println(myNumbers)
```

## Add one element in an existing slice

```go
talks = append(talks, category + title)
```

## Delete one element in an existing slice

##TODO: make

