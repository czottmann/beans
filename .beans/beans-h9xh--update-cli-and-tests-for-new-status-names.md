---
title: Update CLI and tests for new status names
status: completed
type: task
created_at: 2025-12-08T14:03:49Z
updated_at: 2025-12-08T14:52:52Z
---

The status configuration was updated to use: backlog, todo, in-progress, completed, scrapped. But the CLI and tests still reference old statuses (open, done).

## Checklist
- [x] Update create.go flag help to not hardcode status names
- [x] Update list_test.go to use new statuses
- [x] Update roadmap_test.go to use new statuses (already done)
- [x] Migrate existing beans from 'done' to 'completed'
- [x] Run tests to verify everything works
- [x] Update internal/bean/bean_test.go to use new statuses
- [x] Update internal/beancore/core_test.go to use new statuses
- [x] Update internal/ui/styles.go legacy functions for new statuses
- [x] Verify prompt.md and other docs are up to date