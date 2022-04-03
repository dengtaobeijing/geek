package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

func main() {

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	group, errCtx := errgroup.WithContext(ctx)

	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done(): // 因为 cancel、timeout、deadline 都可能导致 Done 被 close
				return errCtx.Err()
			case <-chanel: // 因为 kill -9 或其他而终止
				cancel()
			}
		}
	})

	//启动第一个 app
	srv1 := &http.Server{Addr: "0.0.0.0:9011"}

	group.Go(func() error {
		<-errCtx.Done() //阻塞。因为 cancel、timeout、deadline 都可能导致 Done 被 close
		fmt.Println("http serverApp1 stop")
		return srv1.Shutdown(errCtx)
	})
	group.Go(func() error {
		//err := checkGoroutineErr(errCtx)
		return serverApp1(srv1)
	})

	//启动第二个 app
	srv2 := &http.Server{Addr: ":9012"}
	group.Go(func() error {
		<-errCtx.Done() //阻塞。因为 cancel、timeout、deadline 都可能导致 Done 被 close
		fmt.Println("http serverApp2 stop")
		return srv2.Shutdown(errCtx)
	})

	//serverApp2(srv2)
	group.Go(func() error {
		return serverApp2(srv2)
	})

	//10秒后停止 serverApp2
	err := stopServerApp2(srv2, errCtx)
	if err != nil {
		fmt.Println("http stop error: ", err)
	}

	if err := group.Wait(); err != nil {
		fmt.Println("http server error: ", err)
	}
	fmt.Println("all serverApp stop!")

}

//启动第一个http 服务
func serverApp1(srv *http.Server) error {

	http.HandleFunc("/app1", func(resp http.ResponseWriter, req *http.Request) {

		fmt.Fprintln(resp, "serverApp1")
	})

	err := srv.ListenAndServe()

	return err

}

//启动第二个http服务
func serverApp2(srv *http.Server) error {

	//mux := http.NewServeMux()
	http.HandleFunc("/app2", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "serverApp2")
	})

	err := srv.ListenAndServe()

	//err = http.ListenAndServe(":9012", mux)

	return err

}

//10秒后停止  stopServerApp2
func stopServerApp2(srv *http.Server, errCtx context.Context) error {

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	fmt.Println("time out stopServerApp2")
	err := srv.Shutdown(errCtx)

	return err

}
