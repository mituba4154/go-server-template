# Go Server Template - Agent Context

Read this file first.
Implement changes in this repository using the context below.

## Availability Notice

GoAPI and GoLoader are not public yet.
Do not assume public installation steps are available for external users.

## Project Summary

This repository is a template for writing Minecraft Paper/Spigot plugin logic in Go.
A Java bridge layer (GoAPI + GoLoader) handles communication between the Minecraft server and this Go process.
Users mainly implement gameplay logic in Go modules.

## Role of This Repository

`go-server-template` is the user-facing starter template.

- Keep `main.go` stable.
- Add only module registration lines in `main.go`.
- Put feature logic under `modules/<feature>/`.
- Use the `goapi` package API (`On`, `Command`, `RunLater`, `RunTimer`, etc.).

## Directory Layout

```text
.
├── main.go
├── modules/
│   └── example/
│       ├── register.go
│       └── join.go
└── .goapi/
    ├── AGENT.md
    └── SKILLS.md
```

## Implementation Rules

- Prefer small, focused modules by feature.
- Register event handlers and commands in each module's `register.go`.
- Keep handler logic in separate files for readability.
- Avoid changing infrastructure code unless required by the task.

## Typical Workflow

1. Create `modules/<feature>/`.
2. Write `Register()` in `register.go`.
3. Implement handlers/commands in `*.go`.
4. Add `<feature>.Register()` in `main.go`.
5. Build with `go build -o server .`.
