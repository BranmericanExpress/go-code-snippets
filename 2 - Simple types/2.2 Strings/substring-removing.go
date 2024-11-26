var someString = "Substring removing!"

// There is no built-in function to remove a substring from a string in Go.
someString = someString[0:9] + someString[18:]
// someString is "Substring!"

someString = someString[3:]
// someString is "string!"
