# SmartCommit

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

AI-powered git commit message generator using Google Gemini.

## Overview

SmartCommit is a command-line tool that generates conventional commit messages by analyzing your staged git changes. It connects to a centralized AI server that handles the AI processing, allowing you to use SmartCommit across all your projects without managing API keys in each repository.

**Backend Server:** [SmartCommit AI Server](https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit-AI-Server)

## Features

- AI-generated conventional commit messages
- Follows Conventional Commits specification
- Colored terminal output
- Copy to clipboard support
- Configurable via `.smartcommitrc.json`
- Multiple output options (dry-run, auto-commit)
- Cross-platform support (Linux, macOS, Windows)
- Centralized AI server (no API keys in projects)
- Works with local or cloud-hosted backend

## Prerequisites

- Go 1.25 or higher
- Git
- SmartCommit AI Server (local or deployed)

## Installation

### Option 1: Install from source

```bash
git clone https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit.git
cd Smart-Commit
make install
```

### Option 2: Using Go

```bash
go install github.com/JOSIAHTHEPROGRAMMER/Smart-Commit@latest
```

## Setup

### Step 1: Set up the AI Server

You need a running SmartCommit AI Server. Choose one option:

**Option A: Local Server (Recommended for personal use)**

See the [AI Server repository](https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit-AI-Server) for setup instructions.

Quick start:
```bash
# Clone and set up the server
git clone https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit-AI-Server.git
cd Smart-Commit-AI-Server
npm install
# Add GEMINI_API_KEY to .env
node server.js
```

**Option B: Deploy to Vercel (For remote access)**

Follow the Vercel deployment guide in the [AI Server repository](https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit-AI-Server#deploy-to-vercel).

### Step 2: Configure SmartCommit Client

You have **three options** for configuration. Choose the one that works best for you:

#### Option A: Home Directory Config (Recommended)

Create `.smartcommitrc.json` in your home directory **once** - works for all projects:

**Linux/macOS:**
```bash
nano ~/.smartcommitrc.json
```

**Windows:**
```powershell
notepad ~\.smartcommitrc.json
```

**Content:**
```json
{
  "default_type": "chore",
  "style": "detailed",
  "server_url": "https://your-app.vercel.app",
  "api_key": "your_server_api_key_if_enabled",
  "timeout": 30
}
```

#### Option B: Environment Variables (Global)

Set environment variables **once** - works for all projects:

**Linux/macOS (add to ~/.bashrc or ~/.zshrc):**
```bash
export SMARTCOMMIT_SERVER_URL="https://your-app.vercel.app"
export SMARTCOMMIT_API_KEY="your_server_api_key_if_enabled"
```

**Windows PowerShell (permanent):**
```powershell
[System.Environment]::SetEnvironmentVariable('SMARTCOMMIT_SERVER_URL', 'https://your-app.vercel.app', 'User')
[System.Environment]::SetEnvironmentVariable('SMARTCOMMIT_API_KEY', 'your_server_api_key_if_enabled', 'User')
```

**Windows Command Prompt:**
```cmd
setx SMARTCOMMIT_SERVER_URL "https://your-app.vercel.app"
setx SMARTCOMMIT_API_KEY "your_server_api_key_if_enabled"
```

**IMPORTANT for Windows:** After setting environment variables:
1. Close all PowerShell/Terminal windows completely
2. Reopen PowerShell/Terminal
3. If using VS Code, fully close and reopen VS Code
4. Verify with: `echo $env:SMARTCOMMIT_SERVER_URL`

#### Option C: Per-Project Config

Create `.smartcommitrc.json` in **each project** where you want custom settings:

```json
{
  "default_type": "feat",
  "style": "short",
  "server_url": "http://localhost:8080",
  "api_key": "",
  "timeout": 30
}
```

This overrides home directory config and environment variables for that specific project.

#### Configuration Priority

SmartCommit checks configuration in this order (highest to lowest priority):

1. **Environment variables** (`SMARTCOMMIT_SERVER_URL`, `SMARTCOMMIT_API_KEY`)
2. **Current project directory** (`.smartcommitrc.json`)
3. **Home directory** (`~/.smartcommitrc.json`)
4. **Default values** (`http://localhost:8080`)

### Step 3: Verify Setup

```bash
# Check if smartcommit is installed
smartcommit --version

# Test connection to server
curl http://localhost:8080/health
# or
curl https://your-app.vercel.app/health
```

## Usage

### Basic Usage

```bash
# Stage your changes
git add .

# Generate and commit
smartcommit
```

### Command-line Flags

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
# Preview commit message without committing
smartcommit --dry-run

# Copy commit message to clipboard
smartcommit --copy

# Force specific commit type
smartcommit --type feat

# Force specific scope
smartcommit --scope api

# Combine multiple flags
smartcommit --dry-run --copy --type fix --scope auth
```

## Configuration

### Configuration Options

SmartCommit can be configured in three ways, listed by priority (highest to lowest):

1. **Environment Variables** - Override everything
2. **Project Config File** - Project-specific settings (`.smartcommitrc.json` in project directory)
3. **Home Config File** - Global user settings (`~/.smartcommitrc.json` in home directory)
4. **Default Values** - Built-in fallback values

### Configuration File (`.smartcommitrc.json`)

**Location options:**
- **Home directory** (`~/.smartcommitrc.json`) - Recommended, applies to all projects
- **Project directory** (`.smartcommitrc.json`) - Per-project customization

The configuration file can be placed in:
1. Current project directory (`.smartcommitrc.json`)
2. Home directory (`~/.smartcommitrc.json`)

Priority: Environment variables > Project directory > Home directory > Defaults.

| Option         | Type   | Default                  | Description                                 |
| -------------- | ------ | ------------------------ | ------------------------------------------- |
| `default_type` | string | "chore"                  | Default commit type if not detected         |
| `style`        | string | "detailed"               | Commit message style: "detailed" or "short" |
| `server_url`   | string | "http://localhost:8080"  | AI server URL                               |
| `api_key`      | string | ""                       | Server authentication key (optional)        |
| `timeout`      | number | 30                       | Request timeout in seconds                  |

### Environment Variables

You can also configure via environment variables:

| Variable                  | Description                  |
| ------------------------- | ---------------------------- |
| `SMARTCOMMIT_SERVER_URL`  | Override server URL          |
| `SMARTCOMMIT_API_KEY`     | Override server API key      |

Environment variables take precedence over config file values.

## Conventional Commit Types

SmartCommit supports the following conventional commit types:

- `feat` - New feature
- `fix` - Bug fix
- `docs` - Documentation changes
- `style` - Code style changes (formatting, semicolons, etc.)
- `refactor` - Code refactoring (neither fixes bug nor adds feature)
- `test` - Adding or updating tests
- `chore` - Maintenance tasks (updating dependencies, configs, etc.)
- `perf` - Performance improvements
- `ci` - CI/CD changes
- `build` - Build system changes
- `revert` - Revert previous commit

## Project Structure

```
Smart-Commit/
├── cmd/                  # Command definitions
│   └── root.go          # Main CLI logic
├── internal/
│   ├── ai/              # AI server communication
│   │   └── gemini.go   # HTTP client for server
│   ├── config/          # Configuration handling
│   │   ├── config.go   # Config loading
│   │   └── consent.go  # User consent management
│   ├── formatter/       # Commit message formatting
│   │   └── formatter.go
│   └── git/             # Git operations
│       ├── git.go       # Git commands
│       └── filter.go    # Sensitive data filtering
├── .github/
│   └── workflows/       # CI/CD workflows
├── main.go              # Entry point
├── Makefile             # Build scripts
├── go.mod               # Go module definition
└── README.md
```

## Architecture

```
┌─────────────────┐
│  SmartCommit    │
│  (Go CLI)       │
└────────┬────────┘
         │
         │ HTTP Request
         │ (git diff + options)
         ▼
┌─────────────────┐
│  AI Server      │
│  (Node.js)      │
├─────────────────┤
│  • Auth         │
│  • Rate Limit   │
│  • Validation   │
└────────┬────────┘
         │
         │ API Call
         ▼
┌─────────────────┐
│  Google Gemini  │
│      API        │
└─────────────────┘
```

## Development

### Build Commands

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

# Install globally
make install
```

### Manual Build

```bash
# Build binary
go build -o smartcommit main.go

# Install to system
sudo mv smartcommit /usr/local/bin/
```

### Running Tests

```bash
go test ./...
```

## Troubleshooting

### "Failed to connect to AI server"

**Check server is running:**
```bash
curl http://localhost:8080/health
# or for Vercel
curl https://your-app.vercel.app/health
```

**Verify configuration:**

Check which config is being used:
```bash
# Check environment variables
echo $SMARTCOMMIT_SERVER_URL  # Linux/macOS
echo $env:SMARTCOMMIT_SERVER_URL  # Windows PowerShell

# Check config file in home directory
cat ~/.smartcommitrc.json  # Linux/macOS
cat ~\.smartcommitrc.json  # Windows

# Check config file in current directory
cat .smartcommitrc.json
```

**Windows-specific: Environment variables not working?**

If you set environment variables but they're not working:

1. **Restart your terminal** - Close all PowerShell/Command Prompt windows and reopen
2. **Restart VS Code** - Fully close VS Code (not just the window) and reopen
3. **Verify they're set:**
   ```powershell
   echo $env:SMARTCOMMIT_SERVER_URL
   echo $env:SMARTCOMMIT_API_KEY
   ```
4. **If still not working**, add to VS Code settings:
   
   Create/edit `.vscode/settings.json`:
   ```json
   {
     "terminal.integrated.env.windows": {
       "SMARTCOMMIT_SERVER_URL": "https://your-app.vercel.app",
       "SMARTCOMMIT_API_KEY": "your_api_key"
     }
   }
   ```

**Check network connectivity:**
```bash
ping localhost  # For local server
# or
ping your-app.vercel.app  # For Vercel
```

### "Authentication failed"

**Check API key matches:**
- Server's `API_KEY` in `.env` or Vercel environment variables
- Client's `api_key` in `.smartcommitrc.json`

**Generate new API key:**
```bash
node -e "console.log(require('crypto').randomBytes(32).toString('hex'))"
```

### "No staged changes found"

```bash
# Stage your changes first
git add .

# Or stage specific files
git add file1.js file2.js
```

### Command not found: smartcommit

**Add Go bin to PATH:**
```bash
# Add to ~/.bashrc or ~/.zshrc
export PATH=$PATH:$(go env GOPATH)/bin

# Reload shell
source ~/.bashrc
```

## Security Notice

- Git diffs are sent to the AI server for analysis
- Sensitive data is automatically filtered before transmission
- Review generated messages before committing (use `--dry-run`)
- Use `--dry-run` flag to preview without committing
- Avoid using on repositories with highly sensitive/proprietary code
- For production use, enable server authentication with `API_KEY`
- Consider self-hosted server for maximum security

### Data Filtering

SmartCommit automatically filters sensitive data from git diffs:
- API keys and tokens
- Passwords and secrets
- Private keys
- Email addresses
- Credit card numbers
- GitHub tokens
- AWS access keys

## Benefits

- **Centralized API key management** - Set Gemini API key once on the server
- **No secrets in projects** - No need for `.env` files in every repository
- **Consistent commit messages** - Same AI logic across all projects
- **Team collaboration** - Share server URL with team members
- **Flexible deployment** - Use local server or cloud-hosted
- **Rate limiting** - Server-side rate limiting prevents API quota issues

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Related Projects

- [SmartCommit AI Server](https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit-AI-Server) - Backend Node.js server

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI framework
- Powered by [Google Gemini](https://ai.google.dev/) AI
- Follows [Conventional Commits](https://www.conventionalcommits.org/) specification

## Support

- **Issues**: [GitHub Issues](https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit/issues)
- **Discussions**: [GitHub Discussions](https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit/discussions)
- **Server Setup**: [AI Server Repository](https://github.com/JOSIAHTHEPROGRAMMER/Smart-Commit-AI-Server)
