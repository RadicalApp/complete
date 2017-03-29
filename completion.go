package complete

// Completionable interface is used to implement your own mobile-compatible completion handlers.
type Completionable interface {
	OnSuccess(*Result)
	OnFailure(string)
}

type successHandler func(r *Result)
type failureHandler func(m string)

// CompletionHandler is a Go-compatible completion handler that implements Completionable.
type CompletionHandler struct {
	success successHandler
	failure failureHandler
}

// AddHandlers adds success/failure handlers to a completion handler
func (c *CompletionHandler) AddHandlers(success successHandler, failure failureHandler) {
	c.success = success
	c.failure = failure
}

// OnSuccess is called when a task succeeds.
func (c *CompletionHandler) OnSuccess(r *Result) {
	c.success(r)
}

// OnFailure is called when a task fails.
func (c *CompletionHandler) OnFailure(msg string) {
	c.failure(msg)
}

// NewCompletion returns an empty completion handler. AddHandler is required for these handlers.
func NewCompletion(success successHandler, failure failureHandler) *CompletionHandler {
	complete := CompletionHandler{}
	complete.AddHandlers(success, failure)
	return &complete
}
