# EnvSafe: Secure Environment Variable Manager

<div align="center">

![EnvSafe Logo](https://via.placeholder.com/150x150.png?text=EnvSafe)

[![Go Version](https://img.shields.io/github/go-mod/go-version/vivekjha1213/EnvSafe)](https://github.com/vivekjha1213/EnvSafe/blob/main/go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/vivekjha1213/EnvSafe)](https://goreportcard.com/report/github.com/vivekjha1213/EnvSafe)
[![GoDoc](https://godoc.org/github.com/vivekjha1213/EnvSafe?status.svg)](https://godoc.org/github.com/vivekjha1213/EnvSafe)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Open Source](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://opensource.org/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Build Status](https://github.com/vivekjha1213/EnvSafe/workflows/Go/badge.svg)](https://github.com/vivekjha1213/EnvSafe/actions)
[![codecov](https://codecov.io/gh/vivekjha1213/EnvSafe/branch/main/graph/badge.svg)](https://codecov.io/gh/vivekjha1213/EnvSafe)
[![Go Reference](https://pkg.go.dev/badge/github.com/vivekjha1213/EnvSafe.svg)](https://pkg.go.dev/github.com/vivekjha1213/EnvSafe)
[![Made with Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub issues](https://img.shields.io/github/issues/vivekjha1213/EnvSafe.svg)](https://github.com/vivekjha1213/EnvSafe/issues)
[![GitHub stars](https://img.shields.io/github/stars/vivekjha1213/EnvSafe.svg)](https://github.com/vivekjha1213/EnvSafe/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/vivekjha1213/EnvSafe.svg)](https://github.com/vivekjha1213/EnvSafe/network)

</div>

EnvSafe is a powerful command-line tool designed to securely manage environment variables. It offers robust functionality for setting, retrieving, exporting, and loading environment variables with built-in encryption support.

## Features

- 🔐 Encrypt and decrypt environment variables
- 💾 Store variables in a local encrypted file
- 🚀 Automatically load variables into the shell on startup
- 🌍 Support for multiple environments (e.g., dev, staging, prod)

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/vivekjha1213/EnvSafe.git
   cd EnvSafe
   ```

2. **Build the Project**
   ```bash
   go build -o envsafe ./cmd/envsafe
   ```
   This command creates an executable named `envsafe` in the project directory.

## Usage

### Commands

- `set`: Set a secret in the local store
- `get`: Retrieve a secret from the local store
- `load-env`: Load secrets from environment variables into the local store
- `export-env`: Export secrets from the local store to environment variables

### Flags

- `-f`, `--file`: Path to the secrets store file (default is `secrets.json`)
- `-k`, `--key`: Encryption key for securing secrets

### Examples

1. **Set a Secret**
   ```bash
   ./envsafe set <key> <value> --key=<encryption_key>
   ```
   Example:
   ```bash
   ./envsafe set my_secret_key my_secret_value --key=12345678901234567890123456789012
   ```

2. **Get a Secret**
   ```bash
   ./envsafe get <key> --key=<encryption_key>
   ```
   Example:
   ```bash
   ./envsafe get my_secret_key --key=12345678901234567890123456789012
   ```

3. **Load Secrets from Environment Variables**
   ```bash
   export MYSECRETKEY_my_secret_key=my_secret_value
   ./envsafe load-env MYSECRETKEY
   ```
   This command loads secrets with the prefix `MYSECRETKEY` from environment variables and saves them to `secrets.json`.

4. **Export Secrets to Environment Variables**
   ```bash
   ./envsafe export-env MYSECRETKEY
   ```
   This command exports secrets from `secrets.json` to environment variables with the prefix `MYSECRETKEY`.

5. **Viewing Help Information**
   ```bash
   ./envsafe --help
   ./envsafe set --help
   ./envsafe get --help
   ./envsafe load-env --help
   ./envsafe export-env --help
   ```

## Development

To contribute to the project, follow these steps:

1. **Fork the Repository** and create a new branch
2. **Make Your Changes** and ensure they are well-tested
3. **Submit a Pull Request** with a clear description of the changes

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

For questions, support, or to report issues, please open an issue on the GitHub repository or contact the author:

- **Vivek Kumar Jha**: [GitHub Profile](https://github.com/vivekjha1213)

---

<div align="center">

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/vivekjha1213/EnvSafe)
[![Go Report Card](https://goreportcard.com/badge/github.com/vivekjha1213/EnvSafe)](https://goreportcard.com/report/github.com/vivekjha1213/EnvSafe)
[![codecov](https://codecov.io/gh/vivekjha1213/EnvSafe/branch/main/graph/badge.svg)](https://codecov.io/gh/vivekjha1213/EnvSafe)

EnvSafe is built with Go and utilizes open-source Go packages to provide a professional-grade solution for secure environment variable management.

</div>
