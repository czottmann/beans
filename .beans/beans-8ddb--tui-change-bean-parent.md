---
title: 'TUI: Change bean parent'
status: completed
type: feature
priority: normal
created_at: 2025-12-12T22:23:04Z
updated_at: 2025-12-12T22:41:33Z
parent: beans-xnp8
---

Add the ability to change a bean's parent directly from the TUI.

## Requirements

- When editing a bean in the TUI, and in the bean list, provide a hotkey to change its parent
- Display a filterable list of eligible parent beans (respecting type hierarchy constraints)
- Support clearing the parent (setting to none)
- Show current parent if one exists
- Filter/search should work on both title and ID

## UI Considerations

- Could be triggered via a keybinding from the bean detail view and the list
- Use a popup/overlay pattern similar to other selection interfaces
- Show bean type and title in the selection list to help users choose
