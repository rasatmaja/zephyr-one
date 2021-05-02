---
name: Technical Task
about: Create a task to represents some value to a developer or to a system
title: Golang Project Setup
labels: 'Type: Task'
assignees: rasatmaja

---

**Technical Description**
Initialize GO modules to defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build. Each dependency requirement is written as a module path and a specific semantic version.

**Acceptance Criteria**
- `go.mod` and `go.sum` files should present on root directory
- env example should present on root directory and contain prefix ZO on every variable


**Technical type** :  

- [ ] 🐞 Possible Bug
- [x] 🦌 New Feature
- [ ] 🤴 Code style update (formatting, renaming)
- [ ] 🏇🏼 Refactoring (no functional changes, no api changes)
- [ ] 🏅 Build related changes
- [ ] 🦧 Documentation content changes

**Task** :
- [ ] Init GO Modules
- [ ] Create folder structure
- [ ] Init ENV
