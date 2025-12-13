---
title: Fix missing outer connecting lines in tree display
status: completed
type: bug
priority: normal
created_at: 2025-12-13T10:43:36Z
updated_at: 2025-12-13T10:48:41Z
---

In deeply nested bean structures, the tree display is missing vertical continuation lines for outer nesting levels. For example, when displaying a feature under an epic under a milestone, there should be a vertical line showing the connection back to the milestone level, but it's currently missing.