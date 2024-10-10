# redis

```

func (j *job) Run() {
    // ...
    j.r, err := redis.NewRedis(config)
    j.r.Get(ctx, "key")
    // ...
}
```



