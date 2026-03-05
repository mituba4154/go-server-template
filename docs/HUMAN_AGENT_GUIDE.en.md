# Human Guide (English)

> [!WARNING]
> GoAPI and GoLoader (goloder) are not public yet.

## Purpose

This guide is for people maintaining or extending this template.
Use it as the human-oriented counterpart to `.goapi/AGENT.md`.

## What This Repository Does

- Hosts a Go server process that receives bridged Minecraft events.
- Organizes gameplay logic into module packages under `modules/`.
- Exposes a stable entry point in `main.go`.

## Quick Workflow

1. Copy this template.
2. Create a new module folder under `modules/`.
3. Add a `Register()` function that binds events and commands.
4. Add one registration line in `main.go`.
5. Build and deploy the `server` binary.

```bash
go build -o server .
```

## Module Convention

```text
modules/
└── myfeature/
    ├── register.go
    ├── command.go
    └── events.go
```

- `register.go`: event/command registration only.
- Other files: implementation details.

## Main File Rule

- Keep `main.go` as infrastructure.
- Do not move runtime setup into feature modules.
- Only append module registration lines.

## Change Checklist

- New module compiles.
- `main.go` has the module registration line.
- No unrelated infra changes were introduced.
- Build output `server` is generated successfully.

## Team Notes

- Prefer small modules with clear ownership.
- Keep handlers deterministic and easy to test.
- Document non-obvious behavior near the code that implements it.
