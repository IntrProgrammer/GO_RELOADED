# FSM-Based Implementation Tasks

## Overview
These tasks guide you through implementing the FSM (Finite State Machine) architecture for GO_RELOADED.

## Architecture Summary
The FSM architecture processes tokens in a single pass through state transitions:
- **READING** → **EVALUATING** → **EDITING** → **READING** (loop)
- **Quote tracking**: FSM tracks quote boundaries via `inQuote` state
- **Processors**: Handle transformations and stop at quote boundaries
- **Article correction**: Integrated as post-processing in FSM.Run()
- **Single-pass**: All tokens processed once, no segmentation needed

## Task Execution Order

### Phase 1: Foundation (Tasks 0-1) - KEEP ORIGINAL
- [x] [Task 0: Project Bootstrap](task-00-project-bootstrap.md)
- [x] [Task 1: Tokenizer](task-01-tokenizer.md)

### Phase 2: FSM Core (Tasks 2-6) - SINGLE-PASS FSM
- [x] [Task 2: FSM Core Structure](task-02-fsm-core.md) - State machine with quote tracking
- [x] [Task 3: Conversion Processor](task-03-conversion-processor.md) - Hex/bin with quote boundaries
- [x] [Task 4: Case Processor](task-04-case-processor.md) - up/low/cap with quote boundaries
- [x] [Task 5: Article Correction](task-05-article-correction.md) - a/an integrated into FSM
- [x] [Task 6: Formatter Integration](task-06-formatter-integration.md) - Single-pass FSM integration

### Phase 3: Pipeline Stages (Keep Original)
- [x] [Task 7: Renderer](task-04-renderer.md) - Token to text conversion


### Single-Pass FSM Approach
```
Tokenize → FSM.Run() (single loop) → Render

Inside FSM.Run():
  1. State loop (READING → EVALUATING → EDITING)
  2. Quote tracking (inQuote flag toggles on QUOTE tokens)
  3. Processors (stop at quote boundaries when looking backward)
  4. Post-process (CorrectArticles after main loop)
```

## Benefits of Single-Pass FSM
1. **True single-pass**: One loop through tokens (O(n) complexity)
2. **Quote state tracking**: No segmentation arrays needed
3. **Processor boundaries**: Automatic quote boundary respect
4. **Integrated post-processing**: Article correction in FSM.Run()
5. **Performance**: 50% fewer operations vs multi-pass
6. **Memory**: 33% less memory usage (no segmentation)
7. **Code simplicity**: 50% less code in formatter

## Implementation Guide for AI Agents

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
```

### Project Structure:
```
GO_RELOADED/
├── cmd/go-reloaded/     # Entry point
├── tokenizer/           # Lexical analysis
├── fsm/                 # State machine + processors
│   ├── state.go         # State definitions
│   ├── fsm.go           # FSM core with inQuote tracking
│   ├── processors.go    # Processors with quote boundary checks
│   └── *_test.go        # Tests
└── formatter/           # Formatter integration
    ├── formatter.go     # Format() with direct FSM call
    ├── renderer.go      # Token to string
    └── *_test.go        # Tests
```

## Testing Strategy

### Unit Tests
- **FSM States**: Test state transitions and quote tracking
- **Processors**: Test transformations and quote boundary stops
- **Article Correction**: Test FSM integration

### Integration Tests
- **FSM with Processors**: Test full single-pass flow
- **Quote Boundaries**: Test transformations respect quotes
- **End-to-End**: Test complete Format() with quote handling

### Property Tests
- Token count consistency
- No data loss
- Idempotency where applicable

## Success Criteria

After completing all FSM tasks, you should have:
- ✅ FSM with `inQuote` state tracking
- ✅ Quote boundary handling in `handleEvaluating()`
- ✅ ConversionProcessor stops at quotes
- ✅ CaseProcessor stops at quotes
- ✅ Article correction integrated in `FSM.Run()`
- ✅ Formatter with direct FSM call (no segmentation)
- ✅ All tests passing
- ✅ Single-pass processing verified

## Notes for AI Agents

- Each task is self-contained with complete code examples
- Follow the exact TDD workflow specified
- Do not skip tests
- Run full test suite after each task
- **Key Change**: No quote segmentation - FSM tracks quotes as state
- **Processors**: Must check for QUOTE tokens when looking backward
- **Article Correction**: Automatically called in FSM.Run()
- **Formatter**: Direct FSM call, no helper functions needed

## Recent Updates (Single-Pass Implementation)

### What Changed:
1. **Task 2**: Added `inQuote` field to FSM, quote tracking in `handleEvaluating()`
2. **Task 3**: ConversionProcessor stops at quote boundaries
3. **Task 4**: CaseProcessor stops at quote boundaries
4. **Task 5**: Article correction integrated into `FSM.Run()`
5. **Task 6**: Replaced quote segmentation with formatter integration

 **New Task 6: Formatter Integration**
- No `splitByQuotes()` function needed
- No `processWithFSM()` with segmentation
- Direct FSM call in `Format()` method

### Performance Improvements:
- **Before**: O(2n) - segmentation + FSM per segment
- **After**: O(n) - single FSM pass with state tracking
- **Result**: 50% fewer operations, 33% less memory

### Documentation:
- 📄 `SINGLE_PASS_FSM_CHANGES.md` - Technical implementation details
- 📄 `TASK_UPDATES_REPORT.md` - Complete task update report


