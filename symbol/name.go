package symbol

var SymbolNames = [...]string{
	SINGLE_INPUT:       "SINGLE_INPUT",
	FILE_INPUT:         "FILE_INPUT",
	EVAL_INPUT:         "EVAL_INPUT",
	DECORATOR:          "DECORATOR",
	DECORATORS:         "DECORATORS",
	DECORATED:          "DECORATED",
	ASYNC_FUNCDEF:      "ASYNC_FUNCDEF",
	FUNCDEF:            "FUNCDEF",
	PARAMETERS:         "PARAMETERS",
	TYPEDARGSLIST:      "TYPEDARGSLIST",
	TFPDEF:             "TFPDEF",
	VARARGSLIST:        "VARARGSLIST",
	VFPDEF:             "VFPDEF",
	STMT:               "STMT",
	SIMPLE_STMT:        "SIMPLE_STMT",
	SMALL_STMT:         "SMALL_STMT",
	EXPR_STMT:          "EXPR_STMT",
	TESTLIST_STAR_EXPR: "TESTLIST_STAR_EXPR",
	AUGASSIGN:          "AUGASSIGN",
	DEL_STMT:           "DEL_STMT",
	PASS_STMT:          "PASS_STMT",
	FLOW_STMT:          "FLOW_STMT",
	BREAK_STMT:         "BREAK_STMT",
	CONTINUE_STMT:      "CONTINUE_STMT",
	RETURN_STMT:        "RETURN_STMT",
	YIELD_STMT:         "YIELD_STMT",
	RAISE_STMT:         "RAISE_STMT",
	IMPORT_STMT:        "IMPORT_STMT",
	IMPORT_NAME:        "IMPORT_NAME",
	IMPORT_FROM:        "IMPORT_FROM",
	IMPORT_AS_NAME:     "IMPORT_AS_NAME",
	DOTTED_AS_NAME:     "DOTTED_AS_NAME",
	IMPORT_AS_NAMES:    "IMPORT_AS_NAMES",
	DOTTED_AS_NAMES:    "DOTTED_AS_NAMES",
	DOTTED_NAME:        "DOTTED_NAME",
	GLOBAL_STMT:        "GLOBAL_STMT",
	NONLOCAL_STMT:      "NONLOCAL_STMT",
	ASSERT_STMT:        "ASSERT_STMT",
	COMPOUND_STMT:      "COMPOUND_STMT",
	ASYNC_STMT:         "ASYNC_STMT",
	IF_STMT:            "IF_STMT",
	WHILE_STMT:         "WHILE_STMT",
	FOR_STMT:           "FOR_STMT",
	TRY_STMT:           "TRY_STMT",
	WITH_STMT:          "WITH_STMT",
	WITH_ITEM:          "WITH_ITEM",
	EXCEPT_CLAUSE:      "EXCEPT_CLAUSE",
	SUITE:              "SUITE",
	TEST:               "TEST",
	TEST_NOCOND:        "TEST_NOCOND",
	LAMBDEF:            "LAMBDEF",
	LAMBDEF_NOCOND:     "LAMBDEF_NOCOND",
	OR_TEST:            "OR_TEST",
	AND_TEST:           "AND_TEST",
	NOT_TEST:           "NOT_TEST",
	COMPARISON:         "COMPARISON",
	COMP_OP:            "COMP_OP",
	STAR_EXPR:          "STAR_EXPR",
	EXPR:               "EXPR",
	XOR_EXPR:           "XOR_EXPR",
	AND_EXPR:           "AND_EXPR",
	SHIFT_EXPR:         "SHIFT_EXPR",
	ARITH_EXPR:         "ARITH_EXPR",
	TERM:               "TERM",
	FACTOR:             "FACTOR",
	POWER:              "POWER",
	ATOM_EXPR:          "ATOM_EXPR",
	ATOM:               "ATOM",
	TESTLIST_COMP:      "TESTLIST_COMP",
	TRAILER:            "TRAILER",
	SUBSCRIPTLIST:      "SUBSCRIPTLIST",
	SUBSCRIPT:          "SUBSCRIPT",
	SLICEOP:            "SLICEOP",
	EXPRLIST:           "EXPRLIST",
	TESTLIST:           "TESTLIST",
	DICTORSETMAKER:     "DICTORSETMAKER",
	CLASSDEF:           "CLASSDEF",
	ARGLIST:            "ARGLIST",
	ARGUMENT:           "ARGUMENT",
	COMP_ITER:          "COMP_ITER",
	COMP_FOR:           "COMP_FOR",
	COMP_IF:            "COMP_IF",
	ENCODING_DECL:      "ENCODING_DECL",
	YIELD_EXPR:         "YIELD_EXPR",
	YIELD_ARG:          "YIELD_ARG",
}
