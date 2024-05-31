package package2

import (
	"local.packages/erapse"
)

func Get1() bool {
	return erapse.GetGlobalIsShow()
}

func SetGlobalIsShow(isShow bool) {
	erapse.SetGlobalIsShow(isShow)
}
