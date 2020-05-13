package cmd

import (
	  "github.com/urfave/cli/v2"
    // "fmt"
		"log"
		// "time"
    "sync"
    "github.com/valyala/fasthttp"
    "github.com/binhnt-teko/test_blockchain/app/api/router"
)

func ApiServer(ctx *cli.Context) error{
		router := router.NewRouter()
		server := &fasthttp.Server{
			Name:    "JWT API Server",
			Handler: router.HandleRequest,
		}
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
					log.Println("cmd.ApiServer","Start listening")
					if config.Configuration.Webserver.Tls {
						log.Println("cmd.ApiServer","Start server using TLS ")
						panic(server.ListenAndServeTLS(":"+ config.Configuration.Webserver.Port, config.Configuration.Webserver.CertificateFile,config.Configuration.Webserver.KeyFile))
					} else {
						log.Println("cmd.ApiServer","Start server without TLS : " + config.Configuration.Webserver.Port)
						panic(server.ListenAndServe(":"+ config.Configuration.Webserver.Port))
					}
					defer wg.Done()
		}()

		wg.Wait()
		log.Println("cmd.ApiServer","Finished Server ")
		return nil
}
