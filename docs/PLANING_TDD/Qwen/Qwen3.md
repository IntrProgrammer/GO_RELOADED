Below is a carefully designed sequence of **14 small, incremental Agile tasks** for an entry-level developer. Each task follows TDD principles (test-first), builds incrementally, and collectively leads to a fully functional text formatter. Tasks are ordered to reflect increasing complexity, leveraging a hybrid FSM/Pipeline architecture for optimal performance.

---

### **Task 1: Implement Basic Tag Conversion (hex/bin)**  
**Functionality**: Convert hexadecimal/binary numbers in `"(hex)"` or `"(bin)"` tags to decimal numbers.  
**Test Writing**:  
```go
func TestTagConversion(t *testing.T) {
  // Test cases: 
  // 1. hex conversion: "1E (hex) files" → "30 files"
  // 2. bin conversion: "10 (bin) years" → "2 years"
  // 3. Boundary: No leading space before tag ("1E(hex)files" → "30files")
}
```  
**Implementation Goal**:  
- Parse tag locations via regex (`\s*\(hex\)\s*` or `\(bin\)`).  
- Extract preceding word using sliding window (max 1 token back).  
- Convert hex/bin strings to decimal (use `strconv.ParseInt`).  
- Return modified string (keep non-tag text unchanged).  
**Validation**:  
- Run tests; verify output matches expected results.  
- Handle edge cases: empty input, invalid numbers, and spaces around tags.  

---

### **Task 2: Add Case-Conversion Tags (up/low)**  
**Functionality**: Convert single-word case (`"(up)"`, `"(low)"`).  
**Test Writing**:  
```go
func TestCaseConversion(t *testing.T) {
  // Test cases:
  // 1. "(up)": "Ready, set, go (up) !" → "Ready, set, GO!"
  // 2. "(low)": "I should stop SHOUTING (low)" → "I should stop shouting"
  // 3. Multi-word: "(up, 2)": "This is so exciting (up, 2)" → "This is SO EXCITING"
}
```  
**Implementation Goal**:  
- Implement `"(up)"`/`"(low)"` logic (no `"(cap)"` yet).  
- Support multi-word count via `"(up, 2)"`.  
- Use token buffer to store last N words (N=1 for now).  
- Apply case conversion to specified words (preserve non-alphanumeric chars).  
**Validation**:  
- Verify case changes; ensure buffer correctly tracks prior words.  
- Validate multi-word count works for N=1 (single word).  

---

### **Task 3: Extend Case-Conversion with "(cap)"**  
**Functionality**: Capitalize single word (`"(cap)"`).  
**Test Writing**:  
```go
func TestCapConversion(t *testing.T) {
  // Test cases:
  // 1. "(cap)": "Welcome to the brooklyn bridge (cap)" → "Welcome to the Brooklyn Bridge"
  // 2. Multi-word: "(cap, 2)": "a b c (cap, 2)" → "A B c"
}
```  
**Implementation Goal**:  
- Add `"(cap)"` tag logic (capitalizes first letter, rest lowercase).  
- Extend token buffer to support N=2 for multi-word tags (e.g., `"(cap, 2)"`).  
- Reuse buffer from Task 2 to avoid code duplication.  
**Validation**:  
- Confirm proper capitalization; validate multi-word behavior.  
- Test edge cases: words with non-ASCII chars, empty buffers.  

---

### **Task 4: Implement Punctuation Spacing Rule**  
**Functionality**: Standardize spacing around punctuation (e.g., `", " → ","`).  
**Test Writing**:  
```go
func TestPunctuationSpacing(t *testing.T) {
  // Test cases:
  // 1. Standard: "I was sitting over there ,and then BAMM !!"
  // → "I was sitting over there, and then BAMM!!"
  // 2. Grouped punctuation: "I was thinking ... You were right"
  // → "I was thinking... You were right"
  // 3. Edge: "Hello,world!" (no space after comma)
}
```  
**Implementation Goal**:  
- Scan for punctuation (`, . ! ? ;`) using regex.  
- Insert space after punctuation *only* if preceding char is a letter/number.  
- For grouped punctuation (e.g., `...`), treat as a single unit.  
**Validation**:  
- Confirm spacing is fixed; verify grouped punctuation remains intact.  
- Ensure no extra spaces added between words.  

---

### **Task 5: Add Single-Quote Space Cleanup**  
**Functionality**: Remove spaces after opening quotes and before closing quotes.  
**Test Writing**:  
```go
func TestQuoteSpacing(t *testing.T) {
  // Test cases:
  // 1. Before quote: "As Elton John said: ' I am the most well-known..."
  // → "As Elton John said: 'I am the most well-known..."
  // 2. After quote: "...world ' " → "...world' "
}
```  
**Implementation Goal**:  
- Detect quote boundaries with regex (`' + [\s]*`).  
- Replace `'\s+` → `'` (after opening quote).  
- Replace `[\s]+'` → `'` (before closing quote).  
- **Note**: Apply *after* other rules (per problem statement).  
**Validation**:  
- Verify spaces are removed *only* at quote boundaries.  
- Preserve internal spacing inside quotes.  

---

### **Task 6: Implement "a" → "an" Transformation**  
**Functionality**: Convert "a" to "an" before vowel-starting words.  
**Test Writing**:  
```go
func TestAnConversion(t *testing.T) {
  // Test cases:
  // 1. Basic: "A amazing rock!" → "An amazing rock!"
  // 2. No change: "A apple" → "A apple" (if next word starts with consonant)
  // 3. Edge: "A university" → "An university" (u is vowel)
}
```  
**Implementation Goal**:  
- Scan for standalone "a" (e.g., `a` followed by space/word).  
- Check next word starts with vowel (`a, e, i, o, u`) or `h`.  
- Replace `a` with `an` if condition met.  
**Validation**:  
- Confirm correct replacements; verify no unintended changes.  
- Handle uppercase "A" and non-ASCII vowels.  

---

### **Task 7: Combine Tag Rules into Unified Processor**  
**Functionality**: Apply all tag rules (`hex`, `bin`, `up`, `low`, `cap`, `multi-word`).  
**Test Writing**:  
```go
func TestAllTags(t *testing.T) {
  // Test case: "10 (bin) years (up, 2) I am here"
  // → "2 years I AM HERE" (after all tag rules)
}
```  
**Implementation Goal**:  
- Create a single function `ApplyTagRules` that sequences:  
  1. Hex/bin conversion.  
  2. Case conversion (including multi-word).  
- Use token buffer (size=10) to support multi-word tags.  
- Maintain state across rules (e.g., tags modify text that affects other rules).  
**Validation**:  
- Verify all tag rules work together.  
- Test edge: overlapping tags (e.g., `1E (hex) (bin)`) → process in order.  

---

### **Task 8: Integrate Punctuation with Tag Processor**  
**Functionality**: Apply punctuation spacing *after* tag processing.  
**Test Writing**:  
```go
func TestPunctuationAfterTags(t *testing.T) {
  // Test case: "1E (hex) files were added, then (up)!"
  // → "30 files were added, then GO!" (after tags and punctuation)
}
```  
**Implementation Goal**:  
- Sequence rules: `ApplyTagRules` → `ApplyPunctuationSpacing`.  
- Ensure punctuation rule ignores quotes (per problem statement).  
**Validation**:  
- Confirm punctuation spacing works on modified text (e.g., after tag conversions).  
- Verify quotes remain untouched during punctuation pass.  

---

### **Task 9: Integrate Quote Cleaning After Punctuation**  
**Functionality**: Apply single-quote spacing *after* punctuation.  
**Test Writing**:  
```go
func TestQuoteCleaningOrder(t *testing.T) {
  // Test case: "As Elton John said: ' I am the most well-known... '"
  // → "As Elton John said: 'I am the most well-known...'"
  // *After* punctuation spacing (if any)
}
```  
**Implementation Goal**:  
- Add `ApplyQuoteSpacing` after `ApplyPunctuationSpacing`.  
- Use regex-based replacement for quote boundaries.  
**Validation**:  
- Ensure quotes are processed *after* punctuation.  
- Verify spaces only removed at quote boundaries (not inside).  

---

### **Task 10: Implement "a" → "an" After Quotes**  
**Functionality**: Apply "a" → "an" *after* quote cleaning.  
**Test Writing**:  
```go
func TestAnConversionAfterQuotes(t *testing.T) {
  // Test case: "a university (low) in the quote 'I am here'"
  // → "an university in the quote 'i am here'"
}
```  
**Implementation Goal**:  
- Add `ApplyAnConversion` as the final rule.  
- Ensure it processes text *after* quote cleaning (so quotes are handled correctly).  
**Validation**:  
- Confirm "a" → "an" applies to non-quoted text (e.g., outside quotes).  
- Verify no interference with quote content.  

---

### **Task 11: Build FSM State Machine for Workflow**  
**Functionality**: Manage rule execution order using FSM states.  
**Test Writing**:  
```go
func TestFSMOrder(t *testing.T) {
  // Test sequence: 
  // 1. Tag rules → 2. Punctuation → 3. Quotes → 4. "a"→"an"
  // Verify output matches expected state transitions.
}
```  
**Implementation Goal**:  
- Design FSM states:  
  - `INITIAL` → `TAG_PROCESSING` → `PUNCTUATION` → `QUOTE_CLEANUP` → `FINAL`.  
- Transition states only when a rule completes.  
- Use FSM to enforce rule order *and* handle quotes.  
**Validation**:  
- Confirm rules run in strict sequence (no interleaving).  
- Verify state transitions for quote boundaries.  

---

### **Task 12: Add Chunk Processing for Large Files**  
**Functionality**: Process files in chunks for memory efficiency.  
**Test Writing**:  
```go
func TestChunkProcessing(t *testing.T) {
  // Test case: 100MB file → verify no OOM errors
  // Chunk size: 1MB (or similar)
}
```  
**Implementation Goal**:  
- Split file into chunks (e.g., 1MB each) using `bufio.Scanner`.  
- Apply FSM rules to each chunk independently.  
- **Critical**: Maintain context across chunks (e.g., tags spanning chunk boundaries).  
**Validation**:  
- Test with large files (e.g., 10MB text).  
- Ensure chunk boundaries don't break tag/punctuation logic.  

---

### **Task 13: Implement File Input/Output Pipeline**  
**Functionality**: Read file → apply formatter → write output.  
**Test Writing**:  
```go
func TestFileIO(t *testing.T) {
  // Create test file: "1E (hex) files were added"
  // Verify output: "30 files were added"
}
```  
**Implementation Goal**:  
- Build CLI tool:  
  - `formatter <input.txt> <output.txt>`  
- Use Task 12's chunk processing in `main()`.  
- Handle errors (e.g., file not found).  
**Validation**:  
- Run on sample file; verify output matches expectations.  
- Test with multi-file input (`formatter file1.txt file2.txt`).  

---

### **Task 14: Full Test Suite and Validation**  
**Functionality**: Execute all rules on real-world text.  
**Test Writing**:  
```go
func TestEndToEnd(t *testing.T) {
  // Test full workflow on mixed-input file:
  // "a amazing rock (up,2) ' I am here ' (bin)10 (an)"
}
```  
**Implementation Goal**:  
- Run all 4 rules sequentially.  
- Verify output matches:  
  - `an` for "a", `30` for `1E (hex)`, `GO!` for `(up)`, and `I am here` cleaned.  
- **Critical**: No rule conflicts (e.g., tags inside quotes work as expected).  
**Validation**:  
- Check all test cases pass.  
- Validate performance: < 100ms on 1MB file.  

---

### **Key Architecture Notes**  
- **Hybrid FSM/Pipeline**:  
  - **FSM** manages rule order and state (e.g., tag context, quote boundaries).  
  - **Pipeline** processes chunks for memory efficiency (Task 12).  
- **Order Enforcement**:  
  Rules run in sequence:  
  ```mermaid  
  graph LR  
    A[Tag Rules] --> B[Punctuation]  
    B --> C[Quote Cleaning]  
    C --> D[An Conversion]  
  ```  
- **Why TDD?**  
  Tests expose edge cases (e.g., tags spanning chunks, quote boundaries) early.  
- **Entry-Level Focus**:  
  Each task has **isolated scope** (e.g., only tag rules for Task 1), reducing cognitive load.  

This sequence ensures a smooth learning curve while delivering a robust formatter. Tasks 1-6 build core logic; Tasks 7-12 integrate state/pipeline; Task 14 validates completeness. All tests pass before moving to next task.
