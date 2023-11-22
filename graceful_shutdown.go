package main

import (
	"context"
	"gateway/utils/log"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// gracefulShutdown waits for termination syscalls and doing clean up operations after received it
func gracefulShutdown(ctx context.Context, r *gin.Engine) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGSTOP, syscall.SIGKILL, syscall.SIGHUP)
	go func() {
		sig := <-signalChannel
		defer close(signalChannel)
		log.Info().Msgf("receive signal:%+v,graceful shutdown", sig)
		time.Sleep(3 * time.Second)
		log.Info().Msgf("graceful shutdown done")
	}()
}
