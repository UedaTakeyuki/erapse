package package2

import (
	"time"

	"local.packages/erapse"
)

var IsShow bool = true

func Get1() bool {
	return erapse.GetGlobalIsShow()
}

func SetGlobalIsShow(isShow bool) {
	erapse.SetGlobalIsShow(isShow)
}

func SetLocalIsShow(isShow bool) {
	IsShow = isShow
}

func Call_1() {
	defer erapse.ShowElapsedTIme(time.Now(), IsShow)
}
