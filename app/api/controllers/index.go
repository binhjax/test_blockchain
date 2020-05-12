package controllers

import (
	"github.com/qiangxue/fasthttp-routing"
  // "git.teko.vn/loyalty-system/loyalty-onchain-projects/api_server/app/server/config"
  // "strings"
)
type IndexController struct {}

func (api *IndexController) Index(c  *routing.Context) error {
		return Response(c,505,"Please use POST method",nil)
}
