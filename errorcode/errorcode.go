package errorcode

type ErrorCode int

func (code ErrorCode) String() string {
	return ErrorNames[code]
}

const (
	E_OK         ErrorCode = 10 /* No error */
	E_EOF        ErrorCode = 11 /* End Of File */
	E_INTR       ErrorCode = 12 /* Interrupted */
	E_TOKEN      ErrorCode = 13 /* Bad token */
	E_SYNTAX     ErrorCode = 14 /* Syntax error */
	E_NOMEM      ErrorCode = 15 /* Ran out of memory */
	E_DONE       ErrorCode = 16 /* Parsing complete */
	E_ERROR      ErrorCode = 17 /* Execution error */
	E_TABSPACE   ErrorCode = 18 /* Inconsistent mixing of tabs and spaces */
	E_OVERFLOW   ErrorCode = 19 /* Node had too many children */
	E_TOODEEP    ErrorCode = 20 /* Too many indentation levels */
	E_DEDENT     ErrorCode = 21 /* No matching outer block for dedent */
	E_DECODE     ErrorCode = 22 /* Error in decoding into Unicode */
	E_EOFS       ErrorCode = 23 /* EOF in triple-quoted string */
	E_EOLS       ErrorCode = 24 /* EOL in single-quoted string */
	E_LINECONT   ErrorCode = 25 /* Unexpected characters after a line continuation */
	E_IDENTIFIER ErrorCode = 26 /* Invalid characters in identifier */
	E_BADSINGLE  ErrorCode = 27 /* Ill-formed single statement input */
)
