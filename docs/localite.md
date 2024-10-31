# Localite CLI Documentation

Localite is a powerful CLI tool designed to streamline local development functionalities, including configuration management and AI-powered commit message generation. This documentation provides a detailed overview of available commands, flags, and usage examples.

---

## 📜 Table of Contents

1. [Usage](#usage)
2. [Available Commands](#available-commands)
   - [Completion](#completion)
   - [Config](#config)
     - [del](#config-del)
     - [get](#config-get)
     - [push](#config-push)
     - [set](#config-set)
   - [Generate Commit (genc)](#generate-commit-genc)
3. [Flags](#flags)
4. [Examples](#examples)

---

### 📘 Usage

To get started with Localite, run:

```bash
localite [flags]
localite [command]
```
Use "localite [command] --help" for more details on each command.

### 🎨 Available Commands

🔹 config
Manage application configurations using key-value pairs.
```bash
localite config [flags]
localite config [command]
```
Available Config Commands:
`set`: Sets a key-value pair in the configuration.
```bash
localite config set <key>=<value>
```
`get`: Retrieves the value for a specified key from the configuration.
```bash
localite config get <key>
```
`del`: Deletes a key-value pair from the configuration.
```bash
localite config del <key>
```
`push`: Saves the provided keys to the specified file.
```bash
localite config push -f <file_path> -k <key1> <key2> ...
``` 

🔹 Generate Commit (genc)
Generates an AI-powered commit message based on current Git changes.
```bash
localite genc [flags]
```

### 🚩 Flags
#### Global Flags:
`-h`, `--help`: Display help information for Localite or a specific command.
`-v`, `--version`: Display the current version of the Localite CLI.

#### Command-Specific Flags:
- Config
    - `-l, --list`: Displays the current configurations for the Localite tool.
- Generate Commit (genc)
    - `-l, --length` <int>: Set the maximum length for the generated commit message.
    - `-p, --prefix`: Add a conventional prefix (e.g., feat, fix) to the beginning of the commit message.

### 💡 Examples

**View Help Information**
```bash
localite --help
```

**Manage Configurations**
- Set a configuration value:
```bash
localite config set API_KEY="your-api-key"
```
- Get a configuration value:
```bash
localite config get API_KEY
```
- Delete a configuration key:
```bash
localite config del API_KEY
```
- Push keys to a file 
```bash
localite config push -f ./secrets.env -k DB_USER DB_PASS
```
### 🤖 Generate an AI-Powered Commit Message
```bash
# First, set a Gemini API key
localite config set GEMINI_API_KEY="your-gemini-api-key" 

# Run one of the following commands to generate a commit message
localite genc -l 50 -p
or
localite genc -pl 50
or
localite genc -l 50
```

---

Thanks for checking out Localite! ⭐ If you enjoyed using it, please consider giving us a star—your support means a lot!

---
