# FSM — Compact Notes

## Flow of the FSM
Input Tokens → FSM orchestrates → Processors transform → Output Tokens

<div align="center">

```
  Input tokens
|
v
 +----- FSM -----+
 |  handleReading |
 | -> handleEval  |
 | -> processors  |
 | result updated |
 +----------------+
|
v
  Processed tokens

```
</div>
## Purpose
- Orchestrates token processing: reads tokens, delegates token-specific work to processors, collects output, and reports errors.

## Core Fields
- **state**: current FSM state (Reading, Evaluating, Editing, Done, Error)
- **tokens**: input token slice
- **position**: index of next token to evaluate (0 <= position <= len(tokens))
- **result**: accumulated output tokens
- **errorMsg**: non-empty when in Error state
- **processors**: ordered list of Processor implementations


## Processor Contract
```go
Process(result []tokenizer.Token, currentToken tokenizer.Token) (modified []tokenizer.Token, handled bool)
```

- Inputs: current FSM result and the token under evaluation.

- If handled == true, FSM replaces f.result with modified, advances position, and stops checking further processors.

- Processors must NOT modify FSM.position directly; return the intended result slice instead.