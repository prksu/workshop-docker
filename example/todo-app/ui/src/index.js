import TodoClient from './client';
import './styles.css';

customElements.define('x-app', class extends HTMLElement {
  constructor() {
    super();
    this.client = new TodoClient
  }
  
  TodoList() {
    this.client.listTodo().then((response) => { return response.json() }).then((data) => {
      if (data.todos == null) { return }
      this.todoList.innerHTML = ''
      data.todos.forEach((todo) => {
        this.TodoItem(todo)        
      })
    }).catch(() => this.todoList.innerHTML = `<div>Failed fetching data</div>`)
  }

  TodoItem(todo) {
    let item = document.createElement('li')
    let is_completed = todo.completed ? "completed" : ""
    item.innerHTML = todo.title
    item.className = "todo-item " + is_completed
    item.addEventListener('click', () => { 
      this.client.updateTodo(todo.id, !todo.completed)
      item.classList.toggle('completed')
    })
    let removeButton = document.createElement('span')
    removeButton.innerHTML = "X"
    removeButton.addEventListener('click', () => { 
      this.client.deleteTodo(todo.id)
      item.remove()
    })
    item.appendChild(removeButton)
    this.todoList.appendChild(item)
  }

  connectedCallback() {
    this.innerHTML = this.render();
    this.todoInput = this.querySelector('.todo-input')
    this.todoList = this.querySelector('.todo-list')
    this.TodoList()
    this.todoInput.addEventListener('keypress', event => {
      if (event.which == 13) {
        this.client.createTodo(this.todoInput.value).then((response) => { return response.json() }).then((data) => {
          this.TodoItem(data)
        })
        this.todoInput.value = ""
      }
    })
  }
  
  disconnectedCallback() {}

  attributeChangedCallback(name, oldValue, newValue) {}

  adoptedCallback() {}

  render() {
    return `
      <header>
        <h1>Todos App</h1>
      </header>
      <section class="todo-form">
        <input type="text" class="todo-input" placeholder="input your new todo here..." />
      </section>
      <ul class="todo-list"></ul>
    `
  }
})