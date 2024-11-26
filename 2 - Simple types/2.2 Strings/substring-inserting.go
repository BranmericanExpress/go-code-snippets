var someString = "string"

somestring = "Sub" + someString
// someString is "Substring"

someString = someString + "!"
// someString is "Substring!"

someString = someString[:9] + " inserting" + someString[9:]
// someString is "Substring inserting!"
