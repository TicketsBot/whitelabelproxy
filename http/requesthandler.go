package http

import (
	"encoding/json"
	"github.com/TicketsBot/whitelabelproxy/patreon"
	"github.com/gin-gonic/gin"
)

func RequestHandler(ctx *gin.Context) {
	signature := ctx.GetHeader("x-patreon-signature")

	body, err := ctx.GetRawData(); if err != nil {
		ctx.String(400, err.Error())
		return
	}

	if !patreon.VerifyHmac(body, []byte(signature)) {
		ctx.String(403, "Invalid signature")
		return
	}

	var pledge patreon.Pledge
	if err = json.Unmarshal(body, &pledge); err != nil {
		ctx.String(400, err.Error())
		return
	}

	event := patreon.Event(ctx.GetHeader("x-patreon-event"))
	switch event {
	case patreon.Create: patreon.HandleCreate(pledge)
	case patreon.Update: patreon.HandleUpdate(pledge)
	case patreon.Delete: patreon.HandleDelete(pledge)
	default: ctx.String(400, "Invalid event")
	}
}
