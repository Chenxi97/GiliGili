package main

//流量控制
import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Println("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connction coming: %d\n", c)
}

func ConnLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !CL.GetConn() {
			c.String(http.StatusTooManyRequests, "Too Many Requests!")
			return
		}
		c.Next()
		//使用后释放链接
		CL.ReleaseConn()
	}
}
