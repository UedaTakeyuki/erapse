package package2

import (
	"local.packages/erapse"
)

func Get2() bool {
	return erapse.GetGlobalIsShow()
}
