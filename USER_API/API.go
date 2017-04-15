package USER_API
import(
	"fmt"
	"time"
	"strconv"
	"github.com/Esseh/retrievable"
	"github.com/Esseh/hackfresno-2017-dev/GROUPS"
	"github.com/Esseh/hackfresno-2017-dev/LOGIN"
	"github.com/Esseh/hackfresno-2017-dev/CONTEXT"
)
func CreateGroupAPI(ctx CONTEXT.S){
	l := LOGIN.S{}
	retrievable.GetEntity(ctx,ctx.User.Email,&l)
	keyword1 := ctx.Req.FormValue("university")
	keyword2 := ctx.Req.FormValue("field")
	keyword3 := ctx.Req.FormValue("subject")
	about    := ctx.Req.FormValue("about")
	key , _ := retrievable.PlaceEntity(ctx,int64(0),&GROUPS.S{
		LastUpdated: time.Now(),
		Keyword1: keyword1,
		Keyword2: keyword2,
		Keyword3: keyword3,
		About: about,
	})
	ctx.User.MyGroupIDS = append(ctx.User.MyGroupIDS,key.IntID())
	retrievable.PlaceEntity(ctx,l.UserID,&ctx.User)
}

func MyGroupsAPI(ctx CONTEXT.S){
	response := "["
	for i,v := range ctx.User.MyGroupIDS {
	
	}
	fmt.Fprint(ctx.Res,response+"]")
}

func FindGroupsAPI(ctx CONTEXT.S){
	
}