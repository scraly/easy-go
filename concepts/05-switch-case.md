# 05-Switch-Case

Instead of doing a lot of if-else block, if you have several cases, a common way is to use switch-case.
The program reads the conditions from top to bottom. When a condition is OK, switch is stopped.

## Define a switch

```go
	switch HTTPStatusCode {
	case 200:
		fmt.Println("Success")
	case 202:
		fmt.Println("Status Accepted")
	case 404:
		fmt.Println("Not found!")
	case 500:
		fmt.Println("Internal Server Error")
	default:
		fmt.Println("HTTP Status Code not found")
	}
```

## Define a switch without condition

same as switch true.

```go
os := runtime.GOOS
switch {
case os == "darwin":
	fmt.Println("goarch = amd64 or arm64")
case os == "linux":
	fmt.Println("goarch = amd64, arm64, ppc64le or s390x")
case os == "freebsd", os == "netbsd", os == "openbsd", os == "windows":
	fmt.Println("goarch = amd64")
default:
	fmt.Println("unknown")
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
func detect(input interface{}) {
	switch v := input.(type) {
	case int:
		fmt.Printf("Integer detected: %v\n", v)
	case bool:
		fmt.Printf("Bool detected: %v\n", v)
	case string:
		fmt.Printf("String detected: %s (%v bytes)\n", v, len(v))
	default:
		fmt.Printf("Unknown type detected: %T\n", v)
	}
}

func main() {
	detect(42)
	detect("hello")
	detect(true)
	detect(3.14)
}
```