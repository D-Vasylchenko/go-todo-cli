# 📝 Go TODO CLI Application

![Go](https://img.shields.io/badge/Go-1.22-blue?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Working%20Prototype-yellowgreen?style=for-the-badge)
![PRs](https://img.shields.io/badge/Pull%20Requests-Required-blueviolet?style=for-the-badge)
![CI](https://img.shields.io/github/actions/workflow/status/D-Vasylchenko/go-todo-cli/go.yml?label=CI%20Tests&style=for-the-badge&logo=github)
![GitHub Repo Size](https://img.shields.io/github/repo-size/D-Vasylchenko/go-todo-cli?style=for-the-badge)

---

## 📌 Опис застосунку

Цей проєкт — консольний TODO-додаток, реалізований мовою Go.  
Він дозволяє створювати, переглядати, виконувати, видаляти та очищувати задачі, зберігаючи їх у локальному JSON-файлі.

---

## 🛠 Основні можливості

- ✅ Додавання задач: `go run ./cmd/main.go add "Назва"`
- ✅ Перегляд задач: `go run ./cmd/main.go list`
- ✅ Відмітка виконання: `go run ./cmd/main.go done <id>`
- ✅ Видалення задачі: `go run ./cmd/main.go delete <id>`
- ✅ Очищення виконаних: `go run ./cmd/main.go clear`
- 💾 Зберігання у `tasks.json`

---

## 📄 Структура проєкту

```
go-todo-cli/
│
├── cmd/              # Основна точка входу
│   └── main.go
├── todo/             # Логіка об'єктів (тип Task)
│   └── task.go
├── storage/          # JSON-персистенція
│   └── storage.go
├── tests/            # Юніт-тести
│   └── task_test.go
├── tasks.json        # Дані задач
├── DESIGN.md         # 📄 Дизайн-документ (див. нижче)
└── README.md
```

---

## 🔗 Посилання

- 📐 [Design Document](./DESIGN.md)

---

## ▶️ Як зібрати та запустити

### Вимоги:
- Go 1.22+
- Git

### Клонуй проєкт:
```bash
git clone https://github.com/D-Vasylchenko/go-todo-cli.git
cd go-todo-cli
```

### Запуск прикладів:
```bash
go run ./cmd/main.go add "Написати звіт"
go run ./cmd/main.go list
go run ./cmd/main.go done 1
go run ./cmd/main.go delete 1
go run ./cmd/main.go clear
```

---

## ✅ Запуск тестів

```bash
go test ./tests/...
```

🧪 Повертає результат для всіх доступних функцій.

---

## 🖼️ Інтерфейс програми



📌 **запуск додавання задачі**:
```
go run ./cmd/main.go add "Завершити лабораторну"
```

🔲![image](https://github.com/user-attachments/assets/11994a51-d71b-4ecd-ae1a-443838f0ea46)

---

📌 **перегляд задач**:
```
go run ./cmd/main.go list
```

🔲![image](https://github.com/user-attachments/assets/02b13aba-86ab-4f36-958b-e8497f767b18)


---

📌 **запуск тестів**:
```
go test ./tests/...
```

🔲![image](https://github.com/user-attachments/assets/41bc2b44-5ec2-4471-b906-734e67e55b0b)


---

## 🤝 Командна робота

Цей проєкт реалізований у рамках лабораторної №4 «Командна робота» з дисципліни **Методології та технології розробки ПЗ**.  
Проєкт реалізовано з дотриманням GitFlow: кожна функціональність — окрема гілка, ревʼю — через Pull Request, CI — через GitHub Actions.

---

## 📜 Ліцензія

Проєкт розповсюджується під ліцензією [MIT](LICENSE).
