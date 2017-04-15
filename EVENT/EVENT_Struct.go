package EVENT
import(
	"fmt"
	"github.com/Esseh/hackfresno-2017-dev/CONTEXT"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)
type S struct{
	What, When, Where string
}

func (s *S) GetInfo(ctx CONTEXT.S) {
	fmt.Fprint(ctx.Res,`{"what":"`+s.What+`","when":"`+s.When+`","where":"`+s.Where+`"}`)
}

func (s *S) Key(ctx context.Context, k interface{}) *datastore.Key {
	return datastore.NewKey(ctx, "EventInstance", "", k.(int64), nil)
}