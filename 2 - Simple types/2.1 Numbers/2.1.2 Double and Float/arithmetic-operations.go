var d1 = 8.5 + 3.4 	// d1 is 11.9
var d2 = 8.5 - 3.4 	// d2 is 5.1
var d3 = 8.5 * 2 		// d3 is 17
var d4 = 8.5 / 2 		// d4 is 4.25

// Modulo operator
var d5 = math.Mod(8.5, 3) 	// d5 is 2.5
var d6 = math.Mod(-8.5, 3) 	// d6 is -2.5

// Float division
var f = 7.5
var d7 = int(f) / 2 // d7 is 3
var d8 = -d7 				// d8 is -3

// Increment and decrement
var d9 = 3.5
d9++ 		// d9 is 4.5
d9-- 		// d9 is 3.5
d9 += 2 // d9 is 5.5
d9 -= 2 // d9 is 3.5

// Multiple assignment
d10, d9 := d9, d9+1 			// d10 is 3.5, d9 is 4.5
d11, d9 := d9+1, d9+1 		// d11 is 5.5, d9 is 5.5
var d12 = math.Abs(-8.5) 	// d12 is 8.5
