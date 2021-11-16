package timeconvert_test

import (
	"fmt"
	"testing"

	"github.com/jadahbakar/timeconvert"
)

func TestTime(t *testing.T) {
	// var time2 = time.Date(2021, 11, 16, 14, 45, 0, 0, time.UTC)
	// res := timeconvert.TimeConv(time2)
	// fmt.Println(res)

	// var layoutFormat, value string
	// layoutFormat = "2006-01-02 15:04:05"
	value := "2021-11-16 14:45:00"
	res2 := timeconvert.TimeConv(value)
	fmt.Printf("Testing -> %v\n", res2)

}
