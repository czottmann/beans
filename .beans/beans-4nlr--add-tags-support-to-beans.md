---
title: Add tags support to beans
status: done
type: feature
created_at: 2025-12-07T18:49:41Z
updated_at: 2025-12-07T18:59:00Z
---

Add support for tags on beans. Tags are:
- Stored as a string array in front matter (`tags: [frontend, urgent]`)
- Always lowercase, single-word, URL-formatted (e.g. `wont-fix`, `tech-debt`)
- Free-form (no predefined list required)

## Checklist

- [ ] Add `Tags` field to Bean struct
- [ ] Add tag validation (lowercase, single-word, URL-safe)
- [ ] Add `--tag` flag to `beans update` to add tags
- [ ] Add `--untag` flag to `beans update` to remove tags
- [ ] Add `--tag` flag to `beans create`
- [ ] Add `--tag` filter to `beans list`
- [ ] Add `--no-tag` exclusion filter to `beans list`
- [ ] Show tags in `beans list` output (text and JSON)
- [ ] Show tags in `beans show` output
- [ ] Add tags display to TUI
- [ ] Write tests for tag functionality