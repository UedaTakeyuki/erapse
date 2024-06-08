package erapse

import (
	"log"
	"reflect"
	"runtime"
	"testing"
	"time"

	cp "github.com/UedaTakeyuki/compare"
)

// Check the difference in resolution due to different methods of measuring elapsed time

func Test_04(t *testing.T) {
	err := test_a()
	cp.Compare(t, err, nil)

	test_b()
	test_c()
	test_d()
	test_e()
	test_f()
	test_fμ()
}

func test_b() {
	start := time.Now()
	defer log.Printf("test_b %d ns\n", time.Now().Sub(start).Nanoseconds())
}

func test_c() {
	start := time.Now()
	_ /*pc */, file /*file*/, _ /*line*/, _ := runtime.Caller(0)
	defer log.Printf(" eraps %s: %d ns\n", file, time.Now().Sub(start).Nanoseconds())
}

func test_d() {
	start := time.Now()
	pc /*pc */, _ /*file*/, _ /*line*/, _ := runtime.Caller(0)
	name := runtime.FuncForPC(pc).Name()
	defer log.Printf(" eraps %s: %d ns\n", name, time.Now().Sub(start).Nanoseconds())
}

func test_e() {
	start := time.Now()
	rv := reflect.ValueOf(test_e)
	name := runtime.FuncForPC(rv.Pointer()).Name()
	defer log.Printf(" eraps %s: %d ns\n", name, time.Now().Sub(start).Nanoseconds())
}

func test_f() {
	start := time.Now()
	log.Printf(" eraps %s: %d ns\n", "no defer", time.Now().Sub(start).Nanoseconds())
}

func test_fμ() {
	start := time.Now()
	log.Printf(" eraps %s: %d μs\n", "no defer", time.Now().Sub(start).Microseconds())
}
