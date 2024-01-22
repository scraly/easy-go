# Maps

The zero value of a map is `nil`.

## Initialize a map with key of type string and value of type boolean:

```go
	submitted := map[string](bool) {
    	"Aurelie": true,
    	"Titi": false,
	}
```

## Display the number of key/value elements:

```go
fmt.Println(len(submitted)) // 2
```

## Check if a user is in the map:

```go
	user := "Aurelie"
	if submitted[user] {
		fmt.Println(user, "submitted at the conference")
	}
```

### Check if a key exist (with "comma ok" idiom)

elem, ok := myMap[key]

Return a value of an elem and a bool that check if elem exist or not:

```go
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}

// "comma ok" idiom
seconds, ok := timeZone[tz]; ok {
    return seconds
}
```

### Add an elem

myMap[key] = elem

```go
submitted["Toto"] = true
```

### Get an elem

elem = myMap[key]

### Delete an elem in a map

```go
delete(timeZone, "PDT")
```

