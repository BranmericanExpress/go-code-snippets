// install packages:
// go get golang.org/x/text/language
// go get golang.org/x/text/message

package main

import (
	"fmt"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var e = 2.71828

var s1 = strconv.FormatFloat(e, 'f', 6, 64)
// s1 is "2.718280"

var s2 = fmt.Sprintf("%f.3", e)
// s2 is "2.718"

var s3 = fmt.Sprintf("%.2e", e/100)
// s3 is "2.72e-02"

p := message.NewPrinter(language.English)
var s4 = p.Sprintf("%.2f", e*100000)
// s4 is "271828.00"
