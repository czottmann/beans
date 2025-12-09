---
title: Make GraphQL schema fields non-nullable
status: completed
type: task
created_at: 2025-12-09T07:19:11Z
updated_at: 2025-12-09T07:20:23Z
---

Make Bean fields stricter with non-nullable types and appropriate defaults:

- Bean.type: String! (fallback to "task")
- Bean.priority: String! (fallback to "normal")
- Bean.createdAt: Time! (fallback to updated_at or file mod time)
- Bean.updatedAt: Time! (fallback to created_at)
- Bean.body: String! (empty string if no body)
- Bean.tags: [String!]! (empty array if none)
- Bean.links: [Link!]! (empty array if none)

## Checklist
- [x] Update schema.graphqls with non-nullable types
- [x] Add defaults in beancore/core.go loadBean()
- [x] Run mise codegen
- [x] Update tests if needed (no changes required)
- [x] Run tests to verify