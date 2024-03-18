package utils

import (
	"fmt"
	"time"
)

func Order() string {
	i := time.Now().UnixNano() / int64(time.Millisecond)
	return fmt.Sprintf("SF%v", i)
}
