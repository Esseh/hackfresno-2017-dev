package USER
import(
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)
type S struct{
	Username, Email string
	MyGroupIDS []int64
}

func (s *S) Key(ctx context.Context, k interface{}) *datastore.Key {
	return datastore.NewKey(ctx, "Users", "", k.(int64), nil)	
}