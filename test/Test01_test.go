package erapse

import (
	"testing"
	"time"

	cp "github.com/UedaTakeyuki/compare"
	"local.packages/erapse"
)

func Test_01(t *testing.T) {
	err := test_a()
	cp.Compare(t, err, nil)
}

func test_a() (err error) {
	defer erapse.ShowErapsedTIme(time.Now())
	return
}
