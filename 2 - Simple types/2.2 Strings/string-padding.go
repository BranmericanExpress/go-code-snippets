var str = "123"

// var len = 10
var padStart1 = fmt.Sprintf("%10s", str)
// padStart1 is "       123"

var padStart2 = fmt.Sprintf("%010s", str)
// padStart2 is "0000000123"

var padEnd1 = fmt.Sprintf("%-10s", str)
// padEnd1 is "123       "

var padEnd2 = fmt.Sprintf("%-010s", str)
// padEnd2 is "1230000000"
