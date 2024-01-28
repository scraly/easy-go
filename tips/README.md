# Useful usage

## Convert

### float to int

```go
myFloat := 16.26
myInt := int64(myFloat)
```

/!\ Casting a float to an integer truncates the floating point portion. So the result will be the part at the left of the ".".


## string subtilities

### length of a string

Note: you have to import utf8 package:
```go
import "unicode/utf8"
```

```go
myString := "hello"
length := utf8.RuneCountInString(myString)
```

### Parse a string

Note: you have to import utf8 package:
```go
import "unicode/utf8"
```

```go
fileName := "FR03"

minNbChars := 4
fileNameNbChars := utf8.RuneCountInString(fileName)
if fileNameNbChars == minNbChars {
    fileNameToAnalyze := []rune(fileName)

    //Parsing
    regionCode := string(fileNameToAnalyze[0:2])
    yearCode := string(fileNameToAnalyze[2:4])
	
	fmt.Printf("RegionCode: %s, yearCode: %s", regionCode, yearCode)
}
```

## Random

TODO: xxx