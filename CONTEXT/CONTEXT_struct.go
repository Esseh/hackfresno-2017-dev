package CONTEXT
import(
	"strconv"
	"github.com/Esseh/hackfresno-2017-dev/USER"
	"github.com/Esseh/retrievable"
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)
type S struct{
	Res http.ResponseWriter
	Req *http.Request
	User USER.S
	context.Context
}

func NewContext(res http.ResponseWriter,req *http.Request) S {
	ctx_a := appengine.NewContext(req)
	u := USER.S{}
	c, err := req.Cookie("login")
	if err != nil {
		userKey, _ := strconv.ParseInt(c.Value,10,64)
		retrievable.GetEntity(ctx,userKey,&u)
	}
	return S{res,req,u,ctx_a}
}