package erapse

import (
	//	"fmt"
	"log"
	"runtime"
	"time"
)

// settings if elapse time show or not
var IsShow = true

func SetGlobalIsShow(isShow bool) {
	IsShow = isShow
}

func GetGlobalIsShow() (isShow bool) {
	isShow = IsShow
	return
}

func ShowErapsedTIme(start time.Time) {
	if IsShow {
		pc, _ /*file*/, _ /*line*/, _ := runtime.Caller(1) // 1: caller function
		f := runtime.FuncForPC(pc)                         // refrection for caller function

		log.Printf("eraps %s: %d μs\n", f.Name(), time.Now().Sub(start).Microseconds())
	}
}
