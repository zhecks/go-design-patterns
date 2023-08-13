package proxy

import "fmt"

type VirtualProxy struct {
	UserModel
	loaded bool
}

func (vp *VirtualProxy) GetID() string {
	return vp.ID
}

func (vp *VirtualProxy) GetName() string {
	return vp.Name
}

func (vp *VirtualProxy) GetDetail() string {
	if !vp.loaded {
		vp.reload()
	}
	return vp.Detail
}

func (vp *VirtualProxy) reload() {
	fmt.Println("get detail from sql")
	vp.Detail = "detail"
}

type UserModel struct {
	ID     string
	Name   string
	Detail string
}

type UserManager struct {
}

func (um *UserManager) GetUser() VirtualProxy {
	model := UserModel{
		ID:   "1",
		Name: "Alice",
	}
	return VirtualProxy{model, false}
}
