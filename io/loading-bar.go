package io

import (
	"github.com/cheggaaa/pb/v3"
	"time"
)

func LoadingBar(ms int) {
	count := ms
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()

}
