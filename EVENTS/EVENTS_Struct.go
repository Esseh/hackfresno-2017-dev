package EVENTS
import(
	"fmt"
	"strconv"
	"github.com/Esseh/hackfresno-2017-dev/EVENT"
	"github.com/Esseh/hackfresno-2017-dev/CONTEXT"
	"golang.org/x/net/context"
	"github.com/Esseh/retrievable"
	"google.golang.org/appengine/datastore"
)

type Pair struct{
	EventTitle string
	EventID int64
}

type S struct{
	Events[]Pair
}

// Respond to API
func (s *S) GetEvents(ctx CONTEXT.S){
	response := "["
	for i,v := range s.Events {
		response += `{"title":"`+v.EventTitle+`","id":`+strconv.FormatInt(v.EventID,10)+`}`
		if i != len(s.Events) - 1 {
			response += ","
		}
	}
	fmt.Fprint(ctx.Res,response+"]")
}

// Internal Function
func (s *S) GetEvent(ctx CONTEXT.S) *EVENT.S {
	i, _ := strconv.ParseInt(ctx.Req.FormValue("eventID"),10,64)
	e := &EVENT.S{}
	retrievable.GetEntity(ctx,i,e)
	return e
}

// Side Effect Function, Respond to API with result of API
func (s *S) MakeEvent(ctx CONTEXT.S){
	title := ctx.Req.FormValue("title")
	what  := ctx.Req.FormValue("what")
	when  := ctx.Req.FormValue("when")
	where := ctx.Req.FormValue("where")
	if what == "" || when == "" || where == "" || title == "" {
		fmt.Fprint(ctx.Res,`{"success":false,"reason":"empty field"}`)
		return
	}
	key, err := retrievable.PlaceEntity(ctx,int64(0),&EVENT.S{
		What:  what,
		When:  when, 
		Where: where,
	})
	if err != nil {
		fmt.Fprint(ctx.Res,`{"success":false,"reason":"datastore put error"}`)
		return	
	}
	s.Events = append([]Pair{Pair{title,key.IntID()}},s.Events...)
	selfKey, _ := strconv.ParseInt(ctx.Req.FormValue("groupID"),10,64)
	retrievable.PlaceEntity(ctx,selfKey,s)
}


func (s *S) Key(ctx context.Context, k interface{}) *datastore.Key {
	return datastore.NewKey(ctx, "EventsContainer", "", k.(int64), nil)
}