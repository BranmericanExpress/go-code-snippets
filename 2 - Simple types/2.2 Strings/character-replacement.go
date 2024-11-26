var str = "1-3-2"
var arrStr = strings.Split(str, "-")

fmt.Println(arrStr) 
// arrStr is [1 3 2]

arrStr[2] = "2"
arrStr[4] = "3"

fmt.Println(arrStr)
// arrStr is [1 2 3]

// Replace multiple characters in a string
var anotherStr = "1-/[=2=/]-3"
var replacer = strings.NewReplacer("[", "", "]", "", "=", "", "/", "")
var newStr = replacer.Replace(anotherStr)
// newStr is "1-2-3"
