# GO_RELOADED

A text formatter with tag-based transformations following TDD principles.

## Features

- **Case Transformations**: `(up)`, `(low)`, `(cap)` with multi-word support
- **Number Conversions**: `(hex)`, `(bin)` to decimal
- **Grammar Correction**: Automatic a/an article correction
- **Quote Processing**: Apply rules inside quoted text
- **Punctuation Normalization**: Automatic spacing fixes

**[Read Full Problem Description](docs/Analysis/Understunding%20_the_Problem.md)**

## Installation

### Prerequisites
- Go 1.25.3 or higher

```bash
git clone https://github.com/IntrProgrammer/GO_RELOADED
```

```bash
go build -o go-reloaded ./cmd/go-reloaded
```

## Usage

```bash
./go-reloaded input.txt output.txt
```

## Development

### Project Structure
```
GO_RELOADED/
├── cmd/                 # Entry point (main.go)
├── docs/                # Project documentation
│   ├── Analysis/        # Design docs, problem description, FSM
│   └── test/            # Usage  test cases
├── formatter/           # Text formatting logic (pipeline functions)
├── fsm/                 # Finite state machine implementation
├── tokenizer/           # Tokenization and token types
├── COPYING.md           # License (GPL-3.0)
├── README.md            # Project overview
├── go.mod               # Go module definition
├── go.sum               # Go dependencies checksums
└── go-reloaded          # Compiled binary

```
## Usefull documents

- 📄 **[Architecture Overview](docs/Analysis/Architecture_Type.md)**
- 📄 **[FSM Design Document](docs/Analysis/FSM%20implementation.md)**
- 📄 **[Full License Text](COPYING.md)**

---

**Built with TDD principles | Hybrid FSM-Pipeline Architecture | GPL-3.0 Licensed**