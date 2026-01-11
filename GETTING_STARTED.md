# Getting Started with jira-beads-sync

Welcome! This guide will help you understand what jira-beads-sync is, why you might need it, and how to get started.

## What is jira-beads-sync?

jira-beads-sync is a tool that bridges two issue tracking systems:

- **Jira**: A popular enterprise issue tracking and project management platform
- **beads**: A lightweight, git-backed issue tracker that stores issues as YAML files in your repository

This tool allows you to work with Jira issues locally using beads, then sync your changes back to Jira. It's particularly useful for developers who want to:

- Work on Jira issues offline or in a more lightweight environment
- Use git-based workflows for issue tracking
- Keep issue tracking close to the code in version control
- Leverage AI assistants like Claude Code for natural language issue management

## When Should You Use This Tool?

Consider using jira-beads-sync if you:

- **Want local issue tracking**: Work with Jira issues offline as YAML files in your repository
- **Prefer git-based workflows**: Track issues in git alongside your code
- **Use Claude Code**: Manage Jira issues through natural language with Claude
- **Need bidirectional sync**: Import from Jira, work locally, push changes back
- **Work with complex hierarchies**: Sync epics, stories, subtasks, and their dependencies

## What Gets Synchronized?

jira-beads-sync handles:

### Issue Data
- Title, description, status, priority
- Assignee and reporter
- Created and updated timestamps
- Issue type (Epic, Story, Task, Bug, Subtask)

### Relationships
- Epic-to-issue associations (Jira epics → beads epics)
- Parent-child relationships (Story → Subtask)
- Issue links (blocks, is blocked by, depends on, relates to)

### Metadata
- Original Jira key (e.g., PROJ-123)
- Original Jira ID for syncing back

## Prerequisites

Before you start, you'll need:

1. **Go 1.21 or later** (if installing from source)
2. **Jira access**:
   - A Jira account with access to the projects you want to sync
   - An API token (we'll show you how to get one)
3. **beads** (optional but recommended):
   - Install from: https://github.com/conallob/beads
   - Used to view and manage synced issues locally

## Installation

Choose the method that works best for you:

### Option 1: Homebrew (macOS/Linux - Recommended)

```bash
brew tap conallob/tap
brew install jira-beads-sync
```

### Option 2: Download Pre-built Binary

Visit the [releases page](https://github.com/conallob/jira-beads-sync/releases/latest) and download the binary for your platform:

**macOS (Apple Silicon):**
```bash
curl -LO https://github.com/conallob/jira-beads-sync/releases/latest/download/jira-beads-sync_Darwin_arm64.tar.gz
tar xzf jira-beads-sync_Darwin_arm64.tar.gz
sudo mv jira-beads-sync /usr/local/bin/
```

**Linux (x86_64):**
```bash
curl -LO https://github.com/conallob/jira-beads-sync/releases/latest/download/jira-beads-sync_Linux_x86_64.tar.gz
tar xzf jira-beads-sync_Linux_x86_64.tar.gz
sudo mv jira-beads-sync /usr/local/bin/
```

### Option 3: Install from Source

```bash
go install github.com/conallob/jira-beads-sync/cmd/jira-beads-sync@latest
```

### Verify Installation

```bash
jira-beads-sync version
```

## First-Time Setup

### 1. Get a Jira API Token

You'll need an API token to authenticate with Jira:

1. Visit: https://id.atlassian.com/manage-profile/security/api-tokens
2. Click "Create API token"
3. Give it a name (e.g., "jira-beads-sync")
4. Copy the token (you won't be able to see it again!)

### 2. Configure jira-beads-sync

Run the interactive configuration:

```bash
jira-beads-sync configure
```

You'll be prompted for:
- **Jira Base URL**: Your Jira instance URL (e.g., `https://yourcompany.atlassian.net`)
- **Username**: Your Jira email address
- **API Token**: The token you created in step 1

The configuration is saved to `~/.config/jira-beads-sync/config.yml`.

### 3. Import Your First Issue

Let's try importing a Jira issue:

```bash
jira-beads-sync quickstart PROJ-123
```

Replace `PROJ-123` with an actual Jira issue key from your project.

This will:
1. Fetch the issue from Jira
2. Recursively fetch all related issues (subtasks, dependencies, etc.)
3. Create YAML files in `.beads/issues/` directory

### 4. View Imported Issues

If you have beads installed:

```bash
bd list                  # List all issues
bd show proj-123         # Show details of a specific issue
```

Or view the YAML files directly:

```bash
ls .beads/issues/
cat .beads/issues/proj-123.yaml
```

## Basic Workflow

Here's a typical workflow for using jira-beads-sync:

### Import from Jira

```bash
# Import a single issue and all its dependencies
jira-beads-sync quickstart PROJ-123

# Or use the full URL
jira-beads-sync quickstart https://jira.example.com/browse/PROJ-123
```

### Work Locally with beads

```bash
# List issues
bd list

# Show issue details
bd show proj-123

# Update issue status
bd update proj-123 --status in_progress

# Close an issue
bd close proj-123
```

### Sync Changes Back to Jira (Coming Soon)

```bash
# Sync all modified issues back to Jira
jira-beads-sync sync

# Sync specific issues
jira-beads-sync sync PROJ-123 PROJ-456
```

## Two Ways to Use jira-beads-sync

### 1. Command Line Interface (CLI)

Use the `jira-beads-sync` command directly in your terminal. Great for:
- Scripting and automation
- Direct control over sync operations
- One-time imports or conversions

See [CLI_GUIDE.md](docs/CLI_GUIDE.md) for comprehensive CLI documentation.

### 2. Claude Code Plugin

Use natural language with Claude to manage Jira issues. Great for:
- Natural workflows ("Import PROJ-123 from Jira")
- AI-assisted issue management
- Seamless integration with development tasks

See [PLUGIN_GUIDE.md](docs/PLUGIN_GUIDE.md) for plugin documentation.

## Understanding the Data Flow

```
┌─────────────┐
│    Jira     │ ← Fetch issues via REST API
│  (Cloud)    │
└──────┬──────┘
       │
       ↓ JSON
┌─────────────────┐
│  jira-beads-sync│ ← Convert & map data
│   (Protocol     │
│    Buffers)     │
└──────┬──────────┘
       │
       ↓ YAML
┌─────────────┐
│   beads     │ ← Work locally with issues
│ (.beads/)   │
└──────┬──────┘
       │
       ↓ Sync back
┌─────────────┐
│    Jira     │ ← Push changes via API
│  (Cloud)    │
└─────────────┘
```

## Operation Modes

jira-beads-sync supports three operation modes:

### 1. Quickstart Mode (Recommended)

Fetch issues directly from Jira API with automatic dependency walking:

```bash
jira-beads-sync quickstart PROJ-123
```

**Features:**
- Direct API integration
- Automatic dependency resolution
- Bidirectional sync support

### 2. Sync Mode

Push local changes back to Jira:

```bash
jira-beads-sync sync
```

**Features:**
- Detects changes in beads issues
- Maps beads state to Jira fields
- Maintains bidirectional consistency

### 3. Convert Mode

One-way conversion of exported Jira JSON files:

```bash
jira-beads-sync convert jira-export.json
```

**Features:**
- No API credentials required
- Offline processing
- One-way only (cannot sync back)

## Next Steps

Now that you understand the basics:

1. **Try importing an issue**: Follow the "Import Your First Issue" section above
2. **Learn CLI commands**: Read [CLI_GUIDE.md](docs/CLI_GUIDE.md) for all CLI options
3. **Try the Claude plugin**: Read [PLUGIN_GUIDE.md](docs/PLUGIN_GUIDE.md) for AI-assisted workflows
4. **See real examples**: Check out [EXAMPLES.md](docs/EXAMPLES.md) for practical use cases

## Getting Help

- **Documentation**: Check the [docs/](docs/) directory for detailed guides
- **Issues**: Report bugs at https://github.com/conallob/jira-beads-sync/issues
- **Examples**: See [EXAMPLES.md](docs/EXAMPLES.md) for common scenarios

## Common Questions

### Do I need beads installed?

Not strictly required, but recommended. Without beads, you'll only get YAML files. With beads, you get a full issue tracker with commands like `bd list`, `bd show`, etc.

### Can I sync changes back to Jira?

Yes! The sync mode allows bidirectional synchronization. Import from Jira, work locally, then push changes back.

### What about Jira custom fields?

Currently, jira-beads-sync focuses on standard Jira fields. Custom field support is planned for future releases.

### Is my data safe?

Yes! The tool only reads and updates issues you have access to. It uses secure API tokens (not passwords) and stores credentials locally in your config file.

### Can I use this with Jira Server/Data Center?

Yes! The tool works with both Jira Cloud and on-premise installations. Just configure your Jira base URL accordingly.

---

Ready to start? Head back to [Installation](#installation) or jump to [CLI_GUIDE.md](docs/CLI_GUIDE.md) for detailed usage!
