var a = big.NewFloat(270000000000)
// a is 270000000000

var b = big.NewFloat(2)
var c = big.NewFloat(1000)

a = a.Add(a, b)
// a is 270000000002

a = a.Quo(a, c)
// a is 2.70000000002e+08
