# Todo CLI Application 📝

A simple and interactive Command Line Interface (CLI) application to manage your todos. This app allows you to add tasks, mark them as completed, delete tasks, and view your todo list in a tabular format.

---
![logo]("C:\Users\MUNIYA\GolandProjects\awesomeProject1\logo.webp")

## Features 🚀

- **Add Tasks**: Quickly add a task to your todo list.
- **Mark as Completed**: Mark tasks as completed by their index.
- **Delete Tasks**: Remove tasks from your list using their index.
- **View List**: Display all tasks in a beautifully formatted table.
- **Persistent Storage**: Todos are saved to a JSON file (`todos.json`) for easy retrieval.

---

## Installation 💻

### Prerequisites
- [Go](https://go.dev/) (version 1.16 or later)

### Clone the Repository
```bash
git clone https://github.com/Muniya-64bit/todo-cli.git
```

### Build the Application
Compile the application to create an executable:
```bash
go build -o main
```

---

## Usage 📚

### Add a Task
```bash
./todo-cli -add "Your task here"
```
Example:
```bash
./todo-cli -add "Complete the Go project"
```

### Mark a Task as Completed
```bash
./todo-cli -complete <task_index>
```
Example:
```bash
./todo-cli -complete 1
```

### Delete a Task
```bash
./todo-cli -delete <task_index>
```
Example:
```bash
./todo-cli -delete 2
```

### View All Tasks
```bash
./todo-cli -list
```

### Help
To view all available commands:
```bash
./todo-cli -h
```

---

## Output Example

Here’s an example of the `-list` command output:

```
┌───┬────────────────────────────────┬──────────┬────────────────┬──────────────────┬
│ # │ Task                           │ Done     │ CreatedAt      │ CompletedAt      │
├───┼────────────────────────────────┼──────────┼────────────────┼──────────────────┼
│ 1 │ Complete the Go project        │ No    │ Wed, 15 Jan 2025 14:00:00 +0000      │                    
│ 2 │ Submit the assignment          │ Yes   │ Wed, 15 Jan 2025 13:00:00 +0000      │                    
└───┴────────────────────────────────┴──────────┴────────────────┴──────────────────┴
```

---

## Configuration ⚙️

The app uses `todos.json` to store your todos. This file will be created in the current working directory if it doesn’t already exist.

---

## Dependencies 📦

This project uses the following library:
- [github.com/alexeyco/simpletable](https://github.com/alexeyco/simpletable) for rendering the tabular output.

To install the dependency:
```bash
go get github.com/alexeyco/simpletable
```

---

## Contributing 🤝

Feel free to fork this repository and submit pull requests for:
- New features
- Bug fixes
- Documentation improvements



---

## Author 🙌

- **Isurumuni** ([Muniya-64bit](https://github.com/Muniya-64bit))

---

## Feedback 💬

If you have any feedback, please open an issue or contact me at [Muniya-64bit](https://github.com/Muniya-64bit).

