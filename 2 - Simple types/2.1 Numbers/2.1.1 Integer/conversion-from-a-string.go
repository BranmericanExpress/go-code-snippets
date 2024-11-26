var strNumber = "10"
number1, err := strconv.Atoi(strNumber) // Convert string to int
number2, err := strconv.ParseInt(strNumber, 10, 0) // Convert string to int64
