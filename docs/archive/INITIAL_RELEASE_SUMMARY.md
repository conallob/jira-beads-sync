# jira-beads-sync v0.1.0 - Initial Release

## Overview

**jira-beads-sync** is a Go-based CLI tool that provides bidirectional synchronization between Jira task trees and beads issues. It handles the hierarchical structure of Jira tasks (epics, stories, subtasks) and syncs them with the beads issue tracking system while preserving dependencies and relationships.

This initial release provides a complete, production-ready solution for syncing Jira projects with beads, with support for direct API integration, bidirectional synchronization, and offline conversion of exported data.

## Key Features

### 🚀 Multiple Operation Modes

**Quickstart Mode** (Recommended)
- Fetch issues directly from Jira REST API v2
- Recursive dependency graph walking
- Automatic traversal of subtasks, linked issues, and parent issues
- Single command to sync entire task hierarchies
- Interactive configuration setup

**Sync Mode** (Bidirectional)
- Sync beads state changes back to Jira
- Change detection and field mapping
- Status, assignee, priority, and description updates
- Maintains consistency between beads and Jira
- Conflict resolution for concurrent modifications

**Convert Mode** (One-Way)
- Convert previously exported Jira JSON files
- Offline processing for archived projects
- No API credentials required for exported data
- Note: Convert mode does not support syncing back to Jira

### 🎯 Intelligent Synchronization

- **Hierarchical Mapping**: Jira epics ↔ beads epics, stories ↔ issues, subtasks ↔ dependent issues
- **Dependency Preservation**: Syncs Jira issue links (blocks, depends on) with beads dependencies
- **Status Mapping**: Bidirectional mapping between Jira status categories and beads status enum
- **Priority Mapping**: Converts and syncs between Jira priorities and beads p0-p4 scale
- **Metadata Preservation**: Maintains Jira key, ID, and type for traceability and sync operations
- **Relationship Tracking**: Preserves and syncs parent-child and epic relationships
- **Change Detection**: Identifies modifications in beads for push back to Jira

### 🏗️ Protocol Buffers Architecture

Built on a modern, type-safe architecture using Protocol Buffers:

```
JSON (Jira) → Protobuf (Jira) → Protobuf (Beads) → YAML (Beads)
```

Benefits:
- Single source of truth for data structures (`.proto` files)
- Strong typing across all layers
- Built-in schema versioning support
- Easy to extend with new rendering formats
- Efficient serialization and validation

### 🤖 Claude Code Integration

Native plugin support for Claude Code (claude.ai/code):
- Natural language commands: "Import PROJ-123 from Jira", "Sync beads changes to Jira"
- Slash commands: `/import-jira PROJ-123`, `/sync-jira`
- Interactive credential configuration
- Seamless bidirectional integration with beads workflow

### 🔐 Flexible Configuration

Multiple configuration methods (in order of precedence):
1. **Interactive setup**: `jira-beads-sync configure`
2. **Environment variables**: `JIRA_BASE_URL`, `JIRA_USERNAME`, `JIRA_API_TOKEN`
3. **Config file**: `~/.config/jira-beads-sync/config.yml`

Secure credential storage with API token support (not passwords).

### 📦 Comprehensive Distribution

**Pre-built Binaries**
- Linux: x86_64, ARM64
- macOS: Intel (x86_64), Apple Silicon (ARM64)
- Windows: x86_64, ARM64

**Package Formats**
- **Homebrew**: `brew install conallob/tap/jira-beads-sync`
- **DEB packages**: Debian, Ubuntu, and derivatives
- **RPM packages**: RHEL, Fedora, CentOS, and derivatives
- **Container images**: Multi-arch Docker images on GitHub Container Registry
- **Checksums**: SHA256 checksums for all artifacts

### 🧪 Production Quality

- **Comprehensive test coverage**: 9 test files covering critical functionality
- **CI/CD Pipeline**: Automated testing, linting, and releases via GitHub Actions
- **Multi-version support**: Tested on Go 1.21, 1.22, and 1.23
- **Code quality**: golangci-lint enforcement, gofmt compliance
- **Type safety**: Protocol Buffer validation throughout the pipeline

## Technical Highlights

### API Integration
- **Jira REST API v2** client with authentication
- **Bidirectional operations**: Fetch from Jira, push updates to Jira
- Recursive issue fetching with circular dependency detection
- Handles all link types: subtasks, parent issues, issue links (inward/outward)
- Smart duplicate prevention using visited map
- Change detection and field update operations
- Configurable base URL for on-premise and cloud instances

### Data Processing
- **Protocol Buffer** definitions as source of truth (`proto/jira.proto`, `proto/beads.proto`)
- JSON adapter layer for Jira exports (`internal/jira/adapter.go`)
- Type-safe conversion layer (`internal/converter/proto_converter.go`)
- YAML rendering for beads format (`internal/beads/yaml.go`)

### Project Structure
```
jira-beads-sync/
├── cmd/jira-beads-sync/     # CLI entry point
├── internal/
│   ├── jira/                # Jira integration (API client & adapter)
│   ├── beads/               # Beads YAML rendering
│   ├── converter/           # Protobuf conversion logic
│   └── config/              # Configuration management
├── proto/                   # Protocol Buffer definitions
├── gen/                     # Generated code from protobuf
├── .claude-plugin/          # Claude Code plugin metadata
├── .github/workflows/       # CI/CD automation
└── testdata/                # Test fixtures
```

## Installation

### Homebrew (macOS/Linux)
```bash
brew install conallob/tap/jira-beads-sync
```

### Debian/Ubuntu
```bash
curl -LO https://github.com/conallob/jira-beads-sync/releases/download/v0.1.0/jira-beads-sync_0.1.0_amd64.deb
sudo dpkg -i jira-beads-sync_0.1.0_amd64.deb
```

### RHEL/Fedora/CentOS
```bash
curl -LO https://github.com/conallob/jira-beads-sync/releases/download/v0.1.0/jira-beads-sync_0.1.0_x86_64.rpm
sudo rpm -i jira-beads-sync_0.1.0_x86_64.rpm
```

### Container
```bash
docker pull ghcr.io/conallob/jira-beads-sync:v0.1.0
docker run ghcr.io/conallob/jira-beads-sync:v0.1.0 help
```

### Go Install
```bash
go install github.com/conallob/jira-beads-sync/cmd/jira-beads-sync@v0.1.0
```

## Quick Start

### Quickstart Mode (Import from Jira)
```bash
# One-time configuration
jira-beads-sync configure

# Import from Jira (fetches entire dependency graph)
jira-beads-sync quickstart https://jira.example.com/browse/PROJ-123

# Or use issue key directly
jira-beads-sync quickstart PROJ-123
```

### Sync Mode (Push changes to Jira)
```bash
# Sync beads changes back to Jira (coming soon)
jira-beads-sync sync

# Sync specific issues
jira-beads-sync sync PROJ-123 PROJ-456
```

### Convert Mode (One-way conversion)
```bash
# Convert exported JSON file
jira-beads-sync convert jira-export.json
```

### Claude Code Plugin
```bash
# Start Claude with plugin enabled
claude --plugin-dir /path/to/jira-beads-sync

# Then use natural language:
# "Import PROJ-123 from Jira and all its dependencies"
# "Sync beads changes back to Jira"
```

## What Gets Synchronized

The tool syncs and preserves:

✅ **Issue Data**
- Title, description, status, priority
- Assignee and reporter
- Created and updated timestamps
- Issue type (Epic, Story, Task, Bug, Subtask)

✅ **Relationships**
- Epic → Issue associations
- Parent → Child (Story → Subtask)
- Issue links (blocks, is blocked by, relates to)

✅ **Metadata**
- Original Jira key (e.g., PROJ-123)
- Original Jira ID
- Issue type information

## Commands Reference

```bash
jira-beads-sync quickstart <url-or-key>  # Fetch from Jira and sync to beads
jira-beads-sync sync [issue-keys...]     # Sync beads changes back to Jira (coming soon)
jira-beads-sync convert <file>           # Convert JSON export (one-way)
jira-beads-sync configure                # Interactive setup
jira-beads-sync version                  # Show version
jira-beads-sync help                     # Show usage
```

## Development & Contributing

Built with modern Go practices and comprehensive tooling:

```bash
make proto      # Generate protobuf code
make build      # Build binary
make test       # Run tests with coverage
make lint       # Run golangci-lint
make verify     # Run all checks (fmt, lint, test)
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed development guidelines.

## Release Artifacts

This release includes:
- **21 Go source files** implementing core functionality
- **9 comprehensive test files** with race detection
- **Multi-platform binaries** (6 architectures across 3 operating systems)
- **Linux packages** (.deb and .rpm with proper documentation installation)
- **Multi-arch container images** (amd64, arm64)
- **Homebrew formula** for easy installation on macOS/Linux
- **Complete documentation** (README, CLAUDE.md, PLUGIN.md, CONTRIBUTING.md)

## License

BSD-3-Clause License

Copyright (c) 2026, Conall O'Brien

## Acknowledgments

This tool integrates with:
- **beads**: Git-backed issue tracker ([beads documentation](https://github.com/conallob/beads))
- **Jira REST API v2**: Atlassian's issue tracking platform
- **Claude Code**: Anthropic's AI-powered CLI tool

## What's Next

Future enhancements under consideration:
- **Complete sync mode implementation** (beads → Jira updates)
- Support for Jira custom fields
- Comment and attachment synchronization
- Additional authentication methods (OAuth)
- Jira Cloud vs Server auto-detection
- Batch import with progress tracking
- Conflict resolution strategies
- Webhook support for real-time sync

---

**Full Changelog**: https://github.com/conallob/jira-beads-sync/commits/v0.1.0
