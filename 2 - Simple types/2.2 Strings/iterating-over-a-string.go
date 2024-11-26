import (
	"fmt"
)

var str = "let's iterate over this string"

for _, c := range str {
	fmt.Println(string(c))
}

// Iterating with index
for i, c := range str {
	fmt.Println(
		fmt.Sprintf("Index: %d, Char: %s", i, string(c)),
	)
}
