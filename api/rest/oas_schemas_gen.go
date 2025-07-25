// Code generated by ogen, DO NOT EDIT.

package rest

// Ref: #/components/schemas/RequestTodo
type RequestTodo struct {
	Content string `json:"content"`
}

// GetContent returns the value of Content.
func (s *RequestTodo) GetContent() string {
	return s.Content
}

// SetContent sets the value of Content.
func (s *RequestTodo) SetContent(val string) {
	s.Content = val
}

// Ref: #/components/schemas/Todo
type Todo struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

// GetID returns the value of ID.
func (s *Todo) GetID() int64 {
	return s.ID
}

// GetContent returns the value of Content.
func (s *Todo) GetContent() string {
	return s.Content
}

// GetCompleted returns the value of Completed.
func (s *Todo) GetCompleted() bool {
	return s.Completed
}

// SetID sets the value of ID.
func (s *Todo) SetID(val int64) {
	s.ID = val
}

// SetContent sets the value of Content.
func (s *Todo) SetContent(val string) {
	s.Content = val
}

// SetCompleted sets the value of Completed.
func (s *Todo) SetCompleted(val bool) {
	s.Completed = val
}

// TodosDeleteTodoNoContent is response for TodosDeleteTodo operation.
type TodosDeleteTodoNoContent struct{}
