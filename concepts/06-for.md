# 06-For

## General

```go
for [INITIAL]; [CONDITION]; [AFTER] {
  //...
}
```

INITIAL: run once at the beginning of the loop and can create
variables within the scope of the loop. (can be ommitted)

CONDITION is checked before each iteration. If the condition doesn't pass
then the loop breaks. (can be ommitted)

AFTER: run after each iteration. (can be ommitted)

INITIAL, CONDITION and AFTER are not mandatory, they are optional.

## For loop statement

```go
for i := 0; i < 15; i++ {
  fmt.Println(i)
}
```

## For loop with only a condition

```go
for length < 10 {
  //...
}

food := 1
for food < 5 {
  fmt.Println("cat can eat more food, food currently ate:", food)
  food++
}
fmt.Println("cat has ate", food, "food")
```

## Infinite loop / while

While loop don't exists in Go, you must add a for loop without any condition to do an infinite loop.

```go
for {
//...
}
```

## For loop with continue

continue: stops the current iteration of a for loop and continues to the next iteration. 

Display even numbers:

```go
for i := 0; i < 10; i++ {
  if i % 2 == 0 {
    fmt.Println(i)
    continue
  } 
}
```

## For loop with break

break: stops the current iteration of a loop and exits the loop.

```go
for i := 0; i < 10; i++ {
  if i == 5 {
    break
  }
  fmt.Println(i)
}
```
