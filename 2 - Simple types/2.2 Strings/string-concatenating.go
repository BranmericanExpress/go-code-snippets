var s1 = "Hello, "
var s2 = "world!"

// The first method is to use the + operator to concatenate two strings.
var s3 = s1 + s2
fmt.Println(s3) 
// Output: Hello, world!

// The second method is to use strings.Join() function to concatenate multiple strings.
var s4 = strings.Join([]string{s1, s2}, "")
fmt.Println(s4)
// Output: Hello, world!

// The third method is to use strings.Builder to concatenate multiple strings.
var sb strings.Builder
sb.WriteString(s1)
sb.WriteString(s2)
var s5 = sb.String()
fmt.Println(s5)
// Output: Hello, world!

// String list concatenation
var numbers = []string{"one", "two", "three"}
var numberList = strings.Join(s6, "; ")
fmt.Println(numberList)
// Output: one; two; three
