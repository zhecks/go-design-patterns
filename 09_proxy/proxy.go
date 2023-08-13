package proxy

type Object struct {
}

func (object *Object) Do() string {
	return "real"
}

type Proxy struct {
	object *Object
}

func (proxy *Proxy) Do() string {
	var res string
	// do before
	res += "before:"
	res += proxy.object.Do()
	// afrer do
	res += ":after"
	return res
}
