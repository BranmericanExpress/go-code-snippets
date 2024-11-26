var someString = "Substring search"
var substring = "search"

if strings.Contains(someString, substring) {
		fmt.Println("Found!")
} else {
		fmt.Println("Not found!")
}

if strings.HasPrefix(someString, "Sub") {
		fmt.Println("Has prefix!")
}

if strings.HasSuffix(someString, "search") {
		fmt.Println("Has suffix!")
}
