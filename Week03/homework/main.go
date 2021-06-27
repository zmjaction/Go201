/*
* @Author: ZhaoMingJun
* @Date:   2021/6/6 10:33 下午
 */

package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping"))
	})
	serverOut := make(chan struct{})
		mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})
	server := http.Server{
		Handler: mux,
		Addr: ":8081",
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("exit。。。。")
		case <-serverOut:
			log.Println("out。。。。。")
		}
		timeOutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		log.Println("down")
		return server.Shutdown(timeOutCtx)
	})
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			fmt.Println("terminated  signal%v:",sig)
			return errors.Errorf("terminated  signal%v: ", sig)
		}

	})

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())

}