# Claude Code Plugin Guide

Complete guide for using jira-beads-sync as a Claude Code plugin for natural language issue management.

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Available Commands](#available-commands)
- [Natural Language Usage](#natural-language-usage)
- [Slash Commands](#slash-commands)
- [Features](#features)
- [Workflows](#workflows)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)

## Overview

The jira-beads-sync Claude Code plugin enables you to manage Jira issues through natural language conversation with Claude. Instead of memorizing CLI commands, simply ask Claude to import, update, or sync issues.

**Benefits:**
- Natural language interface: "Import PROJ-123 from Jira"
- Context-aware suggestions: Claude understands your project
- Automatic dependency walking: Fetches related issues automatically
- Bidirectional sync: Import from Jira, work locally, push back
- Seamless workflow: Integrate issue management with coding tasks

## Installation

### Prerequisites

1. **jira-beads-sync CLI tool** installed and in PATH:
   ```bash
   # Using Homebrew
   brew install conallob/tap/jira-beads-sync

   # Or from source
   go install github.com/conallob/jira-beads-sync/cmd/jira-beads-sync@latest
   ```

2. **Claude Code CLI** installed (visit https://claude.ai/code)

3. **beads** (optional but recommended):
   ```bash
   # Install from https://github.com/conallob/beads
   ```

### Plugin Installation

#### Option 1: From Local Directory

```bash
# Clone the repository
git clone https://github.com/conallob/jira-beads-sync.git

# Start Claude Code with plugin enabled
claude --plugin-dir /path/to/jira-beads-sync
```

#### Option 2: Permanent Installation

Add to your Claude Code configuration:

```bash
# In your project's .claude/settings.json
{
  "pluginDirs": ["/path/to/jira-beads-sync"]
}
```

Then start Claude normally:
```bash
claude
```

### Verify Installation

Start Claude and ask:
```
You: Show me the jira-beads-sync plugin commands
```

Claude should list the available slash commands.

## Getting Started

### First-Time Setup

1. **Start Claude with the plugin:**
   ```bash
   cd /path/to/your/project
   claude --plugin-dir /path/to/jira-beads-sync
   ```

2. **Configure Jira credentials:**
   ```
   You: Configure my Jira credentials

   Claude: I'll help you configure Jira credentials.
   [Runs: jira-beads-sync configure]

   Enter Jira Base URL: https://yourcompany.atlassian.net
   Enter Jira Username: user@example.com
   Enter Jira API Token: **********************
   Configuration saved!
   ```

3. **Get an API token:**
   - Visit: https://id.atlassian.com/manage-profile/security/api-tokens
   - Click "Create API token"
   - Copy the token when prompted

4. **Import your first issue:**
   ```
   You: Import PROJ-123 from Jira

   Claude: I'll import PROJ-123 and its dependencies from Jira.
   [Runs: jira-beads-sync quickstart PROJ-123]

   ✓ Fetched 5 issue(s)
   ✓ Conversion complete!

   Issues imported:
   - proj-123: "Implement authentication system"
   - proj-124: "Add OAuth2 support" (subtask)
   - proj-125: "Update user model" (subtask)
   ```

## Available Commands

The plugin provides four slash commands:

### `/import-jira <key-or-url>`

Import a Jira issue and its entire dependency tree.

**Usage:**
```
/import-jira PROJ-123
/import-jira https://jira.example.com/browse/PROJ-123
```

**What it does:**
- Fetches the issue from Jira REST API
- Recursively walks all dependencies
- Creates beads issues in `.beads/issues/`

### `/configure-jira`

Configure Jira API credentials interactively.

**Usage:**
```
/configure-jira
```

**What it does:**
- Prompts for base URL, username, and API token
- Saves to `~/.config/jira-beads-sync/config.yml`

### `/sync-jira [keys...]`

Sync beads changes back to Jira (bidirectional sync).

**Usage:**
```
/sync-jira
/sync-jira PROJ-123 PROJ-456
```

**What it does:**
- Detects changes in beads issues
- Updates corresponding Jira issues via API
- Maintains consistency between systems

### `/convert-jira-export <file>`

Convert a Jira export JSON file to beads format (one-way).

**Usage:**
```
/convert-jira-export jira-export.json
```

**What it does:**
- Parses the JSON export file
- Converts to beads format
- Creates YAML files (cannot sync back)

## Natural Language Usage

The power of the plugin is in natural language interaction. Claude understands various ways of expressing the same intent.

### Importing Issues

All of these work:

```
You: Import PROJ-123 from Jira

You: Fetch the Jira issue TEAM-456 and all its dependencies

You: Pull PROJ-789 from Jira into beads

You: Get me the Jira issue at https://jira.example.com/browse/PROJ-100

You: I need to work on PROJ-123, can you import it?
```

Claude will execute the appropriate command and show you the results.

### Configuring Credentials

```
You: Configure Jira credentials

You: Set up my Jira API token

You: Help me connect to Jira

You: I need to configure jira-beads-sync
```

### Syncing Changes

```
You: Sync beads changes back to Jira

You: Push my beads updates to Jira

You: Update Jira with beads status changes

You: Sync PROJ-123 back to Jira
```

### Converting Exports

```
You: Convert jira-export.json to beads format

You: Import issues from my Jira export file

You: Process the Jira export at ./exports/sprint-42.json
```

## Slash Commands

Slash commands provide explicit control when needed:

### Basic Usage

```
/import-jira PROJ-123
/configure-jira
/sync-jira
/convert-jira-export export.json
```

### With Context

Claude maintains conversation context, so you can combine commands:

```
You: /import-jira PROJ-123

Claude: [Imports issue]

You: Show me what subtasks it has

Claude: PROJ-123 has 3 subtasks:
- PROJ-124: Add OAuth2 support (open)
- PROJ-125: Update user model (in_progress)
- PROJ-126: Write tests (open)

You: What should I work on first?

Claude: Based on dependencies, start with PROJ-125 since...
```

## Features

### Automatic Dependency Walking

The plugin automatically fetches all related issues:

```
You: Import the authentication epic AUTH-100

Claude: Importing AUTH-100 and walking the dependency tree...

Fetched:
- AUTH-100: Authentication System (epic)
  └─ AUTH-101: OAuth2 Integration (story)
     ├─ AUTH-102: Implement OAuth2 flow (subtask)
     ├─ AUTH-103: Add token storage (subtask)
     └─ AUTH-104: Test OAuth2 integration (subtask)
  └─ AUTH-105: User session management (story)
     └─ AUTH-106: Implement session store (subtask)
  └─ AUTH-110: Security audit (blocked by AUTH-101)

✓ Imported 8 issues total
```

### Status and Priority Mapping

The plugin automatically maps between Jira and beads:

**Status mapping:**
- Jira "To Do" ↔ beads `open`
- Jira "In Progress" ↔ beads `in_progress`
- Jira "Blocked" ↔ beads `blocked`
- Jira "Done" ↔ beads `closed`

**Priority mapping:**
- Jira "Blocker"/"Critical" ↔ beads `p0`
- Jira "Highest" ↔ beads `p1`
- Jira "High"/"Medium" ↔ beads `p2`
- Jira "Low" ↔ beads `p3`
- Jira "Lowest" ↔ beads `p4`

### Integration with beads

After import, Claude can help you work with beads:

```
You: Show me all open issues

Claude: Here are the open issues:
[Runs: bd list --status open]

You: What are the dependencies for proj-123?

Claude: Let me check the dependencies:
[Runs: bd show proj-123]

PROJ-123 depends on:
- PROJ-110: Design auth flow
- PROJ-112: Database schema

You: Mark proj-124 as in progress

Claude: [Runs: bd update proj-124 --status in_progress]
✓ Updated proj-124 status to in_progress
```

## Workflows

### Workflow 1: Starting Work on a Feature

```
You: I'm starting work on PROJ-123

Claude: Let me import PROJ-123 and its dependencies.
[Imports issue and subtasks]

PROJ-123: Implement authentication system
├─ PROJ-124: Add OAuth2 support
├─ PROJ-125: Update user model
└─ Depends on: PROJ-110 (Design complete)

All dependencies are resolved. You can start with PROJ-124.

You: Great, mark PROJ-124 as in progress

Claude: [Updates status]
✓ PROJ-124 is now in progress

You: Show me the acceptance criteria

Claude: PROJ-124 acceptance criteria:
1. OAuth2 flow implemented
2. Token storage configured
3. Tests passing
```

### Workflow 2: Planning Sprint Work

```
You: Import the sprint epic SPRINT-42

Claude: Importing SPRINT-42 and all stories...
[Fetches 15 issues]

✓ Imported 15 issues:
- 3 epics
- 8 stories
- 4 subtasks

You: What's the priority order?

Claude: Based on priorities and dependencies:

P0 (Critical):
- PROJ-120: Fix authentication bug

P1 (High):
- PROJ-121: User dashboard
- PROJ-122: API endpoints

P2 (Medium):
- PROJ-123: Documentation
- PROJ-124: Refactoring

You: Let's start with PROJ-120

Claude: PROJ-120: Fix authentication bug
Status: open
Assignee: unassigned
No blocking dependencies - ready to start!
```

### Workflow 3: Daily Standup Preparation

```
You: What Jira issues am I working on?

Claude: [Checks beads]
You're currently working on:
- PROJ-124: In progress (assigned to you)
- PROJ-125: In progress (assigned to you)

You: Show me the status of each

Claude:
PROJ-124: Add OAuth2 support (in_progress)
- Last updated: 2 hours ago
- Blocks: PROJ-126, PROJ-127

PROJ-125: Update user model (in_progress)
- Last updated: yesterday
- No blockers

You: I finished PROJ-124, close it and sync to Jira

Claude: [Closes issue and syncs]
✓ PROJ-124 marked as closed
✓ Synced to Jira
✓ Jira updated: PROJ-124 status changed to "Done"
```

### Workflow 4: Code Review Context

```
You: Show me what PROJ-123 is about

Claude: Let me fetch that issue.
[Imports if not present]

PROJ-123: Implement authentication system

Description:
Add OAuth2 authentication to support third-party logins...

Acceptance Criteria:
- OAuth2 flow implemented
- Tests cover all scenarios
- Documentation updated

Related Issues:
- Blocked by: PROJ-110 (complete)
- Blocks: PROJ-130, PROJ-131

You: Has this been started?

Claude: PROJ-123 status: in_progress
Assigned to: alice
Last updated: 2 days ago

You: Show me the subtasks

Claude: PROJ-123 has 3 subtasks:
- PROJ-124: Add OAuth2 support (closed)
- PROJ-125: Update user model (in_progress)
- PROJ-126: Write tests (open)
```

## Configuration

### Using CLAUDE.md Instructions

You can add Jira integration instructions to your project's `.claude/CLAUDE.md` to make Claude automatically aware of your workflow:

```markdown
# Jira Integration

This project uses Jira for issue tracking. Issues are synced with beads for local development.

## Workflow

When starting work on a Jira issue:
1. Ask Claude to import the issue: "Import PROJ-123 from Jira"
2. View the issue details: "Show me PROJ-123"
3. Work on the code
4. Update status: "Mark PROJ-123 as in progress"
5. When complete: "Close PROJ-123 and sync to Jira"

## Issue Mapping

- Jira epics → beads epics
- Jira stories/tasks → beads issues
- Jira subtasks → beads issues with dependencies
- Issue links → beads dependencies

## Common Commands

- Import issue: "Import PROJ-123 from Jira"
- Check status: "What's the status of PROJ-123?"
- Update status: "Mark PROJ-123 as in progress"
- Sync to Jira: "Sync beads changes to Jira"
- List open: "Show me all open issues"
```

See [EXAMPLES.md](EXAMPLES.md) for more CLAUDE.md examples.

### Environment Variables

For CI/CD or automation:

```bash
export JIRA_BASE_URL=https://acme.atlassian.net
export JIRA_USERNAME=user@example.com
export JIRA_API_TOKEN=your-token

claude --plugin-dir /path/to/jira-beads-sync
```

### Project-Specific Settings

In your project's `.claude/settings.json`:

```json
{
  "pluginDirs": ["/path/to/jira-beads-sync"],
  "environment": {
    "JIRA_BASE_URL": "https://acme.atlassian.net"
  }
}
```

## Troubleshooting

### Plugin Not Loaded

**Problem:** Claude doesn't recognize Jira commands

**Solutions:**
- Verify you started Claude with `--plugin-dir /path/to/jira-beads-sync`
- Check that the plugin directory contains `.claude-plugin/plugin.json`
- Restart Claude Code

### Command Not Found: jira-beads-sync

**Problem:** `jira-beads-sync: command not found`

**Solutions:**
- Install the CLI tool:
  ```bash
  brew install conallob/tap/jira-beads-sync
  # or
  go install github.com/conallob/jira-beads-sync/cmd/jira-beads-sync@latest
  ```
- Verify it's in PATH: `which jira-beads-sync`

### Authentication Failed

**Problem:** `401 Unauthorized` when importing

**Solutions:**
- Check API token: `jira-beads-sync configure`
- Verify token hasn't expired
- Create new token at https://id.atlassian.com/manage-profile/security/api-tokens

### Invalid Configuration

**Problem:** `No configuration found`

**Solutions:**
- Run: `/configure-jira`
- Or set environment variables
- Or create `~/.config/jira-beads-sync/config.yml`

### Natural Language Not Working

**Problem:** Claude doesn't understand Jira requests

**Solutions:**
- Be explicit: "Import PROJ-123 from Jira" instead of "Get PROJ-123"
- Use slash commands for precise control: `/import-jira PROJ-123`
- Check plugin is loaded: Ask "What plugins are available?"

## Advanced Usage

### Batch Operations

```
You: Import these issues: PROJ-123, PROJ-456, PROJ-789

Claude: I'll import all three issues...
[Runs multiple imports]
✓ Imported PROJ-123 (5 related issues)
✓ Imported PROJ-456 (3 related issues)
✓ Imported PROJ-789 (2 related issues)
Total: 10 unique issues
```

### Conditional Imports

```
You: Import PROJ-123 if it's not already imported

Claude: Checking if PROJ-123 exists locally...
[Checks .beads/issues/]
PROJ-123 already exists. Would you like to refresh it?

You: Yes, refresh it

Claude: [Re-imports to get latest data]
✓ Updated PROJ-123 and 4 related issues
```

### Cross-Project Work

```
You: Import the authentication epic from project AUTH and the API epic from project PLATFORM

Claude: Importing from multiple projects...
✓ AUTH-100 and 12 related issues
✓ PLATFORM-200 and 8 related issues

Cross-project dependency found: AUTH-105 depends on PLATFORM-201
Both issues imported successfully.
```

## Next Steps

- Try the [Example Workflows](EXAMPLES.md)
- Read the [CLI Guide](CLI_GUIDE.md) for advanced CLI usage
- See [Getting Started](../GETTING_STARTED.md) for basics

## Getting Help

- **Documentation**: Check [docs/](.) for all guides
- **Issues**: Report problems at https://github.com/conallob/jira-beads-sync/issues
- **Examples**: See real usage in [EXAMPLES.md](EXAMPLES.md)
