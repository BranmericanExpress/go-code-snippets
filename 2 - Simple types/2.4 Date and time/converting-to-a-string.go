var now = time.Now()

var shortStyle = now.Format("1/2/06, 3:04 PM")
var mediumStyle = now.Format("Jan 2, 2006 3:04:05 PM")
var longStyle = now.Format("January 2, 2006 3:04:05 PM MST")
var fullStyle = now.Format("Monday, January 2, 2006 3:04:05 PM MST")

/*
Format 						= "Value"
stdLongMonth      = "January"
stdMonth          = "Jan"
stdNumMonth       = "1"
stdZeroMonth      = "01"
stdLongWeekDay    = "Monday"
stdWeekDay        = "Mon"
stdDay            = "2"
stdUnderDay       = "_2"
stdZeroDay        = "02"
stdHour           = "15"
stdHour12         = "3"
stdZeroHour12     = "03"
stdMinute         = "4"
stdZeroMinute     = "04"
stdSecond         = "5"
stdZeroSecond     = "05"
stdLongYear       = "2006"
stdYear           = "06"
stdPM             = "PM"
stdpm             = "pm"
stdTZ             = "MST"
*/