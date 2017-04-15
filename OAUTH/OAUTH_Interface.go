package OAUTH

import (
	"net/http"
	"strconv"
	"github.com/Esseh/hackfresno-2017-dev/USER"
	"github.com/Esseh/hackfresno-2017-dev/LOGIN"
	"golang.org/x/net/context"
	"github.com/Esseh/retrievable"
)


func Login(ctx context.Context, res http.ResponseWriter, req *http.Request, email string){
	Login := LOGIN.S{}
	if retrievable.GetEntity(ctx,email,Login) != nil {
		// Make a New User
		key, err := retrievable.PlaceEntity(ctx,int64(0),&USER.S{Email:email,})
		if err != nil { return }
		Login.UserID = key.IntID()
		_, err = retrievable.PlaceEntity(ctx,email,&Login)
		if err != nil { return }
	}
	http.SetCookie(res,&Cookie{
		Name: "login",
		Value: strconv.FormatInt(Login.UserID,10),
		HttpOnly: true,
		Path: "/",
	})
	http.Redirect(res,req,"/app",http.StatusSeeOther)
}