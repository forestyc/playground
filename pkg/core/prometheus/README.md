# prometheus


```
import "github.com/forestyc/playground/pkg/prometheus"

func (j *job) Init() {
    // ...
    j.c = prometheus.NewCounter("name","help", "label1","label2")
    // ...
}

func (j *job) Run() {
    // ...
    j.c.Inc("value1", "value2")
    // ...
}
```



