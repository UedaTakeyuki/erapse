# erapse : show elapsed time of go function

## How to use

Import **time** and **github.com/UedaTakeyuki/erapse** and just call ```defer erapse.ShowErapsedTIme(time.Now())``` at the beginning of the function you want to measure elapsed time, that it! Fx:

```
import (
	"time"
	"github.com/UedaTakeyuki/erapse"
)

func DoSomething(){
	defer erapse.ShowErapsedTIme(time.Now())
```

Result is as follows:

```
eraps main.DoSomething: 1896985 μs
```

## Why is "erapse" spelt R? Why not L of "elapse"?
Due to my mother language is Japanese that doesn't have a difference in pronunciation between **R** and **L**, I simply miss spelt. By the time I noticed the mistake, I had already used it a lot, so it was hard to correct the spelling now, so I thought it was okay and left it as is. I'm sure there are many other typos and grammatical errors in the document I write, so please feel free to point them out. Thx!
