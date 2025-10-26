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

### Phase 1: Foundation (Tasks 0-4)
- [x] [Task 0: Project Bootstrap](task-00-project-bootstrap.md) - Initialize Go project
- [x] [Task 1: Tokenizer](task-01-tokenizer.md) - Convert text to tokens
- [x] [Task 2: Punctuation Normalization](task-02-punctuation-normalization.md) - Fix punctuation spacing
- [x] [Task 3: Quote Spacing](task-03-quote-spacing.md) - Clean quote boundaries
- [x] [Task 4: Renderer](task-04-renderer.md) - Convert tokens back to text

### Phase 2: Tag Processing (Tasks 5-9)
- [x] [Task 5: Tag Parser](task-05-tag-parser.md) - Parse command tags
- [x] [Task 6: Number Conversion](task-06-number-conversion.md) - Hex/bin to decimal
- [x] [Task 7: Single-Word Case Tags](task-07-case-single.md) - up/low/cap for one word
- [x] [Task 8: Multi-Word Case Tags](task-08-case-multi.md) - Case changes for multiple words
- [x] [Task 9: Tag Orchestration](task-09-tag-orchestration.md) - Coordinate tag applications

### Phase 3: Advanced Rules (Tasks 10-11)
- [x] [Task 10: A/An Rule](task-10-a-an-rule.md) - Grammatical article correction
- [x] [Task 11: Rules Inside Quotes](task-11-rules-in-quotes.md) - Apply rules within quotes

### Phase 4: Architecture (Tasks 12-13)
- [x] [Task 12: FSM Core](task-12-fsm-core.md) - Finite state machine controller
- [x] [Task 13: Chunked Pipeline](task-13-chunked-pipeline.md) - Large file handling

### Phase 5: Production (Tasks 14-16)
- [x] [Task 14: File I/O and CLI](task-14-file-io-cli.md) - Command-line interface
- [x] [Task 15: Error Handling](task-15-error-handling.md) - Robust error recovery
- [x] [Task 16: Logging](task-16-logging.md) - Diagnostics and monitoring

### Phase 6: Quality (Tasks 17-19)
- [x] [Task 17: Property Tests](task-17-property-tests.md) - Edge case testing
- [x] [Task 18: CI Integration](task-18-ci-integration.md) - Continuous testing
- [x] [Task 19: Documentation](task-19-documentation.md) - Final polish

### Optional
- [x] [Task A: AI Test Generation](task-A-ai-test-generation.md) - Automated test cases

### Task Templates
- [Task Template Folder](task-template/) - Templates for creating custom tasks

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

### Project Structure Convention:
```
go-reloaded/
├── formatter/          # Core formatting logic
├── tokenizer/          # Text tokenization
├── fsm/               # State machine (Task 12+)
├── cmd/               # CLI application (Task 14+)
└── testdata/          # Test fixtures (Task 18+)
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
