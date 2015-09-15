Gython
======

This project is currently a for-fun work in progress.

The main goals of this project are to learn about programming languages by trying to rewrite [CPython][] 3.5.0 in [Go][].

CPython:https://github.com/python/cpython
Go:https://github.com/golang/go/

## Progress
### Scanner
So far I have a mostly working scanner/tokenizer. The main goal was to be able to generate similar output as running `python3 -m tokenize --exact <script.py>`.
Currently there are a few small differences between the output format, but the tokens being produced are the same.


### Parser
Next up is going to be writing the parser to be able to generate an AST which will match the form provided from:
```python
import ast
import pprint


def main(filename):
    with open(filename, 'r') as fp:
        contents = fp.read()
    root = ast.parse(contents, filename)
    pprint.pprint(ast.dump(root, include_attributes=True))


if __name__ == '__main__':
    import sys
    main(sys.argv[1])
```

```bash
python3 parse.py <script.py>
```

### Compiler
The compiler will be up after the parser. The compiler will be responsible for converting the parsed AST into Python bytecode.

### Interpreter
The interpreter will be up after the compiler and will be able to execute on Python bytecode.
