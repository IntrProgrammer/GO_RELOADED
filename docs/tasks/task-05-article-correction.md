# Task 5 — Article Correction (a/an)

## Objective
Implement post-processor to correct grammatical article usage: "a" before consonants, "an" before vowels/h.

## Prerequisites
- Task 4 completed (Case Processor)

## Deliverables
- [ ] CorrectArticles function
- [ ] Vowel/h detection logic
- [ ] Tests for article correction
- [ ] Special case handling (8, 11, 18)

## Correction Rules
- "a" → "an" before vowels (a, e, i, o, u)
- "a" → "an" before h
- "a" → "an" before numbers starting with vowel sound (8, 11, 18)
- "A" → "An" (preserve capitalization)
- Skip punctuation when looking ahead

## Implementation Steps

### Step 1: Write Article Correction Tests
File: `fsm/processors_test.go` (add to existing)
```go
func TestCorrectArticles(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
	}{
		{
			name: "a before vowel",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "an"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
			},
		},
		{
			name: "a before consonant unchanged",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "book"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "book"},
			},
		},
		{
			name: "capital A before vowel",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "A"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "elephant"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "An"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "elephant"},
			},
		},
		{
			name: "a before h",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "hour"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "an"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "hour"},
			},
		},
		{
			name: "a before 8",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "8"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "an"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "8"},
			},
		},
		{
			name: "a with punctuation between",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.PUNCTUATION, ","},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.PUNCTUATION, ","},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CorrectArticles(tt.tokens)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CorrectArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

### Step 2: Implement Vowel Detection Helper
File: `fsm/processors.go` (add to existing)
```go
func startsWithVowelOrH(s string) bool {
	if len(s) == 0 {
		return false
	}
	first := unicode.ToLower(rune(s[0]))
	if first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u' || first == 'h' {
		return true
	}
	// Special cases: 8, 11, 18 (vowel sound)
	if first == '8' || (len(s) >= 2 && s[0] == '1' && (s[1] == '1' || s[1] == '8')) {
		return true
	}
	return false
}
```

### Step 3: Implement CorrectArticles
File: `fsm/processors.go` (add to existing)
```go
func CorrectArticles(tokens []tokenizer.Token) []tokenizer.Token {
	result := make([]tokenizer.Token, len(tokens))
	copy(result, tokens)

	for i := 0; i < len(result); i++ {
		if result[i].Type != tokenizer.WORD {
			continue
		}

		word := result[i].Value
		if strings.ToLower(word) != "a" {
			continue
		}

		// Find next WORD token (skip punctuation/whitespace)
		nextWordIdx := -1
		for j := i + 1; j < len(result); j++ {
			if result[j].Type == tokenizer.WORD {
				nextWordIdx = j
				break
			}
		}

		if nextWordIdx == -1 {
			continue
		}

		// Check if next word starts with vowel or h
		if startsWithVowelOrH(result[nextWordIdx].Value) {
			if word == "a" {
				result[i].Value = "an"
			} else if word == "A" {
				result[i].Value = "An"
			}
		}
	}

	return result
}
```

### Step 4: Integrate into FSM
File: `fsm/fsm.go` (update existing Run method)
```go
func (f *FSM) Run() {
	for f.state != StateDone && f.state != StateError {
		f.step()
	}
	// Post-process: Apply article correction
	f.result = CorrectArticles(f.result)
}
```

### Step 5: Integration Test
File: `fsm/fsm_test.go` (add to existing)
```go
func TestFSMWithArticleCorrection(t *testing.T) {
	tests := []struct {
		name   string
		tokens []tokenizer.Token
		want   []tokenizer.Token
	}{
		{
			name: "a before apple",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "an"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "apple"},
			},
		},
		{
			name: "a before book unchanged",
			tokens: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "book"},
			},
			want: []tokenizer.Token{
				{tokenizer.WORD, "a"},
				{tokenizer.WHITESPACE, " "},
				{tokenizer.WORD, "book"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsm := New(tt.tokens)
			fsm.Run()
			
			result := fsm.Result()
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Result = %v, want %v", result, tt.want)
			}
		})
	}
}
```

## Verification Commands
```bash
go test ./fsm/... -v -run TestCorrectArticles
go test ./fsm/... -v -run TestFSMWithArticleCorrection
go test ./...
```

## Success Criteria
- "a" → "an" before vowels (a, e, i, o, u)
- "a" → "an" before h
- "a" → "an" before 8, 11, 18
- Capitalization preserved (A → An)
- Consonants unchanged
- Punctuation between article and word handled correctly
- All tests pass

## TDD Workflow
1. ✅ RED: Write failing article correction tests
2. ✅ GREEN: Implement CorrectArticles
3. ✅ REFACTOR: Optimize vowel detection

## Git Commit Message
```
feat: implement article correction for a/an

- Add CorrectArticles post-processor function
- Implement startsWithVowelOrH detection
- Handle special cases (8, 11, 18)
- Preserve capitalization (A → An)
- Skip punctuation when looking ahead
```
