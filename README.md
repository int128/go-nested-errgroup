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
2019/10/16 10:35:19 starting wg#main
2019/10/16 10:35:19 starting wg#1
2019/10/16 10:35:19 starting wg#2
2019/10/16 10:35:20 finished wg#1
2019/10/16 10:35:21 finished wg#2
2019/10/16 10:35:21 finished wg#main
2019/10/16 10:35:21 starting eg#main
2019/10/16 10:35:21 starting eg#1
2019/10/16 10:35:21 starting eg#2
2019/10/16 10:35:22 finished eg#1
2019/10/16 10:35:23 finished eg#2
2019/10/16 10:35:23 finished eg#main
```
