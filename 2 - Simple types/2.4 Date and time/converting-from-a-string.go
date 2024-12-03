var stringDate = "2019-08-01T15:00:00Z"
var format = "2006-01-02T15:04:05Z"

date, err := time.Parse(format, stringDate)

if err != nil {
		log.Fatal(err)
}

fmt.Println(date)
// Output: 2019-08-01 15:00:00 +0000 UTC
