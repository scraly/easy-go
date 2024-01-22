# 04-If-conditions

Variables declared inside a if block don't exists outside of the block.

## Comparision operators

    == equal to
    != not equal to
    < less than
    > greater than
    <= less than or equal to
    >= greater than or equal to

## If statement

```go
if age > 6 {
    fmt.Println("older")
} 
```

## If  (and multiple) else

```go
if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}
```

## If with initial statement

normal/common:

```go
length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}
```

Same code but with initial statement:

```go
if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}
```



