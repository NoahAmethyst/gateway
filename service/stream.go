package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type Msg struct {
	Message string
}

func Stream(c *gin.Context) {
	chanStream := make(chan Msg, 10)
	go func() {
		defer close(chanStream)
		for i := 0; i < 5; i++ {
			chanStream <- Msg{Message: fmt.Sprintf("test%d", i)}
			time.Sleep(time.Second * 1)
		}
	}()
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}
