package router

import (
   "github.com/qiangxue/fasthttp-routing"
   "git.teko.vn/loyalty-system/loyalty-onchain-projects/api_server/app/server/controllers"
	 "git.teko.vn/loyalty-system/loyalty-onchain-projects/api_server/app/server/middleware"
)

func NewRouter()  *routing.Router {
	router := routing.New()

  index := new(controllers.IndexController)
	apiauth := new(controllers.ApiAuth)
  apiAdmin := new(controllers.ApiAdmin)
  apiAccount := new(controllers.ApiAccount)
  apiWallet := new(controllers.ApiWallet)
  apiContract := new(controllers.ApiContract)
  apiAsyncTask := new(controllers.ApiAsyncTask)
  apiLoyalty := new(controllers.ApiLoyalty)

  router.Get("/", index.Index)

  //4. Nhóm API hệ thống
  auth := router.Group("/auth")
  {
      auth.Post("/loadjwt", apiauth.LoadJwtAccount)
      auth.Post("/login", apiauth.Login)
  }
	authorized := router.Group("/api")

  // Authentication required
	authorized.Use(middleware.JWTMiddleware())
	{
      //5. Admin
			admin := authorized.Group("/admin")
			{
          // admin.Post("/log",apiAdmin.GetLog)
          network := admin.Group("/network")
          {
             network.Get("/addpeers",apiAdmin.AddPeers)
          }
          cache := admin.Group("/cache")
          {
             cache.Post("/clear",apiAdmin.ClearCache)
          }

          transaction := admin.Group("/transaction")
          {
             transaction.Get("/<hash>",apiAdmin.GetTransactionByHash)
          }
          account := admin.Group("/account")
          {
             account.Post("/keystore",apiAccount.LoadKeystore)
             account.Post("/loadfile",apiAccount.LoadFile)
             account.Post("/savefile",apiAccount.SaveFile)
             account.Get("/tokens",apiAccount.ListToken)
          }

          wallet := admin.Group("/wallet")
          {
             wallet.Post("/load",apiWallet.Load)
             wallet.Post("/list",apiWallet.List)
             wallet.Post("/new",apiWallet.New)
             wallet.Get("/save",apiWallet.Save)
             wallet.Get("/fillgas",apiWallet.FillGas)
             wallet.Get("/balance/<address>",apiWallet.Balance)
          }
          contract := admin.Group("/contract")
          {
             contract.Post("/deploy",apiContract.Deploy)
             contract.Post("/load",apiContract.Load)
             contract.Get("/view",apiContract.View)
             contract.Post("/save",apiContract.Save)
             contract.Post("/code",apiContract.Code)
          }
          //6. Other tasks
          async_task := authorized.Group("/task")
          {
             async_task.Post("",apiAsyncTask.Task)
             async_task.Post("/status/<taskid>",apiAsyncTask.Status)
          }
			}
      v1 := authorized.Group("/v1")
      {
        //1. Nhóm API liên quan tới tài khoản điểm
        accounts := v1.Group("/accounts")
        {
          accounts.Post("",apiLoyalty.StashCreate)
          accounts.Get("/<account_id>",apiLoyalty.StashInfo)
          accounts.Post("/<account_id>",apiLoyalty.StashUpdate)
        }

        //2. Nhóm API liên quan tới giao dịch số dư
        v1.Post("/credit", apiLoyalty.StashCredit)
        v1.Post("/debit", apiLoyalty.StashDebit)
        v1.Post("/transfer", apiLoyalty.StashTransfer)
        v1.Post("/revert", apiLoyalty.StashRevert)

        //3.Nhóm API liên quan tới tra cứu giao dịch
        v1.Get("/history/<txn_type>/<trace_no_org>", apiLoyalty.HistoryTrace)
        v1.Get("/history/time/<txn_type>/<from_time>/<to_time>", apiLoyalty.HistoryTime)

        v1.Get("/stats/<from_time>/<to_time>", apiLoyalty.StatsAll)
      }
      v2 := authorized.Group("/v2")
      {
          v2.Post("",apiLoyalty.ProcessRequest)
      }
	}
	return router
}
