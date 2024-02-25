# 08-Array

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

## Loop over allelements of an array

```go
for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
}
```

Result:
numbers[0] = 1
numbers[1] = 3
numbers[2] = 5
numbers[3] = 9
numbers[4] = 11
numbers[5] = 13


## Display the length/list of elments of an array

```go
myArray := [2]string{"Mulder", "Scully"}
length := len(myArray)
fmt.Println(length)
```

Result:
2

## multi-dimensional array

initializes a 2D array with two rows and three columns

```go
//initializes a 2D array with two rows and three columns
numbers2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
fmt.Println(numbers2D)
```

```go
// get the value at row 1 (second row) and column 2 (third column)
value := numbers2D[1][2]
fmt.Println(value)
```

get the element

```go
numbers2D[0][1] = 7
fmt.Println(numbers2D)
```


display all the element of the multi dimesional array

```go
for i := 0; i < len(numbers2D); i++ {
   for j := 0; j < len(numbers2D[i]); j++ {
     fmt.Printf("numbers2D[%d][%d] = %d\n", i, j, numbers2D[i][j])
   }
	}
```

Result:

numbers2D[0][0] = 1
numbers2D[0][1] = 7
numbers2D[0][2] = 3
numbers2D[1][0] = 4
numbers2D[1][1] = 5
numbers2D[1][2] = 6