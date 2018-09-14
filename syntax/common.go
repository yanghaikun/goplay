package syntax

import (
	"fmt"
	"time"
)

func Duration() {
	var n int64
	n = 10
	d := time.Duration(n)
	fmt.Print("duration is %s, period is %s", d, time.Second)


}
