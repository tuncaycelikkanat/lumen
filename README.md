# Lumen

Lumen is a lightweight, **YAML-based rule engine** designed to run seamlessly in containers or directly on your machine. Its primary purpose is to parse user-defined `rules.yaml` files and execute controls, analysis, or log monitoring based on those rules.

The project is highly portable and has been upgraded to a robust, professional CLI tool.

---

## 🚀 Features

* Reads rules defined in YAML format (supports both `keyword` and `regex` matching).
* Real-time robust log tailing (resistant to log rotation and truncation).
* Monitors single or multiple log files simultaneously.
* Colored and formatted CLI output categorized by severity.
* Exports caught events as CSV or JSON reports.
* Isolated execution environment via Docker.

> **"Define the rule in a file, run it independent of the environment."**

---

## 📦 Requirements

* Go 1.25.6+ (for local development)
* Docker (20.x or higher recommended)

---

## ⚙️ Installation

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/tuncaycelikkanat/lumen.git
cd lumen
```

### 2️⃣ Build Locally

```bash
go mod tidy
go build -o lumen ./cmd
```

---

## 🚀 Usage (CLI)

Lumen uses a modern CLI structure. You can view all options using the `--help` flag:

```bash
./lumen --help
```

### Start Monitoring

To start monitoring logs:

```bash
./lumen start --config rules.yaml --log /var/log/syslog --log /var/log/auth.log --format json
```

#### Flags:
* `-c, --config` : Path to the rules configuration file (default: `rules.yaml`)
* `-l, --log`    : Log file(s) to monitor. Can be used multiple times. (default: `auth.log`)
* `-f, --format` : Report export format upon exit (`csv` or `json`). (default: `csv`)

To gracefully stop monitoring and save the report, simply press `Ctrl+C`.

---

## 📝 rules.yaml Example

```yaml
rules:
  - id: 1
    name: "SSH Brute Force"
    keyword: "Failed password"
    severity: "High"
  - id: 2
    name: "Email Found"
    regex: "[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}"
    severity: "Info"
```

---

## 🐳 Docker Deployment

### Build the Image

```bash
docker build -t lumen:latest .
```

### Run with Docker

```bash
docker run --rm -it \
  -v $(pwd)/rules.yaml:/app/rules.yaml \
  -v /var/log/auth.log:/var/log/auth.log \
  lumen:latest start --config /app/rules.yaml --log /var/log/auth.log
```

---

## 🤝 Contributing

Contributions are welcome! 🙌

1. Fork the project
2. Create a feature branch (`feature/my-feature`)
3. Commit your changes
4. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License.

---

## ✨ Contact

Feel free to reach out via GitHub Issues for any feedback or suggestions.
