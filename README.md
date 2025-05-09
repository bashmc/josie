# 📁 Josie – A File Sharing Platform

**Josie** file sharing platform inspired by Google Drive. Built with a **Vue** frontend and a backend written in **Go**, Josie lets users upload, manage, and share files.

---

> **Note**:🧰🧰🧰This project is currently under active development.🧰🧰🧰

## 🚀 Features

- 📂 Upload and manage files
- 🔗 Generate shareable links for file access
- 🔒 Secure storage and authentication
- 🧭 Fast search and filtering

---

## 🧱 Tech Stack

| Layer        | Technology    |
|--------------|---------------|
| Frontend     | [Vue](https://vuejs.org) |
| Backend      | [Go](https://go.dev) |
| Database     | Postgres |
| Auth         | JWT |
| File Storage | Google Cloud Storage |

---

## 📦 Getting Started

### Prerequisites

- Go 1.24+
- Node.js 20+
- Postgres
- [Goose](https://github.com/pressly/goose)
- [Taskfile](https://github.com/go-task/task)

### 🛠 Run the app

1. Clone the repo and navigate to the backend directory:
   ```bash
   git clone https://github.com/fatcmd/josie.git
   cd josie
    ```
2. Run the app
    ```bash
    task run
    ```