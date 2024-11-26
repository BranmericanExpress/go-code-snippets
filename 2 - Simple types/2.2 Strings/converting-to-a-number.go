// To int
var s string = "42"

// The first method is to use the strconv package's Atoi function.
i, err := strconv.Atoi(s)

// The second method is to use the strconv package's ParseInt function.
i, err := strconv.ParseInt(s, 10, 0)

// To double and float
var strPi string = "3.14159"

// The first method is to use the strconv package's ParseFloat function.
f, err := strconv.ParseFloat(strPi, 64)

// The second method is to use the strconv package's Atof function.
var strHalf string = "0.5"
f, err := strconv.ParseFloat(strHalf, 64)
