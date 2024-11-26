var startString = "3, 2, 1, Go!"
var newString = strings.Replace(startString, "1", "One", 1)
// newString is "3, 2, One, Go!"

// Replace all instances of a substring
var oneString = "1 1 1"
newOneString = strings.Replace(oneString, "1", "One", -1)
// newOneString is "One One One"
