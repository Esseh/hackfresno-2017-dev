package THREADS
import(
	"fmt"
	"strconv"
	"golang.org/x/net/context"
	"github.com/Esseh/retrievable"
	"github.com/Esseh/hackfresno-2017-dev/DISCUSSION"
	"github.com/Esseh/hackfresno-2017-dev/CONTEXT"
	"google.golang.org/appengine/datastore"
)

type Pair struct{
	ThreadTitle string
	DiscussionID int64
}

type S struct{
	Discussions[]Pair
}

func (s *S) GetThreadTitles(ctx CONTEXT.S){
	response := "["
	for i,v := range s.Discussions {
		response += `{"title":"`+v.ThreadTitle+`","id":`+strconv.FormatInt(v.DiscussionID,10)+`}`
		if i != len(s.Discussions) -1 {
			response += ","
		}
	}
	fmt.Fprint(ctx.Res,response+"]")
}

func (s *S) GetThread(ctx CONTEXT.S) *DISCUSSION.S {
	d := &DISCUSSION.S{}
	id , _ := strconv.ParseInt(ctx.Req.FormValue("discussionID"),10,64)
	retrievable.GetEntity(ctx,id,d)
	return d
}

func (s *S) MakeThread(ctx CONTEXT.S){
	threadTitle := ctx.Req.FormValue("threadTitle")
	if threadTitle == "" {
		fmt.Fprint(ctx.Res,`{"success":"false","reason":"empty field"}`)
		return	
	}
	key, err := retrievable.PlaceEntity(ctx,int64(0),&DISCUSSION.S{})
	if err != nil {
		fmt.Fprint(ctx.Res,`{"success":"false","reason":"database error"}`)
		return
	}
	s.Discussions = append([]Pair{Pair{threadTitle,key.IntID()}},s.Discussions...)
	id , _ := strconv.ParseInt(ctx.Req.FormValue("groupID"),10,64)
	retrievable.PlaceEntity(ctx,id,s)
	fmt.Fprint(ctx.Res,`{"success":true}`)
}

func (s *S) Key(ctx context.Context, k interface{}) *datastore.Key {
	return datastore.NewKey(ctx, "Users", "", k.(int64), nil)	
}