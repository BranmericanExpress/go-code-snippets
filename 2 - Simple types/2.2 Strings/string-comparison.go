var first = "A"
var second = "B"
var third = "A"

if first == second {
	fmt.Println("first and second are equal")
} else {
	fmt.Println("first and second are not equal")
}
// Output: first and second are not equal

if first == third {
	fmt.Println("first and third are equal")
} else {
	fmt.Println("first and third are not equal")
}
// Output: first and third are equal

if first != second {
	fmt.Println("first and second are not equal")
} else {
	fmt.Println("first and second are equal")
}
// Output: first and second are not equal

if strings.Compare(first, second) == 0 {
	fmt.Println("first and second are equal")
} else {
	fmt.Println("first and second are not equal")
}
// Output: first and second are not equal
