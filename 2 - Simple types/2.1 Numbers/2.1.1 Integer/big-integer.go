var a = big.NewInt(math.MaxInt64)
// a is 9223372036854775807

var b = big.NewInt(255)
var c = big.NewInt(1000)

a = a.Mul(a, c)
// a is 9223372036854775807000

a = a.Add(a, b).Quo(a, b)
// a is 36170086419038336501
