package errorcode

var ErrorNames = [...]string{
	E_OK:         "E_OK",
	E_EOF:        "E_EOF",
	E_INTR:       "E_INTR",
	E_TOKEN:      "E_TOKEN",
	E_SYNTAX:     "E_SYNTAX",
	E_NOMEM:      "E_NOMEM",
	E_DONE:       "E_DONE",
	E_ERROR:      "E_ERROR",
	E_TABSPACE:   "E_TABSPACE",
	E_OVERFLOW:   "E_OVERFLOW",
	E_TOODEEP:    "E_TOODEEP",
	E_DEDENT:     "E_DEDENT",
	E_DECODE:     "E_DECODE",
	E_EOFS:       "E_EOFS",
	E_EOLS:       "E_EOLS",
	E_LINECONT:   "E_LINECONT",
	E_IDENTIFIER: "E_IDENTIFIER",
	E_BADSINGLE:  "E_BADSINGLE",
}
