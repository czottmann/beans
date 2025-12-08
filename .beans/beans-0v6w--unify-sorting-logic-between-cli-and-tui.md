---
title: Unify sorting logic between CLI and TUI
status: completed
type: task
created_at: 2025-12-08T14:52:57Z
updated_at: 2025-12-08T14:54:56Z
---

Extract the TUI's sortBeans function to a shared location so both CLI list and TUI use the same sorting logic. The TUI sorts by: 1) status order, 2) type order, 3) title. The CLI currently has different behavior.