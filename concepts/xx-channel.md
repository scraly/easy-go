# xx-Channel

Note: data flows in the direction of the arrow `<-`

## Create a channel

```go
myChannel := make(chan int)
```

## Send myVar to myChannel

```go
myChannel <- myVar
```

## Receive from a channel and assign a value to myVar

```go
myVar := myChannel
```