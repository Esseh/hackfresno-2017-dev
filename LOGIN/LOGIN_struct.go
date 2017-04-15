package LOGIN
import(
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)
type S struct{
	UserID int64
}

func (s *S) Key(ctx context.Context, k interface{}) *datastore.Key {
	return datastore.NewKey(ctx, "Logins", k.(string), 0, nil)	
}