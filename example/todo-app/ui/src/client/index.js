export default class TodoClient {
  constructor() {
    this.BASE_URL = typeof process.env.API_URL === "undefined" ? '/api' : process.env.API_URL
  }

  listTodo() {
    return fetch(this.BASE_URL + "/todos")
  }

  createTodo(title) {
    let request = {
      method: 'POST',
      body: JSON.stringify({ title: title })
    }
    return fetch(this.BASE_URL + '/todos', request);
  }

  updateTodo(id, completed) {
    let request = {
      method: 'PUT',
      body: JSON.stringify({ completed: completed })
    }
    return fetch(this.BASE_URL + '/todos' + '/' + id, request);
  }

  deleteTodo(id) {
    return fetch(this.BASE_URL + '/todos' + '/' + id, { method: 'DELETE' })
  }
}

