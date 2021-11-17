package timeconvert_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jadahbakar/timeconvert"
)

func TestTime(t *testing.T) {
	// value := "2021-11-16 14:45:00"
	// value := 1405544146
	value := time.Now()
	res2 := timeconvert.TimeConv(value)
	fmt.Printf("result Testing -> %v\n", res2)

}
