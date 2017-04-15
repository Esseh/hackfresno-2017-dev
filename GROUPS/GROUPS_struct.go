package GROUPS
import(
	"fmt"
	"strconv"
	"github.com/Esseh/hackfresno-2017-dev/EVENT"
	"github.com/Esseh/hackfresno-2017-dev/CONTEXT"
	"golang.org/x/net/context"
	"github.com/Esseh/retrievable"
	"google.golang.org/appengine/datastore"
)

type S struct{}

// Internal Function
func (s *S) GetEvent(ctx CONTEXT.S) *EVENT.S {
	i, _ := strconv.ParseInt(ctx.Req.FormValue("eventID"),10,64)
	e := &EVENT.S{}
	retrievable.GetEntity(ctx,i,e)
	return e
}


func (s *S) Key(ctx context.Context, k interface{}) *datastore.Key {
	return datastore.NewKey(ctx, "EventsContainer", "", k.(int64), nil)
}