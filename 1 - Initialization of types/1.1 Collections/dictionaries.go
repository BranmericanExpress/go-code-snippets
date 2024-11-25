// Dictionary<String, String>
var languages = map[string]string{
		"en": "English",
		"es": "Spanish",
		"fr": "French",
}

// Dictionary<Int, String>
var numbers = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
}

type employee struct {
	name string
	age int
}

// Dictionary<Int, Employee>
var employees = map[int]employee{
		1: employee{name: "John", age: 25},
		2: employee{name: "Jane", age: 30},
		3: employee{name: "Doe", age: 35},
}
