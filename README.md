# SmartCommit

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

AI-powered git commit message generator using Google Gemini.

## Features

- AI-generated conventional commit messages
- Follows Conventional Commits specification
- Colored terminal output
- Copy to clipboard support
- Configurable via `.smartcommitrc.json`
- Multiple output options (dry-run, auto-commit)
- Cross-platform support (Linux, macOS, Windows)

## Installation

### From source

```bash
git clone https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit.git
cd Smart-Commit
make install
```

### Using Go

```bash
go install github.com/JOSIAHTHEPROGRAMMER/Smart-Commit@latest
```

### Pre-built binaries

Download the latest release from the [releases page](https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit/releases).

## Setup

### Step 1: Get Gemini API Key

Get a Gemini API key from [Google AI Studio](https://makersuite.google.com/app/apikey)

### Step 2: Configure environment

Create a `.env` file in your project root:

```bash
GEMINI_API_KEY=your_api_key_here
```

### Step 3: Optional configuration

Create a `.smartcommitrc.json` config file:

```json
{
  "default_type": "feat",
  "model": "gemini-3-flash-preview",
  "style": "detailed"
}
```

## Usage

### Basic usage

```bash
# Stage your changes
git add .

# Generate and commit
smartcommit
```

### Command-line flags

| Flag        | Short | Description                                |
| ----------- | ----- | ------------------------------------------ |
| `--dry-run` | `-d`  | Preview commit message without committing  |
| `--copy`    | `-c`  | Copy commit message to clipboard           |
| `--type`    | `-t`  | Force commit type (feat, fix, chore, etc.) |
| `--scope`   | `-s`  | Force commit scope                         |
| `--version` | `-v`  | Show version information                   |
| `--help`    | `-h`  | Show help message                          |

### Examples

```bash
# Dry run (preview only, don't commit)
smartcommit --dry-run

# Copy to clipboard
smartcommit --copy

# Force commit type
smartcommit --type feat

# Force commit scope
smartcommit --scope api

# Combine flags
smartcommit --dry-run --copy --type fix --scope auth
```

## Configuration

### Config file (`.smartcommitrc.json`)

| Option         | Type   | Default                  | Description                                 |
| -------------- | ------ | ------------------------ | ------------------------------------------- |
| `default_type` | string | "chore"                  | Default commit type if not detected         |
| `model`        | string | "gemini-3-flash-preview" | Gemini model to use                         |
| `style`        | string | "detailed"               | Commit message style: "detailed" or "short" |

### Environment variables (`.env`)

| Variable         | Required | Description         |
| ---------------- | -------- | ------------------- |
| `GEMINI_API_KEY` | Yes      | Your Gemini API key |

## Conventional Commit Types

SmartCommit supports the following conventional commit types:

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Test additions or changes
- `chore`: Maintenance tasks
- `perf`: Performance improvements
- `ci`: CI/CD changes
- `build`: Build system changes
- `revert`: Revert previous commit

## Development

### Build commands

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Clean build artifacts
make clean

# Run application
make run
```

### Project structure

```
Smart-Commit/
├── cmd/              # Command definitions
├── internal/
│   ├── ai/          # AI integration (Gemini)
│   ├── config/      # Configuration handling
│   ├── formatter/   # Commit message formatting
│   └── git/         # Git operations
├── .github/
│   └── workflows/   # CI/CD workflows
├── main.go          # Entry point
├── Makefile         # Build scripts
└── README.md
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Security Notice

- Your git diffs are sent to Google Gemini for analysis
- Sensitive data is filtered before transmission
- Review generated messages before committing
- Use `--dry-run` flag to preview without committing
- Avoid using on repositories with sensitive/proprietary code
- Consider self-hosted AI alternatives for sensitive projects
- For Global usage, add /go/bin to your path

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI
- Powered by [Google Gemini](https://ai.google.dev/) AI
- Follows [Conventional Commits](https://www.conventionalcommits.org/) specification
