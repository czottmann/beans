---
title: Move .beans/config.yaml to .beans.yml at project root
status: completed
type: task
tags:
    - config
    - breaking-change
created_at: 2025-12-06T23:50:54Z
updated_at: 2025-12-08T14:28:31Z
links:
    - parent: beans-7lmv
---

Move configuration from `.beans/config.yaml` to `.beans.yml` at the project root. This follows the common pattern of root-level dotfiles for configuration (`.eslintrc`, `.prettierrc`, `.goreleaser.yml`).

## Changes

The new `.beans.yml` should support:
- `beans_path` - path to the beans directory (defaults to `.beans/`)
- All existing config options (`types`, `statuses`, etc.)

The CLI should:
- Walk upward from CWD to find `.beans.yml` (like git finds `.git/`)
- Support `--config` flag to specify an explicit config path
- Fall back to sensible defaults if no config file is found

## Checklist

- [x] Add `beans_path` field to config struct with default `.beans/`
- [x] Implement upward directory walk to find `.beans.yml`
- [x] Update config loading to use new location
- [x] Add `--config` flag to CLI for explicit config path
- [x] Update all commands to resolve beans path from config
- [x] Migrate this project's config from `.beans/config.yaml` to `.beans.yml`
- [x] Update documentation in CLAUDE.md and README
- [x] Add tests for config discovery and beans_path resolution
