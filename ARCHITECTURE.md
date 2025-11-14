# GO_RELOADED Architecture

## Architecture Pattern: Hybrid FSM-Pipeline

You've implemented a **Hybrid Architecture** that combines the strengths of both Finite State Machine (FSM) and Pipeline patterns, leaning more toward FSM for intelligent text transformations while maintaining pipeline simplicity for preprocessing.

---

## High-Level Flow

```
Input Text
    ↓
[PIPELINE: Tokenization]
    ↓
[PIPELINE: Punctuation Normalization]
    ↓
[PIPELINE: Quote Spacing Cleanup]
    ↓
[FSM: Quote Segment Processing]
    ├─→ Segment 1 → FSM → ConversionProcessor → CaseProcessor → Article Correction
    ├─→ Segment 2 → FSM → ConversionProcessor → CaseProcessor → Article Correction
    └─→ Segment N → FSM → ConversionProcessor → CaseProcessor → Article Correction
    ↓
[PIPELINE: Rendering]
    ↓
Output Text
```

---

## Architecture Layers

### 1. **Entry Point Layer** (`cmd/go-reloaded/main.go`)
**Pattern**: Simple CLI Handler

**Responsibilities**:
- Parse command-line arguments (input/output files)
- Read input file
- Delegate to Formatter
- Write output file
- Error handling

**Design**: Minimal logic, delegates all processing to formatter layer

---

### 2. **Tokenization Layer** (`tokenizer/`)
**Pattern**: Pipeline Stage (Lexical Analysis)

**Components**:
- `Token`: Data structure (Type + Value)
- `TokenType`: Enum (WORD, TAG, QUOTE, PUNCTUATION, WHITESPACE)
- `Tokenize()`: Regex-based tokenizer

**Strategy**:
- Single master regex with capture groups
- Pattern priority: TAG → QUOTE → PUNCTUATION → WHITESPACE → WORD
- Immutable token stream output

**Why Pipeline Here**: Tokenization is a pure transformation with no state dependencies

---

### 3. **Preprocessing Layer** (`formatter/punctuation.go`, `formatter/quote.go`)
**Pattern**: Pipeline Stages (Normalization)

**Stage 1: Punctuation Normalization**
```
Input:  [WORD:"Hello", WHITESPACE:" ", PUNCTUATION:",", WORD:"world"]
Output: [WORD:"Hello", PUNCTUATION:",", WHITESPACE:" ", WORD:"world"]
```
- Remove whitespace before punctuation
- Add whitespace after punctuation

**Stage 2: Quote Spacing Cleanup**
```
Input:  [QUOTE:"'", WHITESPACE:" ", WORD:"hello", WHITESPACE:" ", QUOTE:"'"]
Output: [QUOTE:"'", WORD:"hello", QUOTE:"'"]
```
- Remove whitespace after opening quote
- Remove whitespace before closing quote

**Why Pipeline Here**: Simple, stateless transformations that always execute in order

---

### 4. **Core Processing Layer** (`formatter/formatter.go` + `fsm/`)
**Pattern**: Hybrid (Pipeline orchestration + FSM execution)

#### 4.1 Quote Segmentation (Pipeline Logic)
```go
func splitByQuotes(tokens []Token) [][]Token
```
- Splits token stream by quote boundaries
- Each segment processed independently
- Prevents transformations from crossing quote boundaries

#### 4.2 FSM Processing (State Machine Logic)
```go
func processWithFSM(tokens []Token) []Token
```

**FSM States**:
1. **READING**: Check if more tokens exist
2. **EVALUATING**: Examine token type, route to processor or append
3. **EDITING**: Apply transformations via processors
4. **DONE**: Terminal state
5. **ERROR**: Error handling state

**FSM Flow**:
```
READING → EVALUATING → {TAG? → EDITING → READING}
                     → {WORD? → Append → READING}
                     → {OTHER? → Append → READING}
```

**Why FSM Here**: Complex state-dependent transformations with backward-looking logic

---

### 5. **Transformation Layer** (`fsm/processors.go`)
**Pattern**: Strategy Pattern (Pluggable Processors)

#### Processor Interface
```go
type Processor interface {
    Process(result []Token, currentToken Token) (modified []Token, handled bool)
}
```

#### Implemented Processors

**ConversionProcessor** (Priority 1)
- Handles: `(hex)`, `(bin)`
- Logic: Look backward for previous WORD, convert to decimal
- Modifies: Previous token value, removes tag and trailing whitespace

**CaseProcessor** (Priority 2)
- Handles: `(up)`, `(low)`, `(cap)` with optional count
- Logic: Look backward for N WORD tokens, apply case transformation
- Modifies: Previous N tokens, removes tag and trailing whitespace

**Article Correction** (Post-processor)
- Handles: `a` → `an` before vowels/h
- Logic: Look forward for next WORD, check first character
- Modifies: Current token value

**Why Strategy Pattern**: 
- Easy to add new transformations
- Processors are independent and testable
- Execution order is configurable

---

### 6. **Rendering Layer** (`formatter/renderer.go`)
**Pattern**: Pipeline Stage (Final Assembly)

**Responsibility**: Concatenate token values into final string

**Why Pipeline Here**: Pure data assembly, no logic needed

---

## Key Architectural Decisions

### 1. **Immutable Input, Mutable Output**
- Input token array never changes
- FSM builds result array incrementally
- Processors create modified copies

**Benefit**: Easier debugging, no side effects

### 2. **Single-Pass Processing**
- Each token processed exactly once
- No backtracking or re-scanning
- Processors look backward in result array

**Benefit**: O(n) time complexity, efficient memory usage

### 3. **Quote Segmentation**
- Pre-split tokens by quotes before FSM
- Process each segment independently
- Merge results after processing

**Benefit**: Transformations respect quote boundaries naturally

### 4. **Processor Ordering**
- Conversion before case (numbers must convert before case changes)
- Article correction after FSM (needs final token values)

**Benefit**: Correct transformation semantics

### 5. **Backward-Looking Transformations**
- Tags modify previous tokens in result array
- No lookahead needed during FSM execution

**Benefit**: Simpler state machine, clearer logic

---

## Data Flow Example

**Input**: `"hello 1E (hex) world (up)"`

### Step 1: Tokenization
```
[WORD:"hello", WHITESPACE:" ", WORD:"1E", WHITESPACE:" ", TAG:"(hex)", 
 WHITESPACE:" ", WORD:"world", WHITESPACE:" ", TAG:"(up)"]
```

### Step 2: Punctuation Normalization
```
(No changes - no punctuation)
```

### Step 3: Quote Spacing
```
(No changes - no quotes)
```

### Step 4: Quote Segmentation
```
Segment 1: [WORD:"hello", WHITESPACE:" ", WORD:"1E", WHITESPACE:" ", TAG:"(hex)", 
            WHITESPACE:" ", WORD:"world", WHITESPACE:" ", TAG:"(up)"]
```

### Step 5: FSM Processing

| State | Token | Action | Result Array |
|-------|-------|--------|--------------|
| READING | - | Check position | [] |
| EVALUATING | WORD:"hello" | Append | [WORD:"hello"] |
| READING | - | Check position | [WORD:"hello"] |
| EVALUATING | WHITESPACE:" " | Append | [WORD:"hello", WHITESPACE:" "] |
| READING | - | Check position | [...] |
| EVALUATING | WORD:"1E" | Append | [..., WORD:"1E"] |
| READING | - | Check position | [...] |
| EVALUATING | WHITESPACE:" " | Append | [..., WORD:"1E", WHITESPACE:" "] |
| READING | - | Check position | [...] |
| EVALUATING | TAG:"(hex)" | Move to EDITING | [...] |
| EDITING | TAG:"(hex)" | ConversionProcessor: "1E"→"30", remove " " | [..., WORD:"30"] |
| READING | - | Check position | [...] |
| EVALUATING | WHITESPACE:" " | Append | [..., WORD:"30", WHITESPACE:" "] |
| READING | - | Check position | [...] |
| EVALUATING | WORD:"world" | Append | [..., WORD:"world"] |
| READING | - | Check position | [...] |
| EVALUATING | WHITESPACE:" " | Append | [..., WORD:"world", WHITESPACE:" "] |
| READING | - | Check position | [...] |
| EVALUATING | TAG:"(up)" | Move to EDITING | [...] |
| EDITING | TAG:"(up)" | CaseProcessor: "world"→"WORLD", remove " " | [..., WORD:"WORLD"] |
| READING | - | Position >= len | [...] |
| DONE | - | Return result | [WORD:"hello", WHITESPACE:" ", WORD:"30", WHITESPACE:" ", WORD:"WORLD"] |

### Step 6: Article Correction
```
(No changes - no articles)
```

### Step 7: Rendering
```
"hello 30 WORLD"
```

---

## Why This Hybrid Approach?

### Pipeline Advantages (Used for preprocessing/postprocessing)
✅ Simple, predictable flow  
✅ Easy to understand and debug  
✅ Stateless transformations  
✅ Efficient for batch operations  

### FSM Advantages (Used for core transformations)
✅ Handles complex state-dependent logic  
✅ Backward-looking transformations  
✅ Flexible processor ordering  
✅ Clear error handling  
✅ Single-pass efficiency  

### Hybrid Benefits
✅ **Speed**: Single-pass processing for small files  
✅ **Scalability**: Can process segments in parallel (future enhancement)  
✅ **Maintainability**: Clear separation of concerns  
✅ **Extensibility**: Easy to add new processors or pipeline stages  
✅ **Testability**: Each component independently testable  

---

## Testing Strategy

### Unit Tests
- **Tokenizer**: Test regex patterns and token types
- **Processors**: Test each transformation in isolation
- **FSM States**: Test state transitions
- **Pipeline Stages**: Test punctuation and quote normalization

### Integration Tests
- **Full Pipeline**: Test complete text transformations
- **Edge Cases**: Empty input, malformed tags, nested quotes
- **Property Tests**: Verify invariants (e.g., token count consistency)

### Golden Tests
- **Reference Outputs**: Compare against known-good outputs
- **Regression Prevention**: Ensure changes don't break existing behavior

---

## Future Enhancements

### 1. Parallel Segment Processing
```go
// Process quote segments concurrently
var wg sync.WaitGroup
results := make([][]Token, len(segments))
for i, segment := range segments {
    wg.Add(1)
    go func(idx int, seg []Token) {
        defer wg.Done()
        results[idx] = processSingleSegment(seg)
    }(i, segment)
}
wg.Wait()
```

### 2. Streaming for Large Files
- Process file in chunks
- Maintain state across chunks
- Write output incrementally

### 3. Custom Processor Registration
```go
formatter.RegisterProcessor("custom", &MyProcessor{})
```

### 4. Error Recovery Strategies
- Continue on error vs fail fast
- Collect all errors vs stop at first

---

## Conclusion

Your architecture successfully combines:
- **Pipeline simplicity** for straightforward transformations
- **FSM intelligence** for complex, state-dependent logic
- **Strategy pattern** for extensible transformations
- **TDD principles** for reliable, tested code

This hybrid approach achieves both **speed** (single-pass) and **flexibility** (pluggable processors), making it suitable for both small files and large-scale processing.