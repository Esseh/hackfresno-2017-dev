package DISCUSSION
import (
	"fmt"
	"strconv"
	"github.com/Esseh/retrievable"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
	"github.com/Esseh/hackfresno-2017-dev/CONTEXT"
)
type Pair struct{
	PosterName, Post string
}

type S struct{
	Posts []Pair
}

// API Response
func (s *S) GetPosts(ctx CONTEXT.S){
	response := "["
	for i,v := range s.Posts {
		response += `{"poster":"`+v.PosterName+`","post":`+v.Post+`}`
		if i != len(s.Posts) - 1 {
			response += ","
		}
	}
	fmt.Fprint(ctx.Res,response+"]")
}

// Side Effect
func (s *S) MakePost(ctx CONTEXT.S){
	postername := ctx.User.Email
	post := ctx.Req.FormValue("post")
	if post == "" {
		fmt.Fprint(ctx.Res,`{"success":"false","reason":"empty field"}`)
		return
	}
	s.Posts = append(s.Posts,Pair{postername,post})
	id, _ := strconv.ParseInt(ctx.Req.FormValue("discussionID"),10,64)
	retrievable.PlaceEntity(ctx,id,s)
}

func (s *S) Key(ctx context.Context, k interface{}) *datastore.Key {
	return datastore.NewKey(ctx, "DiscussionInstance", "", k.(int64), nil)
}