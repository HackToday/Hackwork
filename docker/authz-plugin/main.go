package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/go-ini/ini"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/openstack/identity/v2/users"
)

type authzdriver struct{}

func (*authzdriver) AuthZReq(req authorization.Request) authorization.Response {
	fmt.Println("Auth Request...")
	reqPara := []string {req.RequestMethod, req.RequestURI}
	reqPath := strings.Join(reqPara, ",")
	fmt.Println("The request path is: ", reqPath)

	role := rule.getRole(req.RequestMethod, req.RequestURI)
	fmt.Printf("Get the role is %s\n", role)

	//TODO to get user if docker support AuthN
	// Note: This is very correct and secure ways to use keystone authentication
	// as usually need user access info to fetch token and use that token to go
        // further operation.

	user := req.RequestHeaders["X-Auth-User"]
	tenant := req.RequestHeaders["X-Auth-Project"]
	req.User = user
	//TODO check if user has such role
	var roles []users.Role
	pager := users.ListRoles(ks.getKsClient(), tenant, user)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		roleList, err := users.ExtractRoles(page)
		if err != nil {
			fmt.Printf("Got error %v", err)
			return false, nil
		}
		for _, ele := range roleList {
			roles = append(roles, ele)
		}
		if len(roleList) == 0 {
			return false, nil
		} else {
			return true, nil
		}
	})
	if err != nil {
		fmt.Printf("List roles got error %v", err)
		return authorization.Response{Err: "Hit an unexpected error"}
	}
	if req.User != "" {
		for _, ro := range roles {
			if ro.Name == role || ro.Name == "admin" {
				return authorization.Response{Allow: true}
			}
		}
		return authorization.Response{Msg: "User not role, Opearation is not allowed!"}
	} else {
		return authorization.Response{Msg: "User not exist, Opearation is not allowed!"}
	}
}

func (*authzdriver) AuthZRes(req authorization.Request) authorization.Response {
	fmt.Println("Auth Resp....")
	return authorization.Response{Allow: true}
}

type actionRule struct {
	ruleMap map[string]string
}
 
func (ar *actionRule) getRole(action string, path string) string {
	fmt.Printf("Action is : %s,  Path is: %s\n", action, path)
	for key, obj := range dockerObjMap {
		if obj.action == action {
			validPath := obj.pathReg
			fmt.Printf("action is %s, validpath is %s\n", obj.action, validPath)
			if validPath.MatchString(path) {
				fmt.Printf("Matched the route: %s, key: %s\n", path, key)
				return ar.ruleMap[key]
			}
		}
	}
	return ""
}

func loadPolicy(filepath string) *actionRule {
	//TODO check if exist
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Rule File error: %v\n", err)
		os.Exit(1)
	}
	var jsonType map[string]string
	err = json.Unmarshal(file, &jsonType)
	if err != nil {
		fmt.Printf("Decode json failure %v", err)
		os.Exit(1)
	}
	fmt.Println("Unmarshal got %v", jsonType)
	actionTarget := &actionRule{ruleMap: jsonType}
	return actionTarget

}

type objPathReg struct {
	action string
	pathReg *regexp.Regexp
}

var (
	rule *actionRule
	dockerObjMap  map[string]objPathReg
	ks  *keystone
)


func init() {
	rule = loadPolicy(os.Args[1])
	dockerObjMap = map[string]objPathReg {
		"container:create": {"POST", regexp.MustCompile(`/v[0-9|.]+/containers/\w+/start`)},
		"container:getall": {"GET", regexp.MustCompile(`/v[0-9|.]+/containers/json`)},
		"image:getall": {"GET", regexp.MustCompile(`/v[0-9|.]+/images/json`)},

	}

	// TODO find file
	fileName := "keystone.conf"
	fmt.Println("file name is: ", fileName)

	cfg, err := ini.Load(fileName)
	if err != nil {
		panic("Could not load config file")
	}
	hash := cfg.Section("").KeysHash()
	ks = &keystone{hash}
}

func main() {
	p := &authzdriver{}
	h := authorization.NewHandler(p)
	fmt.Println("Authz plugin is listening now.\n")
	h.ServeUnix("root", "test_authz_plugin")
}
