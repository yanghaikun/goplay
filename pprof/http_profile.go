package pprof

import (
	"log"
	 _ "net/http/pprof"
)

func ServeHttp() {
	log.Printf("run hello world server")
	ServeHttp()
}
