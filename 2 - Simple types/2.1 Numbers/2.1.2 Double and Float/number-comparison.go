var a = 0.1
var b = 0.2
var c = 0.3

// Wrong way to compare floating-point numbers
if a + b == c {
	fmt.Println("Equal")
} else {
	fmt.Println("Not Equal")
}
// Output: Not Equal

// Correct way to compare floating-point numbers
if math.Abs(a+b-c) < 0.0001 {
	fmt.Println("Equal")
} else {
	fmt.Println("Not Equal")
}
// Output: Equal
