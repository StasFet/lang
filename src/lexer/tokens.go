package lexer

import "fmt"

type TokenKind int

// define kinda of tokens
const (
	EOF TokenKind = iota
	NUMBER
	STRING
	IDENTIFIER

	// grouping
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	// equivalence
	ASSIGNMENT // =
	EQUALS     // ==
	NOT        // !
	NOT_EQUALS // !=

	// comparison
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL

	// boolean logic
	OR
	AND

	// symbols
	DOT
	DOT_DOT
	SEMI_COLON
	COLON
	QUESTION
	COMMA

	// shorthand
	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS
	SLASH_EQUALS
	STAR_EQUALS

	// math
	PLUS
	DASH
	SLASH
	STAR
	PERCENT

	// reserved keyworks
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FUNCTION
	IF
	ELSE
	FOR
	WHILE
	FOREACH
	EXPORT
	TYPEOF
	IN
)

type Token struct {
	Kind  TokenKind
	Value string
}

// returns true if token.Type is one the params, false otherwise
func (token Token) kindIsAmong(expectedToken ...TokenKind) bool {
	for _, expected := range expectedToken {
		if expected == token.Kind {
			return true
		}
	}
	return false
}

// debug function that prints the kind and value (if applicable) of a token
func (token Token) Debug() {
	if token.kindIsAmong(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s, ()\n", TokenKindString(token.Kind))
	}
}

func NewToken(kind TokenKind, value string) Token {
	return Token{kind, value}
}

// ugly little function to get the name of a TokenKind (did not wanna use reflect for this)
func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_CURLY:
		return "open_curly"
	case CLOSE_CURLY:
		return "close_curly"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case ASSIGNMENT:
		return "assignment"
	case EQUALS:
		return "equals"
	case NOT_EQUALS:
		return "not_equals"
	case NOT:
		return "not"
	case LESS:
		return "less"
	case LESS_EQUAL:
		return "less_equal"
	case GREATER:
		return "greater"
	case GREATER_EQUAL:
		return "greater_equal"
	case OR:
		return "or"
	case AND:
		return "and"
	case DOT:
		return "dot"
	case DOT_DOT:
		return "dot_dot"
	case SEMI_COLON:
		return "semi_colon"
	case COLON:
		return "colon"
	case QUESTION:
		return "question"
	case COMMA:
		return "comma"
	case PLUS_PLUS:
		return "plus_plus"
	case MINUS_MINUS:
		return "minus_minus"
	case PLUS_EQUALS:
		return "plus_equals"
	case MINUS_EQUALS:
		return "minus_equals"
	case PLUS:
		return "plus"
	case DASH:
		return "dash"
	case SLASH:
		return "slash"
	case STAR:
		return "star"
	case PERCENT:
		return "percent"
	case LET:
		return "let"
	case CONST:
		return "const"
	case CLASS:
		return "class"
	case NEW:
		return "new"
	case IMPORT:
		return "import"
	case FROM:
		return "from"
	case FUNCTION:
		return "function"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case FOREACH:
		return "foreach"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case EXPORT:
		return "export"
	case IN:
		return "in"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}
