var movieString = "2010|Inception|Christopher Nolan"
var movieArray = strings.Split(movieString, "|")

year, err := strconv.Atoi(movieArray[0])
var name = movieArray[1]
var director = movieArray[2]

fmt.Println(year, name, director)
// Output: 2010 Inception Christopher Nolan
