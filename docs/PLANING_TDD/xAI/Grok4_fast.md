### Task 1: Project Setup and Basic Go Module Initialization
- **Functionality Description:** Establish the foundational Go project structure, including module initialization, directory layout, and a simple main function to read a file, process it minimally (e.g., echo the content), and write output to a new file. This sets up the entry point for the text formatter without any rules yet.
- **TDD Step:** Write unit tests in `main_test.go` to verify: (1) the module loads without errors, (2) a sample input file is read correctly as a string, (3) the output file is created with identical content (basic echo), and (4) error handling for non-existent input files. Use `testing` package and `os`/`ioutil` for file I/O mocks.
- **Implementation Goal:** In `main.go`, initialize the Go module with `go mod init textformatter`, create directories like `cmd/textformatter/` and `internal/processor/`, implement a basic `ProcessFile(inputPath, outputPath string)` function that reads the entire file content using `os.ReadFile`, echoes it to output using `os.WriteFile`, and call it from `main`.
- **Validation:** Run `go test ./...` to ensure all tests pass. Manually test by running the binary with a sample input file (e.g., create `test_input.txt` with plain text) and verify the output file matches the input unchanged.

### Task 2: Implement Basic Finite State Machine Skeleton for Single-File Processing
- **Functionality Description:** Create a core FSM structure to manage states for reading, processing, and writing text chunks. Start with three states: `Reading` (load text), `Processing` (apply rules later), and `Writing` (output result). This leans into the FSM for flexible, state-based flow on single files.
- **TDD Step:** In `internal/processor/fsm.go` and `fsm_test.go`, write tests for: (1) FSM initializes in `Reading` state, (2) transitions to `Processing` after reading succeeds, (3) transitions to `Writing` after processing (mock for now), (4) handles invalid transitions (e.g., error on direct `Writing` from `Reading`), and (5) FSM execution returns the final output string or error.
- **Implementation Goal:** Define a `FSM` struct with `currentState` enum (use `string` constants: "reading", "processing", "writing") and methods like `Read(input string) error`, `Process() string`, `Write(output string) error`. Implement transitions using a switch on current state, starting with basic string passing through states (no real processing yet).
- **Validation:** Run `go test ./internal/processor/...` to confirm tests pass. Integrate with Task 1's main: update `ProcessFile` to use FSM, test on a small file to ensure input reads, FSM runs through states, and output matches input.

### Task 3: Tokenize Input Text into Words and Punctuation for FSM Processing
- **Functionality Description:** In the `Processing` state, split input text into tokens (words, punctuation, tags) to prepare for rule application. Treat punctuation groups (e.g., `...`, `!!`) as single tokens and preserve spaces minimally for reassembly.
- **TDD Step:** In `internal/processor/tokenizer.go` and `tokenizer_test.go`, write tests for: (1) simple sentence tokenizes to words and spaces, (2) punctuation attaches correctly (e.g., "hello ," -> ["hello", ","]), (3) groups like "BAMM!!" -> single token "BAMM!!", (4) handles mixed cases (e.g., "a (up)"), and (5) reassembles tokens to original text.
- **Implementation Goal:** Implement `Tokenize(text string) []Token` where `Token` is a struct with `Value string` and `Type string` (e.g., "word", "punct", "tag"). Use regex (e.g., `\s+` for spaces, `([.!?,;]+)` for punct groups, `$$.*?$$` for tags) combined with string splitting. Add `Reassemble(tokens []Token) string` to join with single spaces where needed.
- **Validation:** Run `go test ./internal/processor/...` to ensure tests pass. Update FSM's `Process` to tokenize and reassemble (still echo), test full pipeline on sample text with punctuation/tags to verify no loss of content.

### Task 4: Implement Rule 1 - Basic Tag Processing (Single-Word: hex and bin Conversions)
- **Functionality Description:** In the FSM `Processing` state, detect `(hex)` and `(bin)` tags, convert the preceding word from hex/base-16 or binary/base-2 to decimal, and replace it. Focus on single-word tags first.
- **TDD Step:** In `internal/processor/rules.go` and `rules_test.go`, write tests for: (1) "1E (hex)" -> "30", (2) "10 (bin)" -> "2", (3) invalid hex/bin inputs error gracefully, (4) non-numeric words before tag unchanged, and (5) integration with tokenizer (pass tokens, get updated tokens).
- **Implementation Goal:** Add `ProcessTags(tokens []Token) []Token` function. Scan tokens for tag type, extract preceding word token, use `strconv.ParseInt` with base 16/2 for conversion, replace `Value` in place. Update FSM `Process` to call this after tokenizing.
- **Validation:** Run `go test ./internal/processor/...` to pass tests. Test full app on sample file like "1E (hex) files added" -> output "30 files added". Verify FSM transitions correctly.

### Task 5: Extend Rule 1 - Text Case Tags (up, low, cap for Single Word)
- **Functionality Description:** Add support for `(up)`, `(low)`, and `(cap)` tags, modifying the case of the single preceding word: uppercase all, lowercase all, or capitalize first letter.
- **TDD Step:** Extend `rules_test.go` with tests for: (1) "go (up)" -> "GO", (2) "SHOUT (low)" -> "shout", (3) "brooklyn (cap)" -> "Brooklyn", (4) tags ignored if no preceding word, and (5) case-insensitive tag matching.
- **Implementation Goal:** In `ProcessTags`, add cases for "up"/"low"/"cap": use `strings.ToUpper`, `strings.ToLower`, and manual capitalization (`unicode.ToUpper` on first rune + `ToLower` on rest). Ensure it handles only single word (no comma yet).
- **Validation:** Run `go test ./internal/processor/...`. Test app on inputs like "Ready, set, go (up) !" -> "Ready, set, GO !", confirm FSM processes without breaking prior rules.

### Task 6: Implement Rule 1 Extension - Multi-Word Case Tags (up/low/cap with Count)
- **Functionality Description:** Handle multi-word tags like `(up, 2)` by applying case changes to the specified number of preceding words.
- **TDD Step:** Add tests to `rules_test.go` for: (1) "is so exciting (up, 2)" -> "is SO EXCITING", (2) invalid count (e.g., "(up, abc)") errors, (3) count exceeds available words applies to all, and (4) combines with single-word tags.
- **Implementation Goal:** Parse tag as `(tag, num)` using regex or string split, count back `num` words from tag position in tokens, apply case function to each. Update `ProcessTags` to handle comma-separated args.
- **Validation:** Run `go test ./internal/processor/...`. Test on multi-word sample, ensure integration with hex/bin doesn't conflict (process all tags in one pass).

### Task 7: Implement Rule 2 - Punctuation Spacing Normalization
- **Functionality Description:** After tag processing, ensure punctuation (.,!?,;) attaches to the previous word (no space before) and has exactly one space after. Treat groups (e.g., `...`, `!?.`) as single units.
- **TDD Step:** In new `punctuation.go` and `punctuation_test.go`, test: (1) "there , and" -> "there, and", (2) "BAMM !!" -> "BAMM!! ", (3) "thinking ... You" -> "thinking... You", (4) standalone punct unchanged, and (5) reassemble with proper spacing.
- **Implementation Goal:** Add `NormalizePunctuation(tokens []Token) []Token`. Scan for punct tokens, merge with prior word if space-separated, insert single space after. Use tokenizer's punct detection. Call after tags in FSM `Process`.
- **Validation:** Run `go test ./internal/processor/...`. Test full app on "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!", verify no interference with tags.

### Task 8: Implement Rule 3 - Single Quotes Spacing Cleanup
- **Functionality Description:** Detect single-quoted sections, remove extra spaces right after opening `'` and before closing `'`, while applying all other rules inside quotes.
- **TDD Step:** In `quotes.go` and `quotes_test.go`, test: (1) "' I am ... world '" -> "'I am ... world'", (2) no extra spaces unchanged, (3) nested quotes not handled (error or ignore), and (4) tags/punct inside quotes process first (mock integration).
- **Implementation Goal:** Add `CleanQuotes(text string) string` using regex to find `'([^']*)'` and trim spaces in capture group. But integrate via tokens: scan for quote tokens, process inner tokens with all rules, then reassemble without inner spaces. Call in FSM `Process` after punct.
- **Validation:** Run `go test ./internal/processor/...`. Test on "As Elton John said: ' I am the most well-known homosexual in the world '" -> proper output, ensure rules like tags apply inside (e.g., add a tag test case).

### Task 9: Implement Rule 4 - "a" to "an" Transformation
- **Functionality Description:** Scan for "a" followed by a word starting with vowel (a,e,i,o,u) or 'h', change to "an". Apply after other rules to avoid interference.
- **TDD Step:** In `articles.go` and `articles_test.go`, test: (1) "A amazing" -> "An amazing", (2) "a house" -> "an house" (h rule), (3) "A book" unchanged, (4) case-insensitive for "a/A", and (5) doesn't change inside quotes or after tags.
- **Implementation Goal:** Add `TransformArticles(tokens []Token) []Token`. Scan for word "a" token, check next token's first letter (vowel or 'h'), replace if match. Call last in FSM `Process`.
- **Validation:** Run `go test ./internal/processor/...`. Test on "There it was. A amazing rock!" -> "There it was. An amazing rock!", confirm order: tags/punct/quotes first.

### Task 10: Enhance FSM for Hybrid Pipeline Support (Multiple Files/Chunks)
- **Functionality Description:** Extend FSM to handle multiple files or large-file chunks in a pipeline-like flow: read chunks sequentially, process each in FSM, aggregate outputs. Use FSM for per-chunk flexibility, pipeline for batching.
- **TDD Step:** In `fsm.go` and `fsm_test.go`, add tests for: (1) `ProcessMultipleFiles(paths []string) []string` outputs per file, (2) chunk large file (e.g., split by lines/bytes) into FSM, (3) errors in one chunk don't stop others, and (4) final multi-output write.
- **Implementation Goal:** Add `ProcessChunk(chunk string) string` to FSM (reuse existing states). For multiples, loop over files/chunks, run FSM per item. In main, support CLI flag for multiple inputs (e.g., `go run main.go -files file1.txt file2.txt` using `flag` package).
- **Validation:** Run `go test ./internal/processor/...`. Test on two small files with mixed rules, verify outputs correct and aggregated. For large file, chunk by 1KB, ensure full processing.

### Task 11: Full Integration, Error Handling, and End-to-End Validation
- **Functionality Description:** Ensure all rules integrate seamlessly in FSM `Process` order: tokenize -> tags -> punct -> quotes -> articles -> reassemble. Add error states in FSM for malformed input (e.g., unclosed tags).
- **TDD Step:** Create integration tests in `internal/processor/integration_test.go`: (1) full sample with all rules (tags, punct, quotes, articles), (2) error cases (invalid hex, unclosed quote), (3) multi-file processing, and (4) performance edge (1000-line file chunks).
- **Implementation Goal:** In FSM, add `Error` state with transitions (e.g., from `Processing` on parse fail). Update main to log errors but continue for multiples. Order rule calls explicitly in `Process`.
- **Validation:** Run `go test ./...` for all tests. Create comprehensive test file with all rules combined (e.g., "a (hex) (up,2) amazing ! 'test'", process to expected). Run binary on varied inputs (single/multi-file, chunked), confirm 100% test coverage with `go test -cover`, and manually inspect outputs match spec.
