package router

type Router struct {
	routerRules []RouterRule
}

func (r *Router) Set(routerRule *RouterRule) {
	r.routerRules = append(r.routerRules, routerRule)
}