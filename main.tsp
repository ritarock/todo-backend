import "@typespec/http";

using Http;
@service(#{
  title: "todos service",
})
@server("http://localhost:8080", "Single server endpoiont")

namespace Todo;

model Todo {
  id: int64;
  content: string;
  completed: boolean;
}

model RequestTodo {
  content: string;
}

@route("/todos")
namespace Todos {
  @get
  op listTodos(): {
    @statusCode statusCode: 200;
    @body todos: Todo[];
  };

  @get
  op getTodo(@path todoID: int64): {
    @statusCode statusCode: 200;
    @body todo: Todo;
  };

  @post
  op createTodo(@body todo: RequestTodo): {
    @statusCode statusCode: 201;
    @body newTodo: Todo;
  };

  @put
  op updateTodo(@path todoID: int64, @body todo: Todo): {
    @statusCode statusCode: 200;
    @body updateTodo: Todo;
  };

  @delete
  op deleteTodo(@path todoID: int64): {
    @statusCode statusCode: 204;
  };
}
