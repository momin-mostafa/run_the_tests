# run_my_tests

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)  
> A lightweight Go-based CLI tool to execute all Flutter tests (unit, widget, and integration) **in parallel** for faster feedback and CI pipelines.

---

## 📌 Overview

`run_my_tests` is a command-line utility written in Go that compiles into a single executable.  
It detects and runs your Flutter **unit tests**, **widget tests**, and **integration tests** simultaneously, leveraging parallelism to reduce test execution time.

This tool is designed for:
- **Local development** — get faster feedback when running tests.
- **CI/CD pipelines** — optimize Flutter test jobs by running them in parallel.

---

## 🚀 Features

- Runs **multiple Flutter test suites** in parallel.
- Compatible with **macOS**, **Linux**, and **Windows**.
- Generates a **single portable binary** (`run_my_tests`).
- Simple CLI — no external dependencies besides Flutter & Go.
- Supports **custom test file paths**.
- Exits with non-zero code if any test fails (CI-friendly).

---

## 📦 Installation

### 1. Clone the repository
```bash
git clone https://github.com/your-username/run_my_tests.git
cd run_my_tests
```

### 2. Build the executable
Make sure you have Go installed (`go version` ≥ 1.18):
```bash
go build -o run_my_tests
```

---

## 🛠 Usage

### Run all tests in parallel
```bash
./run_my_tests
```

By default, it runs:
```bash
flutter test test/unit_test/unit_test.dart
flutter test test/widget_test/widget_test.dart
flutter test integration_test/integration_test.dart
```

### Run with custom test files
```bash
./run_my_tests --unit test/my_unit_test.dart --integration integration_test/login_flow_test.dart
```

### CLI Options
| Flag           | Description                      | Default                                      |
|----------------|----------------------------------|----------------------------------------------|
| `--unit`       | Path to unit test file(s)        | `test/unit_test/unit_test.dart`              |
| `--widget`     | Path to widget test file(s)      | `test/widget_test/widget_test.dart`          |
| `--integration`| Path to integration test file(s) | `integration_test/integration_test.dart`     |

---

## 📂 Project Structure

```
run_my_tests/
├── logger
|   ├── logger.go
|
├── main.go
├── go.mod
├── LICENSE
└── README.md
```

---

## ⚡ Example in CI/CD (GitHub Actions)
```yaml
name: Flutter Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: subosito/flutter-action@v2
        with:
          flutter-version: '3.19.0' #your desired flutter version
      - name: Build run_my_tests
        run: go build -o run_my_tests
      - name: Run tests in parallel
        run: ./run_my_tests
```

---

## 📝 License
This project is licensed under the [Apache License 2.0](LICENSE).  
© 2025 Momin Mostafa

---

## 💡 Author
**Momin Mostafa**  
Passionate about Flutter, Go, and building efficient developer tools.
