var a = 5 // 0101
var b = 6 // 0110

// Bitwise AND
fmt.Println(a & b) // 0100 = 4

// Bitwise OR
fmt.Println(a | b) // 0111 = 7

// Bitwise XOR
fmt.Println(a ^ b) // 0011 = 3

// Bitwise AND NOT
fmt.Println(a &^ b) // 0001 = 1

// Bitwise left shift
fmt.Println(a << 1) // 1010 = 10

// Bitwise right shift
fmt.Println(b >> 1) // 0011 = 3
