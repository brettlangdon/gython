Gython
======

This project is currently a for-fun work in progress.

The main goals of this project are to learn about programming languages by trying to rewrite [CPython][] 3.5.0 in [Go][].

[CPython]:https://github.com/python/cpython
[Go]:https://github.com/golang/go/

## Progress
### Scanner
So far I have a mostly working scanner/tokenizer. The main goal was to be able to generate similar output as running `python3 -m tokenize --exact <script.py>`.
Currently there are a few small differences between the output format, but the tokens being produced are the same.


### Grammar Parser
Next up is going to be writing the parser to be able to validate the source code grammar; which will match the form provided from:
```python
import parser
import pprint
import symbol
import token


def resolve_symbol_names(part):
    if not isinstance(part, list):
        return part

    if not len(part):
        return part

    symbol_id = part[0]
    if symbol_id in symbol.sym_name:
        symbol_name = symbol.sym_name[symbol_id]
        return [symbol_name] + [resolve_symbol_names(p) for p in part[1:]]
    elif symbol_id in token.tok_name:
        token_name = token.tok_name[symbol_id]
        return [token_name] + part[1:]
    return part


def main(filename):
    with open(filename, 'r') as fp:
        contents = fp.read()
    st = parser.suite(contents)
    ast = resolve_symbol_names(st.tolist())
    pprint.pprint(ast)

if __name__ == '__main__':
    import sys
    main(sys.argv[1])
```

```bash
python3 grammar.py <script.py>
```

```bash
$ echo "print('hello world')" > test.py
$ python3 parse.py test.py
['file_input',
 ['stmt',
  ['simple_stmt',
   ['small_stmt',
    ['expr_stmt',
     ['testlist_star_expr',
      ['test',
       ['or_test',
        ['and_test',
         ['not_test',
          ['comparison',
           ['expr',
            ['xor_expr',
             ['and_expr',
              ['shift_expr',
               ['arith_expr',
                ['term',
                 ['factor',
                  ['power',
                   ['atom_expr',
                    ['atom', ['NAME', 'print']],
                    ['trailer',
                     ['LPAR', '('],
                     ['arglist',
                      ['argument',
                       ['test',
                        ['or_test',
                         ['and_test',
                          ['not_test',
                           ['comparison',
                            ['expr',
                             ['xor_expr',
                              ['and_expr',
                               ['shift_expr',
                                ['arith_expr',
                                 ['term',
                                  ['factor',
                                   ['power',
                                    ['atom_expr',
                                     ['atom',
                                      ['STRING',
                                       "'hello world'"]]]]]]]]]]]]]]]]]],
                     ['RPAR', ')']]]]]]]]]]]]]]]]]]],
   ['NEWLINE', '']]],
 ['NEWLINE', ''],
 ['ENDMARKER', '']]
```

### AST Parsing
AST parsing will take the validated source grammar and convert it into a valid AST.

The goal is to get a similar AST output as the following:

```python
import ast


def main(filename):
    with open(filename, 'r') as fp:
        contents = fp.read()
    module = ast.parse(contents)
    print(ast.dump(module))

if __name__ == '__main__':
    import sys
    main(sys.argv[1])
```

```bash
$ echo "print('hello world')" > test.py
$ python3 parser.py test.py
Module(body=[Expr(value=Call(func=Name(id='print', ctx=Load()), args=[Str(s='hello world')], keywords=[]))])
```

### Compiler
The compiler will be up after the parser. The compiler will be responsible for converting the parsed AST into Python bytecode.

### Interpreter
The interpreter will be up after the compiler and will be able to execute on Python bytecode.
