var str = "ABC"
var charA = str[0]
// charA is now 65, which is the ASCII code for 'A'

var charB = rune(str[1])
// charB is now 66, which is the ASCII code for 'B'

var charC = string(str[2])
// charC is now "C"

var charList = ""
for i := 0; i < len(str); i++ {
		charList += string(str[i]) + " "
}
// charList is now "A B C "
