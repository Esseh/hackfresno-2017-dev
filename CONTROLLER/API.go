package CONTROLLER
import (
	"time"
	"strconv"
	"github.com/Esseh/retrievable"
	"github.com/Esseh/hackfresno-2017-dev/CONTEXT"
	"github.com/Esseh/hackfresno-2017-dev/GROUPS"
)

func API(ctx CONTEXT.S){
	g := &GROUPS.S{}
	id, _ := strconv.ParseInt(ctx.Req.FormValue("groupID"),10,64)
	retrievable.GetEntity(ctx,id,g)
	var modified bool
	switch ctx.Req.FormValue("control") {
		case "makeEvent":	
			g.Events(ctx).MakeEvent(ctx)
			modified = true
		case "makeThread":
			g.Threads(ctx).MakeThread(ctx)
			modified = true
		case "makePost":
			g.Threads(ctx).GetThread(ctx).MakePost(ctx)
			modified = true
		case "makeDiscussionPost":
			g.Discussion(ctx).MakePost(ctx)
			modified = true
	}
	if modified {
		g.LastUpdated = time.Now()
		retrievable.PlaceEntity(ctx,id,g)
	}
}