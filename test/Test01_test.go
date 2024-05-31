package erapse

import (
	"testing"
	"time"

	cp "github.com/UedaTakeyuki/compare"
	"local.packages/erapse"
	"local.packages/package1"
	"local.packages/package2"
)

// Basic Test
func Test_01(t *testing.T) {
	erapse.SetGlobalIsShow(false)
	err := test_a()
	cp.Compare(t, err, nil)

	erapse.SetGlobalIsShow(true)
	err = test_a()
	cp.Compare(t, err, nil)

	erapse.SetGlobalIsShow(false)
	err = test_a()
	cp.Compare(t, err, nil)
}

func test_a() (err error) {
	defer erapse.ShowErapsedTIme(time.Now())
	return
}

// Test global variable of the module which is imported multiplly from different package
func Test_02(t *testing.T) {
	package1.SetGlobalIsShow(true)
	cp.Compare(t, package1.Get1(), true)
	cp.Compare(t, package1.Get2(), true)
	cp.Compare(t, package2.Get1(), true)
	cp.Compare(t, package2.Get2(), true)
}

// Test ShowElapsedTIme(start time.Time, isShow bool) {
func Test_03(t *testing.T) {
	package1.SetGlobalIsShow(true)
	package2.SetGlobalIsShow(true)
	package1.SetLocalIsShow(true)
	package2.SetLocalIsShow(false)
	package1.Call_1()
	package1.Call_2()
	package2.Call_1()
	package2.Call_2()

}
