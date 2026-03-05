package example

import goapi "github.com/mituba4154/goapi-sdk"

func onJoin(ctx *goapi.Context) {
	ctx.Player().SendMessage("§aGoAPILib 動作確認 OK！")
}
