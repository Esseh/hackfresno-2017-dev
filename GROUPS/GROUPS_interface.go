package GROUPS
import(
	"fmt"
	"time"
	"strconv"
	"github.com/Esseh/hackfresno-2017-dev/THREADS"
	"github.com/Esseh/hackfresno-2017-dev/DISCUSSION"
	"github.com/Esseh/hackfresno-2017-dev/EVENTS"
	"github.com/Esseh/hackfresno-2017-dev/CONTEXT"
	"golang.org/x/net/context"
	"github.com/Esseh/retrievable"
	"google.golang.org/appengine/datastore"
)

type S struct {
	Title string
	LastUpdated time.Time
	Keyword1, Keyword2, Keyword3 string
	About string
}

func (s *S)Discussion(ctx CONTEXT.S) *DISCUSSION.S {
	o := &DISCUSSION.S{}
	id, _ := strconv.ParseInt(ctx.Req.FormValue("groupID"),10,64)
	retrievable.GetEntity(ctx,id,o)
	return o
}
func (s *S)Threads(ctx CONTEXT.S) *THREADS.S {
	o := &THREADS.S{}
	id, _ := strconv.ParseInt(ctx.Req.FormValue("groupID"),10,64)
	retrievable.GetEntity(ctx,id,o)
	return o
}
func (s *S)Events(ctx CONTEXT.S) *EVENTS.S {
	o := &EVENTS.S{}
	id, _ := strconv.ParseInt(ctx.Req.FormValue("groupID"),10,64)
	retrievable.GetEntity(ctx,id,o)
	return o
}
func (s *S)GetAbout(ctx CONTEXT.S){
	fmt.Fprint(ctx.Res,`{"about":"`+s.About+`"}`)
}

func (s *S)LastModified(ctx CONTEXT.S){
	fmt.Fprint(ctx.Res,strconv.FormatInt(s.LastUpdated.UnixNano(),10))
}

func (s *S) Key(ctx context.Context, k interface{}) *datastore.Key {
	return datastore.NewKey(ctx, "GroupInstance", "", k.(int64), nil)
}