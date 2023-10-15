package erapse

import (
//	"fmt"
	"log"
	"runtime"
	"time"
)

func ShowErapsedTIme(start time.Time) {
	pc, _ /*file*/, _ /*line*/, _ := runtime.Caller(1) // 1: caller function
	f := runtime.FuncForPC(pc)                         // refrection for caller function
	//fmt.Printf("eraps %s: %d μs\n", f.Name(), time.Now().Sub(start).Microseconds())
	log.Printf("eraps %s: %d μs\n", f.Name(), time.Now().Sub(start).Microseconds())
}
