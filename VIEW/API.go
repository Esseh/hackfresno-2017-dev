package VIEW
import (
	"strconv"
	"github.com/Esseh/retrievable"
	"github.com/Esseh/hackfresno-2017-dev/CONTEXT"
	"github.com/Esseh/hackfresno-2017-dev/GROUPS"
)

func API(ctx CONTEXT.S){
	g := &GROUPS.S{}
	id, _ := strconv.ParseInt(ctx.Req.FormValue("groupID"),10,64)
	retrievable.GetEntity(ctx,id,g)
	switch ctx.Req.FormValue("control") {
		case "lastModified":	
			g.LastModified(ctx)
		case "about":
			g.GetAbout(ctx)
		case "events":
			g.Events(ctx).GetEvents(ctx)
		case "eventInfo":
			g.Events(ctx).GetEvent(ctx).GetInfo(ctx)
		case "threadTitles":
			g.Threads(ctx).GetThreadTitles(ctx)
		case "posts":
			g.Threads(ctx).GetThread(ctx).GetPosts(ctx)
		case "discussionPosts":
			g.Discussion(ctx).GetPosts(ctx)
	}
}