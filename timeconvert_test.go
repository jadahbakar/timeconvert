package timeconvert_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jadahbakar/timeconvert"
)

func TestTimeString(t *testing.T) {
	value := "2021-11-16 14:45:00"
	res := timeconvert.TimeConv(value)
	fmt.Printf("result Testing -> %v\n", res)

}

func TestTimeInt(t *testing.T) {
	value := 1405544146
	res := timeconvert.TimeConv(value)
	fmt.Printf("result Testing -> %v\n", res)

}

func TestTimeNow(t *testing.T) {
	value := time.Now()
	res := timeconvert.TimeConv(value)
	fmt.Printf("result Testing -> %v\n", res)
}
