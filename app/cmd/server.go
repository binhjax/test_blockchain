package cmd
import (
	  // "fmt"
		"log"
		// "time"
	  "github.com/urfave/cli/v2"
)

func NodeServer(ctx *cli.Context) error{
		// router := router.NewRouter()
		// server := &fasthttp.Server{
		// 	Name:    "JWT API Server",
		// 	Handler: router.HandleRequest,
		// }
		//
		// params := config.Configuration.Parameters.Cmd
		// NUM_RESPONSE_HANDLER  = params.NumResponseHandler
		//
		// nbResponseHandler := NUM_RESPONSE_HANDLER
		// var wg sync.WaitGroup
		//
		// queueName := config.GetQueue("outqueue")
		// workers.SubPool.HandleQueue(&wg, "API",queueName,nbResponseHandler,tasks.HandlerOutQueueMessage)
		//
		// wg.Add(1)
		// go func() {
		// 			log.Println("cmd.ApiServer","Start listening")
		// 			if config.Configuration.Webserver.Tls {
		// 				log.Println("cmd.ApiServer","Start server using TLS ")
		// 				panic(server.ListenAndServeTLS(":"+ config.Configuration.Webserver.Port, config.Configuration.Webserver.CertificateFile,config.Configuration.Webserver.KeyFile))
		// 			} else {
		// 				log.Println("cmd.ApiServer","Start server without TLS : " + config.Configuration.Webserver.Port)
		// 				panic(server.ListenAndServe(":"+ config.Configuration.Webserver.Port))
		// 			}
		// 			defer wg.Done()
		// }()
		//
		// connection.MaintainConnection(&wg)
		// workers.MaintainPubConnection(&wg)
		// workers.MaintainSubConnection(&wg)
		//
		// wg.Wait()
		log.Println("cmd.ApiServer","Finished Server ")
		return nil
}

func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		// for _, name := range ctx.FlagNames() {
		// 	if ctx.IsSet(name) {
		// 		ctx.GlobalSet(name, ctx.String(name))
		// 	}
		// }
		return action(ctx)
	}
}
