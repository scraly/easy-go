# 05-Switch-Case

Instead of doing a lot of if-else block, if you have several cases, a common way is to use switch-case.
The program reads the conditions from top to bottom. When a condition is OK, switch is stopped.

## Define a switch

```go
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
```

## Define a switch without condition

same as switch true.

```go
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
```

## Break the switch in a for loop

```go
for {
	switch myFunction() {
		case true:
			break;
		case false:
			//do something
	}
}
```

Switch break, not the for loop

## Break in switch AND break the loop

Use "labeled break" to break for/switch (works for for/select too).

```go
loop:
for {
	switch myVar {
		case true:
			//do something
		case false:
			break loop
	}

}
```

## Type switch

A type switch is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value. 

Define a switch and do things depending of the type of the variable:

```go
func doThings(input interface{}) {
	switch v := input.(type) {
	case int:
		fmt.Printf("You entered an integer: %v\n", v)
    case bool:
		fmt.Printf("You entered a bool: %v\n", v)
	case string:
		fmt.Printf("You entered a string: %s (%v bytes long)\n", v, len(v))
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	doThings(42)
	doThings("hello readers")
	doThings(true)
    doThings(4.2)
}
```