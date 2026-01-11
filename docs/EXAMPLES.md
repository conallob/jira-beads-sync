# Real-World Examples

Practical examples of using jira-beads-sync in different scenarios, including CLAUDE.md integration.

## Table of Contents

- [CLI Examples](#cli-examples)
- [Claude Code Plugin Examples](#claude-code-plugin-examples)
- [CLAUDE.md Integration](#claudemd-integration)
- [Workflow Examples](#workflow-examples)
- [Team Scenarios](#team-scenarios)

## CLI Examples

### Example 1: Import a Feature Branch

You're starting work on a new feature tracked in Jira.

```bash
# Setup (one-time)
jira-beads-sync configure

# Import the feature issue
jira-beads-sync quickstart FEAT-100

# Output:
# Fetching FEAT-100...
# Fetched issue: FEAT-100 - New user dashboard
# Walking dependencies...
#   Fetching subtask: FEAT-101 - Design wireframes
#   Fetching subtask: FEAT-102 - Implement layout
#   Fetching subtask: FEAT-103 - Add user widgets
#   Fetching linked issue: FEAT-90 - User API (blocks)
# ✓ Fetched 5 issue(s)
# ✓ Conversion complete!

# View imported issues
bd list
# proj-123    FEAT-100  New user dashboard        open     p1
# proj-124    FEAT-101  Design wireframes         open     p2
# proj-125    FEAT-102  Implement layout          open     p2
# proj-126    FEAT-103  Add user widgets          open     p2
# proj-127    FEAT-90   User API                  closed   p1

# Check what to work on first
bd show proj-123
# Dependencies:
#   - proj-127 (FEAT-90) [closed] ✓

# Start working
bd update proj-123 --status in_progress
git checkout -b feat/user-dashboard
```

### Example 2: Sprint Planning

Import all issues for an upcoming sprint.

```bash
# Import the sprint epic
jira-beads-sync quickstart SPRINT-23

# List by priority
bd list --sort priority

# Export to JSON for analysis
bd list --format json > sprint-23-issues.json

# Count issues by status
bd list --format json | jq 'group_by(.status) | map({status: .[0].status, count: length})'
```

### Example 3: Sync Status Updates Back

Update issue status locally and sync to Jira.

```bash
# Work on issues
bd update proj-123 --status in_progress
bd update proj-124 --status in_progress

# Complete some work
bd close proj-124

# Sync changes back to Jira
jira-beads-sync sync

# Output:
# Syncing changes to Jira...
# ✓ FEAT-100: Updated status to "In Progress"
# ✓ FEAT-101: Updated status to "Done"
# ✓ 2 issues synced successfully
```

### Example 4: Convert Legacy Export

You have a JSON export from an old Jira project.

```bash
# Convert the export file
jira-beads-sync convert jira-exports/legacy-project-2023.json

# Output:
# Reading jira-exports/legacy-project-2023.json...
# Found 156 issues
# Converting to beads format...
# ✓ Conversion complete!
# Created 156 issues in .beads/issues/

# Review converted issues
bd list --format table | head -20
```

## Claude Code Plugin Examples

### Example 1: Natural Language Import

```
You: I'm starting work on PROJ-456 from Jira. Can you import it for me?

Claude: I'll import PROJ-456 and its dependencies from Jira.

[Runs: jira-beads-sync quickstart PROJ-456]

✓ Fetched 4 issue(s)
✓ Conversion complete!

Imported issues:
- PROJ-456: "Add payment integration" (Story, p1, open)
  Dependencies:
  └─ PROJ-457: "Stripe API integration" (Subtask, p1, open)
  └─ PROJ-458: "Payment UI components" (Subtask, p2, open)
  └─ Blocks: PROJ-470 "Billing dashboard"

All dependencies are resolved. You can start with PROJ-457 (Stripe integration).

You: Great! Mark PROJ-457 as in progress

Claude: [Runs: bd update proj-457 --status in_progress]
✓ Updated PROJ-457 to in_progress

You: Show me the acceptance criteria for PROJ-457

Claude: [Runs: bd show proj-457]

PROJ-457: Stripe API integration

Description:
Integrate Stripe payment processing API for handling customer payments.

Acceptance Criteria:
- Stripe SDK integrated
- Payment endpoints implemented
- Error handling for failed payments
- Tests with 90% coverage
- Documentation updated

Assignee: unassigned
Priority: p1 (high)
Status: in_progress
```

### Example 2: Code Review Workflow

```
You: I'm reviewing a PR for PROJ-123. What's this issue about?

Claude: Let me fetch PROJ-123 from Jira.

[Runs: jira-beads-sync quickstart PROJ-123]

PROJ-123: Implement OAuth2 authentication

Type: Story
Priority: p0 (critical)
Status: in_progress
Assignee: alice

Description:
Add OAuth2 authentication flow to support Google and GitHub login...

Related Work:
- Part of epic: AUTH-100 "Authentication System"
- Blocks: PROJ-130 "User profile management"
- Depends on: PROJ-110 "Auth service setup" (✓ complete)

Subtasks:
- PROJ-124: OAuth2 provider setup (✓ complete)
- PROJ-125: Frontend login buttons (in progress)
- PROJ-126: Integration tests (open)

You: Has the OAuth2 provider setup been completed?

Claude: Yes, PROJ-124 (OAuth2 provider setup) is marked as complete.
According to the issue, PROJ-125 (Frontend login buttons) is currently
in progress, which is the logical next step.

You: What needs to be done for PROJ-125?

Claude: [Runs: bd show proj-125]

PROJ-125: Frontend login buttons

Acceptance Criteria:
- Login buttons for Google and GitHub
- Proper OAuth2 redirect handling
- Error state handling
- Loading indicators
- Accessibility compliance

Current status: in_progress (updated 3 hours ago)
```

### Example 3: Daily Standup Prep

```
You: What Jira issues am I currently working on?

Claude: Let me check your current work in beads.

[Runs: bd list --status in_progress --assignee $USER]

You're working on 3 issues:

1. PROJ-234: Refactor database queries (p2)
   - Started: 2 days ago
   - No blockers

2. PROJ-245: Add caching layer (p1)
   - Started: yesterday
   - Blocked by: PROJ-243 (still open)

3. PROJ-251: Update documentation (p3)
   - Started: 1 week ago
   - No blockers

You: What's blocking PROJ-245?

Claude: [Runs: bd show proj-245]

PROJ-245 is blocked by PROJ-243: "Configure Redis cluster"

Let me check the status of PROJ-243...
[Runs: bd show proj-243]

PROJ-243: Configure Redis cluster
Status: in_progress
Assignee: bob
Updated: 1 hour ago

It looks like Bob is actively working on it. Would you like me to import
the latest status from Jira to see if there are any updates?

You: Yes please

Claude: [Runs: jira-beads-sync quickstart PROJ-243]

✓ Updated PROJ-243 from Jira

Good news! PROJ-243 status changed to "Done" in Jira 30 minutes ago.
The blocker for PROJ-245 is now resolved.
```

### Example 4: Epic Planning

```
You: Import the Q1 authentication epic from Jira

Claude: I'll import the authentication epic. Do you have the issue key?

You: AUTH-100

Claude: Importing AUTH-100 and walking the entire dependency tree...

[Runs: jira-beads-sync quickstart AUTH-100]

✓ Fetched 18 issue(s)
✓ Conversion complete!

Imported Epic: AUTH-100 "Q1 Authentication Improvements"

Stories (6):
- AUTH-101: OAuth2 integration (8 subtasks)
- AUTH-102: Two-factor authentication (4 subtasks)
- AUTH-103: Session management (3 subtasks)
- AUTH-104: Security audit (2 subtasks)
- AUTH-105: API authentication (3 subtasks)
- AUTH-106: Documentation (2 subtasks)

Total: 1 epic, 6 stories, 22 subtasks = 29 issues

You: What's the priority order?

Claude: Based on priority and dependencies:

Critical (p0):
1. AUTH-104: Security audit
   - No dependencies, should start ASAP

High (p1):
2. AUTH-101: OAuth2 integration
   - No dependencies, can start immediately
3. AUTH-102: Two-factor authentication
   - Depends on: AUTH-101

Medium (p2):
4. AUTH-103: Session management
5. AUTH-105: API authentication
   - Depends on: AUTH-101, AUTH-103

Low (p3):
6. AUTH-106: Documentation
   - Depends on: all others (document after completion)

Recommended start: AUTH-104 (security audit) and AUTH-101 (OAuth2) in parallel.
```

## CLAUDE.md Integration

### Example 1: Basic Project Integration

Add to your project's `.claude/CLAUDE.md`:

```markdown
# Project: Acme Web Application

## Issue Tracking

This project uses Jira for issue tracking. Issues are synced with beads for local development.

### Workflow

When starting work on a feature:
1. Import the Jira issue: "Import PROJ-123 from Jira"
2. Review the issue: "Show me PROJ-123"
3. Create a feature branch: `git checkout -b feat/PROJ-123-description`
4. Mark as in progress: "Mark PROJ-123 as in progress"
5. Work on the feature
6. When done: "Close PROJ-123 and sync to Jira"

### Issue Commands

Quick reference for common operations:
- Import: "Import PROJ-123 from Jira"
- Status: "What's the status of PROJ-123?"
- Update: "Mark PROJ-123 as in progress"
- Complete: "Close PROJ-123"
- Sync: "Sync all changes to Jira"
- List: "Show me all open issues"

### Issue Mapping

- Jira epics → beads epics (track large features)
- Jira stories → beads issues (track user stories)
- Jira tasks → beads issues (track technical work)
- Jira subtasks → beads issues with dependencies
- Issue links → beads dependencies
```

**Usage:**

```
You: I'm working on the user dashboard

Claude: Let me help you get started. Which Jira issue are you working on?

You: PROJ-456

Claude: [Imports PROJ-456 based on CLAUDE.md instructions]

✓ Imported PROJ-456: "New user dashboard"

Following the project workflow, I've:
1. ✓ Imported the issue from Jira
2. Ready for you to create a feature branch:
   `git checkout -b feat/PROJ-456-user-dashboard`
3. Ready to mark as in progress

Would you like me to mark PROJ-456 as in progress?
```

### Example 2: Team-Specific Workflow

```markdown
# Project: Mobile Banking App

## Development Workflow

### Sprint Issues

At sprint start, import the sprint epic:
- Sprint 23: SPRINT-23
- Sprint 24: SPRINT-24

Command: "Import sprint epic SPRINT-23 from Jira"

### Feature Development

For each feature, follow this workflow:

1. **Import**: "Import [JIRA-KEY] from Jira"
2. **Branch**: Create branch matching Jira key: `feat/[JIRA-KEY]-description`
3. **Status**: Mark as in progress: "Mark [JIRA-KEY] as in progress"
4. **Code**: Implement the feature
5. **PR**: Create PR with "[JIRA-KEY]" in title
6. **Review**: Get code review
7. **Merge**: Merge to main
8. **Complete**: "Close [JIRA-KEY] and sync to Jira"

### Dependency Checks

Before starting work, always check dependencies:
- "Show me dependencies for [JIRA-KEY]"
- "Are all blockers resolved for [JIRA-KEY]?"

### Daily Standup

Use these commands to prepare:
- "What issues am I working on?"
- "Show me my in-progress issues"
- "What did I complete yesterday?"

### Critical Issues

P0/P1 issues should be imported immediately:
- "Import [JIRA-KEY] from Jira"
- Notify team in Slack
- Update standup board

## Examples

### Starting a New Feature

```
User: I'm starting work on the payment screen
Claude: Which Jira issue is this for?
User: MOBILE-234
Claude: [Imports MOBILE-234]
       [Shows dependencies and acceptance criteria]
       [Suggests creating branch: feat/MOBILE-234-payment-screen]
       Ready to mark as in progress?
```

### Completing Work

```
User: I finished MOBILE-234
Claude: Great! Let me:
       1. Mark as closed in beads
       2. Sync status to Jira
       3. Check if this unblocks any other issues

       [Performs actions]

       ✓ MOBILE-234 closed
       ✓ Synced to Jira
       ✓ This unblocks: MOBILE-240, MOBILE-241
```
```

**Usage with this CLAUDE.md:**

```
You: I'm starting work on the payment screen

Claude: I'll help you follow the team workflow. Which Jira issue is this for?

You: MOBILE-234

Claude: [Imports MOBILE-234 following project guidelines]

MOBILE-234: Payment screen redesign
Priority: p1
Status: open

Following the project workflow:

1. ✓ Imported from Jira (with 3 subtasks)
2. Create branch: `git checkout -b feat/MOBILE-234-payment-screen`
3. Ready to mark as in progress

Dependencies check:
- MOBILE-220: Payment API (✓ complete)
- MOBILE-225: Design system (✓ complete)

All dependencies resolved! You're good to start.

Shall I mark MOBILE-234 as in progress?

You: Yes

Claude: ✓ Marked MOBILE-234 as in progress
       ✓ Synced to Jira

According to the acceptance criteria, you need to:
- Implement new payment form UI
- Add credit card validation
- Add PayPal integration
- Write integration tests

Good luck! Let me know when you complete MOBILE-234.
```

### Example 3: CI/CD Integration

```markdown
# CI/CD: Jira Issue Tracking

## Automated Issue Import

### On PR Creation

When a PR is created with a Jira key in the title:
1. Import the issue automatically
2. Check that all dependencies are resolved
3. Update Jira with PR link

```bash
# .github/workflows/pr-create.yml
- name: Import Jira issue
  run: |
    ISSUE_KEY=$(echo "${{ github.event.pull_request.title }}" | grep -oE '[A-Z]+-[0-9]+')
    if [ -n "$ISSUE_KEY" ]; then
      jira-beads-sync quickstart "$ISSUE_KEY"
      bd show "$ISSUE_KEY"
    fi
```

### On PR Merge

When a PR is merged:
1. Close the issue in beads
2. Sync status to Jira
3. Check for unblocked issues

```bash
# .github/workflows/pr-merge.yml
- name: Close Jira issue
  run: |
    ISSUE_KEY=$(echo "${{ github.event.pull_request.title }}" | grep -oE '[A-Z]+-[0-9]+')
    if [ -n "$ISSUE_KEY" ]; then
      bd close "$ISSUE_KEY"
      jira-beads-sync sync "$ISSUE_KEY"
    fi
```

## For Claude

When working with this repository:
- Always check PR titles for Jira keys
- Import issues mentioned in commits
- Validate dependencies before suggesting implementation
- Update Jira when completing work
```

## Workflow Examples

### Example 1: Full Feature Development Cycle

Complete workflow from planning to deployment.

```bash
# 1. Sprint Planning
jira-beads-sync quickstart SPRINT-15
bd list --sort priority

# 2. Start Feature
jira-beads-sync quickstart FEAT-234
bd show feat-234
bd update feat-234 --status in_progress
git checkout -b feat/FEAT-234-user-settings

# 3. Development
# ... write code ...
git commit -m "feat(FEAT-234): implement user settings page"

# 4. Mark Subtask Complete
bd close feat-235  # Subtask: UI implementation
jira-beads-sync sync FEAT-235

# 5. Create PR
gh pr create --title "FEAT-234: User settings page" --body "..."

# 6. After Merge
bd close feat-234
jira-beads-sync sync FEAT-234

# 7. Check Unblocked Issues
bd list --format json | jq '.[] | select(.depends_on[] == "feat-234")'
```

### Example 2: Bug Triage

Handling a critical bug report.

```bash
# 1. Import the bug
jira-beads-sync quickstart BUG-789

# 2. Check severity
bd show bug-789
# Priority: p0 (critical)
# Affects: Authentication system

# 3. Check dependencies
bd show bug-789 --deps
# Needs: BUG-788 (root cause analysis) [open]

# 4. Import dependencies
jira-beads-sync quickstart BUG-788

# 5. Start investigation
bd update bug-788 --status in_progress
git checkout -b bugfix/BUG-788-auth-crash

# 6. After fix
bd close bug-788
bd update bug-789 --status in_progress  # Now unblocked
jira-beads-sync sync
```

## Team Scenarios

### Scenario 1: Onboarding a New Developer

```
Manager: Import the onboarding epic for our new developer

Claude: [Runs: jira-beads-sync quickstart ONBOARD-10]

Imported: ONBOARD-10 "Developer Onboarding - Alice"

Setup Tasks (7):
- ONBOARD-11: Grant repository access (complete)
- ONBOARD-12: Setup development environment (in progress)
- ONBOARD-13: Review architecture docs (open)
- ONBOARD-14: Pair program with senior dev (open)

Training Tasks (5):
- ONBOARD-15: Complete React training (open)
- ONBOARD-16: Complete testing workshop (open)
- ONBOARD-17: Code review training (open)

First Project (3):
- ONBOARD-20: Implement contact form (open, depends on training)
- ONBOARD-21: Write tests for contact form (open)
- ONBOARD-22: Deploy to staging (open)

Manager: Great. Alice, as you complete each task, tell Claude to
         "Close ONBOARD-XX" and it will update Jira automatically.
```

### Scenario 2: Sprint Retrospective

```
Team: Let's review what we completed in Sprint 23

Claude: [Runs: jira-beads-sync quickstart SPRINT-23]
        [Runs: bd list --format json]

Sprint 23 Results:

Completed (12 issues):
✓ FEAT-234: User settings page (with 4 subtasks)
✓ FEAT-240: Dark mode support
✓ BUG-156: Fix login redirect
✓ BUG-157: Fix profile image upload
... (8 more)

In Progress (3 issues):
⧗ FEAT-250: Notification system (85% complete, 1 subtask remaining)
⧗ FEAT-255: Email templates (blocked by design)
⧗ REFACTOR-50: Code cleanup (non-critical)

Not Started (2 issues):
○ FEAT-260: Admin dashboard (moved to Sprint 24)
○ DOCS-30: API documentation (moved to Sprint 24)

Velocity: 12 completed / 17 planned = 71%
```

### Scenario 3: Cross-Team Dependency

```
Team A: We need the API endpoints from Team B before we can start

Claude: What's the Jira key for that dependency?

Team A: PLATFORM-456

Claude: [Runs: jira-beads-sync quickstart PLATFORM-456]

PLATFORM-456: User API endpoints
Team: Platform Team
Status: in_progress
Assigned: Bob (Team B)
Target: End of week

This is blocking:
- MOBILE-234: Mobile app user screen (Team A)
- WEB-567: Web dashboard (Team A)

Team A: Can you notify us when PLATFORM-456 is complete?

Claude: I can help with that. I'll check the status and let you know
        when it's marked as done in Jira. You can also run:

        `jira-beads-sync quickstart PLATFORM-456`

        anytime to get the latest status.
```

## Next Steps

- Try these examples in your own projects
- Customize the workflows for your team
- Add your own CLAUDE.md instructions
- Share your examples with the community!

## Contributing Examples

Have a great example? Contribute to this doc:
1. Fork the repository
2. Add your example to this file
3. Submit a PR

See [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.
