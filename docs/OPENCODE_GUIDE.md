# Opencode Integration Guide

Guide for using jira-beads-sync with [Opencode](https://opencode.ai), an open-source AI coding assistant.

## Overview

jira-beads-sync is a standard CLI binary that works with any AI coding tool — not just Claude Code. Opencode can run shell commands and invoke `jira-beads-sync` directly during your sessions.

**Benefits:**
- Use natural language to import and manage Jira issues
- Opencode runs `jira-beads-sync` commands on your behalf
- All standard CLI workflows are available
- No special plugin installation required

## Prerequisites

1. **jira-beads-sync CLI** installed and in PATH:
   ```bash
   # Using Homebrew
   brew install conallob/tap/jira-beads-sync

   # Or from source
   go install github.com/conallob/jira-beads-sync/cmd/jira-beads-sync@latest
   ```

2. **Opencode** installed (see [opencode.ai](https://opencode.ai))

3. **Jira credentials** configured:
   ```bash
   jira-beads-sync configure
   ```

## Project Instructions

Add a project instructions file so Opencode understands how to use jira-beads-sync in your project. Create a `.opencode/instructions.md` (or add to your existing project rules file):

```markdown
# Jira Integration

This project uses jira-beads-sync to sync Jira issues with beads.

## Importing Issues

To import a Jira issue and its dependency tree:
```bash
jira-beads-sync quickstart <JIRA-KEY>
# Example: jira-beads-sync quickstart PROJ-123
```

To import all issues with a specific label:
```bash
jira-beads-sync fetch-by-label <label>
# Example: jira-beads-sync fetch-by-label sprint-23
```

To import using a JQL query:
```bash
jira-beads-sync fetch-jql '<jql>'
# Example: jira-beads-sync fetch-jql 'project = PROJ AND assignee = currentUser()'
```

## Viewing Issues

```bash
bd list          # List all issues
bd show <id>     # Show issue details
```

## Syncing Back to Jira

```bash
jira-beads-sync sync
```
```

## Natural Language Usage

Once your project instructions are in place, ask Opencode naturally:

- "Import PROJ-123 from Jira and show me what was fetched"
- "Fetch all issues labeled sprint-23"
- "Run a JQL query for my open tasks in the MYPROJ project"
- "Show me all the imported issues with bd list"
- "Sync my beads changes back to Jira"
- "Configure my Jira credentials"

Opencode will run the appropriate `jira-beads-sync` commands on your behalf.

## Workflows

### Import and Start Working

```
You: Import PROJ-456 from Jira

Opencode: [runs: jira-beads-sync quickstart PROJ-456]
          Fetching PROJ-456...
          Fetching PROJ-457 (subtask)...
          ✓ Fetched 3 issue(s)
          Issues written to .beads/issues.jsonl

You: Show me what was imported

Opencode: [runs: bd list]
          ...lists the imported issues...
```

### Sprint Planning

```
You: Fetch all issues for sprint-47

Opencode: [runs: jira-beads-sync fetch-by-label sprint-47]
          Found 12 issue(s) with label sprint-47
          ✓ Fetched 18 issue(s) total (including dependencies)

You: Which ones are high priority?

Opencode: [runs: bd list --priority p0,p1]
          ...
```

### Daily Workflow

```
You: What are my open tasks?

Opencode: [runs: jira-beads-sync fetch-jql 'assignee = currentUser() AND status = Open']
          [runs: bd list]
          ...

You: Mark PROJ-123 as in progress, then sync to Jira

Opencode: [runs: bd update proj-123 --status in_progress]
          [runs: jira-beads-sync sync]
          ✓ Synced 1 change(s) to Jira
```

## Troubleshooting

### "command not found: jira-beads-sync"

Ensure the binary is in your PATH:
```bash
which jira-beads-sync
# If not found, add the install location to your PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```

### "no configuration found"

Run the configure command and follow the prompts:
```bash
jira-beads-sync configure
```

You'll need a Jira API token from https://id.atlassian.com/manage-profile/security/api-tokens

### Authentication issues

Test your credentials with:
```bash
jira-beads-sync whoami
```

## See Also

- [CLI Guide](CLI_GUIDE.md) — Complete command reference
- [Plugin Guide](PLUGIN_GUIDE.md) — Claude Code plugin documentation
- [Examples](EXAMPLES.md) — Real-world usage examples
