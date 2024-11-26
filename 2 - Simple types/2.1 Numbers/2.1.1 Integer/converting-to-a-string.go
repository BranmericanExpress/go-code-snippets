var number = 42
var s1 = strconv.Itoa(number) // s1 is a string with the value "42"
var s2 = strconv.FormatInt(int64(number), 10) // s2 is a string with the value "42"
var s3 = fmt.Sprintf("%d", number) // s3 is a string with the value "42"
var s4 = fmt.Sprintf("%03d", number) // s4 is a string with the value "042"