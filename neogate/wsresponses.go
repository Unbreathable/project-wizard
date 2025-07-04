package neogate

import (
	"runtime/debug"
)

func NormalResponse(ctx *Context, data map[string]interface{}) Event {
	return Response(ctx, data, ctx.Instance)
}

func SuccessResponse(ctx *Context) Event {
	return Response(ctx, map[string]interface{}{
		"success": true,
	}, ctx.Instance)
}

func ErrorResponse(ctx *Context, message string, err error) Event {

	if DebugLogs {
		Log.Println("error with action "+ctx.Action+" (", message, "): ", err)
		debug.PrintStack()
	}

	return Response(ctx, map[string]interface{}{
		"success": false,
		"message": message,
	}, ctx.Instance)
}

func Response(ctx *Context, data map[string]interface{}, instance *Instance) Event {
	return Event{
		Name: "res:" + ctx.Action + ":" + ctx.ResponseId,
		Data: data,
	}
}
