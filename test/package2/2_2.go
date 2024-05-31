package package2

import (
	"time"

	"local.packages/erapse"
)

func Get2() bool {
	return erapse.GetGlobalIsShow()
}

func Call_2() {
	defer erapse.ShowElapsedTIme(time.Now(), IsShow)
}
