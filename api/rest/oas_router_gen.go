// Code generated by ogen, DO NOT EDIT.

package rest

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos"

			if l := len("/todos"); len(elem) >= l && elem[0:l] == "/todos" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch r.Method {
				case "GET":
					s.handleTodosListTodosRequest([0]string{}, elemIsEscaped, w, r)
				case "POST":
					s.handleTodosCreateTodoRequest([0]string{}, elemIsEscaped, w, r)
				default:
					s.notAllowed(w, r, "GET,POST")
				}

				return
			}
			switch elem[0] {
			case '/': // Prefix: "/"

				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "todoID"
				// Leaf parameter, slashes are prohibited
				idx := strings.IndexByte(elem, '/')
				if idx >= 0 {
					break
				}
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "DELETE":
						s.handleTodosDeleteTodoRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					case "GET":
						s.handleTodosGetTodoRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					case "PUT":
						s.handleTodosUpdateTodoRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "DELETE,GET,PUT")
					}

					return
				}

			}

		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos"

			if l := len("/todos"); len(elem) >= l && elem[0:l] == "/todos" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch method {
				case "GET":
					r.name = TodosListTodosOperation
					r.summary = ""
					r.operationID = "Todos_listTodos"
					r.pathPattern = "/todos"
					r.args = args
					r.count = 0
					return r, true
				case "POST":
					r.name = TodosCreateTodoOperation
					r.summary = ""
					r.operationID = "Todos_createTodo"
					r.pathPattern = "/todos"
					r.args = args
					r.count = 0
					return r, true
				default:
					return
				}
			}
			switch elem[0] {
			case '/': // Prefix: "/"

				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "todoID"
				// Leaf parameter, slashes are prohibited
				idx := strings.IndexByte(elem, '/')
				if idx >= 0 {
					break
				}
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "DELETE":
						r.name = TodosDeleteTodoOperation
						r.summary = ""
						r.operationID = "Todos_deleteTodo"
						r.pathPattern = "/todos/{todoID}"
						r.args = args
						r.count = 1
						return r, true
					case "GET":
						r.name = TodosGetTodoOperation
						r.summary = ""
						r.operationID = "Todos_getTodo"
						r.pathPattern = "/todos/{todoID}"
						r.args = args
						r.count = 1
						return r, true
					case "PUT":
						r.name = TodosUpdateTodoOperation
						r.summary = ""
						r.operationID = "Todos_updateTodo"
						r.pathPattern = "/todos/{todoID}"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}

			}

		}
	}
	return r, false
}
