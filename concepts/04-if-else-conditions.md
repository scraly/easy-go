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

## If with several conditions

```go
if name == "myName" && isActiveUser {
    //...
}
```

## If (and multiple) else

```go
if HTTPStatusCode == 200 {
    fmt.Println("Success")
} else if height == 404 {
    fmt.Println("Not found!")
} else {
    fmt.Println("HTTP Status code not implement yet")
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



