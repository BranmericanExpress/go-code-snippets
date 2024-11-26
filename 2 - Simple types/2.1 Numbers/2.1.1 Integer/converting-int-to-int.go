var d1 = math.MaxInt64
// d1 is 9223372036854775807

var d2 = int32(d1)
// d2 is -1

var d3 = int(d1)
// d3 is 9223372036854775807

d1 = 10
var d4 = int(d1)
// d4 is 10

var d5 = math.MaxInt32
// d5 is 2147483647

var d6 = int64(d5)
// d6 is 2147483647
