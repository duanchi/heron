package reflec

type InvocationHandler interface {
	Invoke(proxy *Proxy, method *Method, args []interface{}) (error)
}