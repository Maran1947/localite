
# Localite
Localite is an **AI-powered Go CLI** tool designed to empower developers by automating commit messages and managing local dev environment configurations.

## 🚀 Features

- **AI Commit Message Generation**: Automatically generate meaningful and descriptive commit messages based on your code changes.
- **Manage Keys Configuration**: Easily manage your configuration keys.
- **New Features Coming Soon**: Stay tuned for more exciting features and enhancements!


## 🛠 Built With

- **Go**: The programming language that powers Localite.
- **Cobra**: A powerful library for creating command-line applications in Go.
- **Gemini AI**: Leveraging advanced AI capabilities for intelligent commit message generation.


## 📦 Installation

To install Localite, you can clone the repository and build it using Go:

```bash
git clone https://github.com/maran1947/localite.git
cd localite
go build
go install
```

Alternatively, you can download the pre-built binaries from the [releases page](https://github.com/maran1947/localite/releases).

---

## 📝 Usage

### Command Overview
For detailed usage, run `localite -h` or `localite --help`.

```bash
localite [flags]
localite [command]
```

### Available Commands

- `localite config` - Manage your Localite configuration.
- `localite genc` - Generate commit messages based on your git diff.

### Example Usage

Generate a commit message:

```bash
localite genc -pl 50
```
- `-p` / `--prefix`: enable conventional prefix for the commit message (e.g., feat, fix) to categorize the change.

- `-l` / `--length`: Sets the maximum length of the generated commit message (e.g., -l 50 limits it to 50 characters).

View configuration settings:

```bash
localite config --list
```

## 🤝 Contributing

We welcome contributions to improve Localite! Please follow these steps to contribute:

1. Fork and clone the repository.
2. Create a new branch (`git checkout -b <branch_name>`).
3. Make your changes and commit them using Localite commands (`localite genc -pl 100'`) instead of the traditional `git commit` command. This way, you can take advantage of the AI-powered commit message generation that Localite offers!

4. Push to the branch (`git push origin <branch_name>`).
5. Create a pull request.

## ✨ Support Localite’s Development

Localite is here to make your local development easier! Your support helps us add new features so developers can spend more time building and less time on repetitive tasks.

By supporting Localite, you’re backing an AI-powered CLI tool that simplifies dev environment management and automates commit messages—making life easier for developers.

If you believe in Localite’s mission and find it helpful, please consider showing your support! Whether it’s starring the repo, contributing, or sharing with others, _every bit counts_. Together, we can create something remarkable.

☕️ **[Buy Me a Coffee](https://www.buymeacoffee.com/abhishekmaran)**  
⭐️ **Don’t forget to star the repo on GitHub!**

Thank you for joining us on this journey! Together, we have a long way to go!

## 🛡️ License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.