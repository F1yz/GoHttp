package http

// import "fmt"

// Session，需要一个map,和上次访问时间（用于Session GC)
type Session interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Delete(key interface{})
	SessionId() string
}

// 维护Session列表，以及将过期的Session清除掉
type SessionHandler interface {
	SessionInit() Session
	SessionDelete(sessionId string)
	SessionGet(sessionId string) Session
	SessionGC(maxLifeTime int64)
}
