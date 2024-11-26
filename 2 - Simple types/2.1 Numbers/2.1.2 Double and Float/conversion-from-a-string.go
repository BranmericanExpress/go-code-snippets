// The first method
var strPi = "3.14"
piFloat, err := strconv.ParseFloat(strPi, 64)
// piFloat is 3.14 (float64)

// The second method
var e = "2.71828"
eFloat, err := strconv.ParseFloat(e, 64)
// eFloat is 2.71828 (float64)

// The third method
var strHalf = "0,5"
strHalf = string.Replace(strHalf, ",", ".", 1)
halfFloat, err := strconv.ParseFloat(strHalf, 64)
// halfFloat is 0.5 (float64)
