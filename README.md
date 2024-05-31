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

The result is as follows:

```
eraps main.DoSomething: 1896985 μs
```

## Why is "erapse" spelt R? Why not L of "elapse"?
Due to my mother language being Japanese which doesn't have a difference in pronunciation between **R** and **L**, I carelessly misspelt. ^^; By the time I noticed the mistake, I had already used it a lot, so it was hard to correct the spelling now, so I thought it was okay and left it as is. I'm sure there are many other typos and grammatical errors in the document I wrote, so please feel free to point them out. Thx!

## for other language
**C++** version is available at https://github.com/UedaTakeyuki/elapse.

## history
- 2020.07.06 created from scratch
- 2024.05.31 add test
