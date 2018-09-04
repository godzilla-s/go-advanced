package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserResource struct {
	users map[string]User
}

func (u UserResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/users").Consumes(restful.MIME_XML, restful.MIME_JSON, restful.MIME_OCTET)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{user-id}").To(u.findUser))
	ws.Route(ws.GET("/name/{user-name}").To(u.findUser))
	ws.Route(ws.POST("").To(u.updateUser))
	ws.Route(ws.PUT("/{user-id}").To(u.createUser))
	ws.Route(ws.DELETE("/{user-id}").To(u.deleteUser))

	container.Add(ws)
}

func (u UserResource) findUser(req *restful.Request, resp *restful.Response) {
	log.Print("find user")
	id := req.PathParameter("user-id")
	user, ok := u.users[id]
	if !ok {
		log.Printf("not found user id:%v", id)
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		resp.WriteEntity(user)
	}
}

func (u *UserResource) updateUser(req *restful.Request, resp *restful.Response) {
	log.Print("update user")
	//user := new(User)
	var user User
	err := req.ReadEntity(&user)
	if err == nil {
		u.users[user.Id] = user
		resp.WriteEntity(&user)
	} else {
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (u *UserResource) createUser(req *restful.Request, resp *restful.Response) {
	log.Print("create user")
	user := User{Id: req.PathParameter("user-id")}
	err := req.ReadEntity(&user)
	if err == nil {
		u.users[user.Id] = user
		resp.WriteHeaderAndEntity(http.StatusCreated, user)
	} else {
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (u *UserResource) deleteUser(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("user-id")
	delete(u.users, id)
}

func main() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	u := UserResource{map[string]User{}}
	u.Register(wsContainer)

	log.Printf("start listen on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

// 用curl测试
// POST:
// 		curl -H "Content-Type:application/json" -X POST -d '{"id":"ug221","name":"james"}' http://localhost:8080/users
// PUT:
// 		curl -H "Content-Type:application/json" -X PUT -d '{"name":"harris"}' http://localhost:8080/users/pe309
// DELETE:
// 		curl curl -X DELETE http://localhost:8080/users/ug221
// GET:
// 		curl http://localhost:8080/pe309
