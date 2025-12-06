# Beans

Beans is an issue tracker for you, your team, and your coding agents. Instead of tracking tasks in a separate application, Beans stores them right alongside your code. You can use the `beans` CLI to interact with your tasks, but more importantly, so can your favorite coding agent!

This gives your robot friends a juicy upgrade: now they get a complete view of your project, make suggestions for what to work on next, track their progress, create bug issues for problems they find, and more.

You've been programming all your life; now you get to be a product manager. Let's go! 🚀

## Installation

### Homebrew

```bash
brew install hmans/beans/beans
```

## Usage

```bash
beans init          # Initialize a .beans/ directory
beans list          # List all beans
beans show <id>     # Show a bean's contents
beans create "Title" # Create a new bean
beans status <id> <status>  # Change status (open, in-progress, done)
beans archive       # Delete all done beans
```

All commands support `--json` for machine-readable output.

## Using with Claude Code

Beans integrates with [Claude Code](https://claude.ai/code) via hooks. Add this to your `.claude/settings.json`:

```json
{
  // ... other settings ...
  "hooks": {
    "SessionStart": [
      {
        "matcher": "",
        "hooks": [{ "type": "command", "command": "beans prompt" }]
      }
    ],
    "PreCompact": [
      {
        "matcher": "",
        "hooks": [{ "type": "command", "command": "beans prompt" }]
      }
    ]
  }
}
```

This runs `beans prompt` at session start and before context compaction, injecting instructions that teach Claude to use Beans for task tracking instead of its built-in TodoWrite tool.

## Using with other Agents

You can use Beans with other coding agents by configuring them to run `beans prompt` to get the prompt instructions for task management. We'll add specific integrations for popular agents over time.

## Contributing

This project currently does not accept contributions -- it's just way too early for that!
But if you do have suggestions or feedback, please feel free to open an issue.
