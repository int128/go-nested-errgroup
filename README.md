# go-nested-errgroup

## RQ

Questions:

1. Does `sync.WaitGroup` wait for all goroutines if `wg.Add(1)` is called in a running goroutine?
1. Does `errgroup.Group` wait for all goroutines if `eg.Go()` is called in a running goroutine?

Answers:

1. Yes
1. Yes

## Verify

Run [main.go](main.go).

```sh
go run main.go
```

```
2019/10/16 10:53:13 starting wg#main
2019/10/16 10:53:13 starting wg#1
2019/10/16 10:53:14 starting wg#2
2019/10/16 10:53:15 finished wg#1
2019/10/16 10:53:17 finished wg#2
2019/10/16 10:53:17 finished wg#main
2019/10/16 10:53:17 starting eg#main
2019/10/16 10:53:17 starting eg#1
2019/10/16 10:53:18 starting eg#2
2019/10/16 10:53:19 finished eg#1
2019/10/16 10:53:21 finished eg#2
2019/10/16 10:53:21 finished eg#main
```
