package lexical

// structure
type Token struct {
	token    int
	intValue int
}

// Tokens
const (
	T_PLUS = iota // 0
	T_MINUS       // 1
	T_STAR        // 2
	T_SLASH       // 3
	T_INTLIT      // 4
)

/*
We will start with a language that has only five lexical elements:

	the four basic maths operators: *, /, + and -
	decimal whole numbers which have 1 or more digits 0 .. 9
*/
