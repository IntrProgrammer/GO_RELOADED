# GO_RELOADED

A text formatter that applies tag-based transformations and grammar corrections.

## The Problem

Build a program that reads text, identifies special patterns and commands, then applies formatting rules automatically. Think of it as an automated proofreader with specific transformation capabilities.

### What It Does

- **Case Transformations**: `(up)`, `(low)`, `(cap)` - change text casing
- **Number Conversions**: `(hex)`, `(bin)` - convert to decimal
- **Grammar Fixes**: Automatic `a` → `an` correction
- **Punctuation**: Auto-spacing for `.`, `,`, `!`, `?`, `;`
- **Quote Processing**: Apply rules inside `'single quotes'`

### Examples

```
Input:  hello world (up)
Output: hello WORLD

Input:  FF (hex) files
Output: 255 files

Input:  a apple
Output: an apple
```

📖 **[Read Full Problem Description](docs/Analysis/Understunding%20_the_Problem.md)**

## Architecture

This project uses a **hybrid FSM-Pipeline architecture** combining:

- **FSM (Finite State Machine)**: For predictable state-based control flow
- **Pipeline**: For concurrent processing of large files

### Why This Approach?

- **Speed**: FSM handles small files efficiently
- **Scalability**: Pipeline processes large files in chunks
- **Flexibility**: State transitions handle complex transformations
- **Testability**: Each component is independently testable

```
Text → Tokenize → Normalize → Transform → Render → Output
         ↓           ↓            ↓          ↓
       [FSM States: READING → EVALUATING → EDITING]
```

📐 **[Read Architecture Details](docs/Analysis/Architecture_Type.md)**

## Project Structure

```
GO_RELOADED/
├── docs/
│   ├── Analysis/          # Problem and architecture docs
│   ├── tasks/             # TDD task breakdown (20 tasks)
│   └── PLANING_TDD/       # Development planning
├── COPYING.md             # GPL-3.0 License
└── Readme.md              # This file
```

## Development Approach

This project follows **Test-Driven Development (TDD)** with 20 incremental tasks:

### Task Phases

1. **Foundation** (0-4): Tokenizer, normalization, rendering
2. **Tag Processing** (5-9): Parse and apply transformations
3. **Advanced Rules** (10-11): Grammar and quote handling
4. **Architecture** (12-13): FSM and pipeline implementation
5. **Production** (14-16): CLI, error handling, logging
6. **Quality** (17-19): Testing, CI/CD, documentation

📋 **[View All Tasks](docs/tasks/README.md)**

### Task Template

Need to add custom tasks? Use the provided template:

🛠️ **[Task Template Guide](docs/tasks/task-template/README.md)**

## Quick Links

| Resource | Description |
|----------|-------------|
| [Problem Analysis](docs/Analysis/Understunding%20_the_Problem.md) | Detailed problem breakdown |
| [Architecture Guide](docs/Analysis/Architecture_Type.md) | FSM vs Pipeline comparison |
| [Task List](docs/tasks/README.md) | All 20 TDD tasks |
| [Task Templates](docs/tasks/task-template/) | Create custom tasks |

## License

This project is licensed under the **GNU General Public License v3.0**.

```
Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
```

📄 **[Full License Text](COPYING.md)**

---

**Built with TDD principles | Hybrid FSM-Pipeline Architecture | GPL-3.0 Licensed**
