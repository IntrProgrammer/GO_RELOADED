# GO_RELOADED Tasks - Agent Execution Guide

This directory contains individual task files structured for execution by coding agents. Each task follows TDD principles and includes complete implementation guidance.

## Task Structure

Each task file contains:
- **Objective**: Clear goal statement
- **Prerequisites**: Dependencies on previous tasks
- **Deliverables**: Checklist of outputs
- **Implementation Steps**: Detailed code examples
- **Verification Commands**: Testing and validation
- **Success Criteria**: Definition of done
- **TDD Workflow**: Red-Green-Refactor cycle
- **Git Commit Message**: Standardized commit format

## Task Execution Order

### Phase 1: Foundation (Tasks 0-1)
- [x] [Task 0: Project Bootstrap](task-00-project-bootstrap.md) - Initialize Go project
- [x] [Task 1: Tokenizer](task-01-tokenizer.md) - Convert text to tokens

### Phase 2: FSM Core (Tasks 2-6)
- [x] [Task 2: FSM Core Structure](task-02-fsm-core.md) - State machine with quote tracking
- [x] [Task 3: Conversion Processor](task-03-conversion-processor.md) - Hex/bin with quote boundaries
- [x] [Task 4: Case Processor](task-04-case-processor.md) - up/low/cap with quote boundaries
- [x] [Task 5: Article Correction](task-05-article-correction.md) - a/an integrated into FSM
- [x] [Task 6: Formatter Integration](task-06-formatter-integration.md) - Single-pass FSM integration

### Phase 3: Pipeline Stages (Task 7)
- [x] [Task 7: Renderer](task-12-renderer.md) - Token to text conversion

### Phase 4: Production (Task 8)
- [x] [Task 8: File I/O and CLI](task-14-file-io-cli.md) - Command-line interface


## Agent Execution Instructions

### For Each Task:
1. Read the task file completely
2. Verify prerequisites are met
3. Follow TDD workflow (Red → Green → Refactor)
4. Run verification commands
5. Confirm success criteria
6. Commit with provided message
7. Move to next task

### Verification After Each Task:
```bash
go test ./...
go vet ./...
go fmt ./...
go test -race ./...
```

### Project Structure:
```
GO_RELOADED/
├── cmd/go-reloaded/     # Entry point (main.go)
├── docs/                # Project documentation
│   ├── Analysis/        # Design docs, problem description
│   └── tasks/           # Task files for implementation
├── formatter/           # Pipeline orchestration
│   ├── formatter.go     # Main Format() + quote segmentation
│   ├── punctuation.go   # Punctuation normalization
│   ├── quote.go         # Quote spacing cleanup
│   └── renderer.go      # Token to string conversion
├── fsm/                 # Finite state machine
│   ├── state.go         # State definitions
│   ├── fsm.go           # FSM core logic
│   └── processors.go    # Transformation processors
└── tokenizer/           # Tokenization
    ├── token.go         # Token type definitions
    └── tokenizer.go     # Tokenize function
```

## Notes for Coding Agents

- Each task is self-contained with complete code examples
- Follow the exact TDD workflow specified
- Do not skip tests
- Maintain clean git history (one commit per task)
- Run full test suite after each task
- Ask for clarification if prerequisites are unclear

## Progress Tracking

Update the checkboxes above as tasks are completed. Each completed task should have:
- ✅ All tests passing
- ✅ No vet warnings
- ✅ Code formatted
- ✅ Git commit created
- ✅ Documentation updated
