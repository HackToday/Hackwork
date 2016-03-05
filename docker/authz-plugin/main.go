package main

import (
	"fmt"
	"github.com/docker/go-plugins-helpers/authorization"
)

type authzdriver struct{}

func (*authzdriver) AuthZReq(req authorization.Request) authorization.Response {
	fmt.Println("Auth Request...")
	if req.RequestMethod == "POST" {
		return authorization.Response{Msg: "Create is not allowed!"}
	} else {
		return authorization.Response{Allow: true}
	}
}

func (*authzdriver) AuthZRes(req authorization.Request) authorization.Response {
	fmt.Println("Auth Resp....")
	return authorization.Response{Allow: true}
}

func main() {
	p := &authzdriver{}
	h := authorization.NewHandler(p)
	fmt.Println("Authz plugin is listening now.")
	h.ServeUnix("root", "test_authz_plugin")

}
