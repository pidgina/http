package service

import "time"

func TimeNow() string {
	NowTime := time.Now()
	layout := "02.01.2006 15:04"
	time := NowTime.Format(layout)
	return time
}
