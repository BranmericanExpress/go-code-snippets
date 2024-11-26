// decimal number system
var decimal = int64(42)

// octal number system
var octal = 042
// octal is 34

// hexadecimal number system
var hexadecimal = 0x2a
// hexadecimal is 42

// 42 to decimal string
var strDecimal = strconv.FormatInt(decimal, 10)
// strDecimal is "42"

// 42 to octal string
var strOctal = strconv.FormatInt(decimal, 8)
// strOctal is "52"

// 42 to hexadecimal string
var strHexadecimal = strconv.FormatInt(decimal, 16)
// strHexadecimal is "2a"

// 42 to binary string
var strBinary = strconv.FormatInt(decimal, 2)
// strBinary is "101010"
