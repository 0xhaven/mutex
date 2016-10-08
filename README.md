# mutex
Simple mutex implementation interview question

## Task
Implement a simple [spinlock](https://en.wikipedia.org/wiki/Spinlock) based on Go [sync/atomic](https://golang.org/pkg/sync/atomic) primatives.

```Go
var m Mutex
m.Lock()
m.Unlock()
```
### Extra credit
* [Read-write Mutex](https://golang.org/pkg/sync/#RWMutex)
* [WaitGroup](https://golang.org/pkg/sync/#WaitGroup)
* Measure performance against sync.Mutex and explain differences.
