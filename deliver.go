package deliver

// Deliver >> A courier
type Deliver struct {
}

// Context >> Request Context
type Context struct {
}

// Middleware >> Func type
type Middleware func(*Context) bool

// Use >> Add middleware
func (_d *Deliver) Use(_m *Middleware) {
}

// Listen >> Start listening port
func (_d *Deliver) Listen() {
}
