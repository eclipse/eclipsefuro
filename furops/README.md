# furops
 
`furops` is a pattern scaffolding utility, which gives the ability to create templates for patterns with the use of 
the furo specs.

## Writing a template
Take a look at the sample as a starting point.

## variables:

```yaml
variables:
  - varname: Name
    default: null
    prompt: "Enter the component name (xxx-xxx)"
    inputKind: "string"
    regexp: "^[^\\s-]+-[^\\s-]+$"
    regexpText: "Component name must contain at least one dash (-)"
```

### varname
This defines the name of the variable. 

This name can be used in any subsequent expression.

```yaml
  default: "ToCamel(Name)"
```

### [default]
The default value is a optional expression.
If you just need a string or number, write it as following:
```yaml 
    inputKind: "directory"
    default: "'./output'"
```    
    
```yaml     
    inputKind: "number"
    default: "'42'"
```

### prompt
This defines the question, which is promted to the user.


### inputKind
Define the input type. 

**string** will treat your input as a string.

**number** will treat your input as a number (float 64).

**directory** will give you a autocompletion assistant for a path. "./" is used as a starting point, if no default expression was given.

**type** will give you a autocompletion assistant for all types, which are defined in the specified spec export file.

**service** will give you a autocompletion assistant for all services, which are defined in the specified spec export file.

### regexp / regexpText
Enforce a input form with a regular expression. 

The expression
> Negative lookahead isn't supported for technical reasons, specifically because it conflicts with the O(n)-time guarantees of the library.

If the expression does not compile, or produces a error, `furops` will exit and log the expression.

The user is only then allowed to continue, if the expression matches
```yaml
    regexp: "^[^\\s-]+-[^\\s-]+$"
    regexpText: "Component name must contain at least one dash (-)"
```

### expression
With expression you can define a "calculated variable".

Given the following example, the variable `Small` is true if NumOfRows is smaller then 5.
```yaml

  - varname: NumOfRows
    default: "'4'"
    prompt: "How many rows you need?"
    inputKind: "number"

  - varname: Small
    expression: "NumOfRows < 5" #

```

## The structure section
In this section you define which target file is produced with wich template.
All templates are receiving the values, defined in the variables section.

```yaml
structure:
  - target: "'./output/view-' + Name + '.js'"
    template: "templatefile.tpl"
    notes: "This field is informative only"
```

### target
The target is a expression like default or expression in the variables definition.
You can construct a target with the variables and given functions.

### template
This field defines which template is to be used for generating the target.

## Expressions
We follow the philosophy that the template only has to render and everything that is needed for this is resolved via the expressions.

The expressions allow c++ style syntax, which is much easier than writing in the golang template language.

Besides all known comparison operators, some functions are also provided. If you are missing a method, feel free to make a feature request.

We are using [Knetic/govaluate](https://github.com/Knetic/govaluate) under the hood. 
```go
s = "AnyKind of_string"
t = "furo.fat.String"
srv = "furo.fat.String"
```

### Built in functions

| Function                    | Result                          |
|-----------------------------|---------------------------------|
| `Strlen(s)`                 | 17                              |
| `GetService(srv)`           | struct with spec of the service |
| `GetType(t)`                | struct with spec of the type    |
| `ToSnake(s)`                | `any_kind_of_string`            |
| `ToSnakeWithIgnore(s, '.')` | `any_kind.of_string`            |
| `ToScreamingSnake(s)`       | `ANY_KIND_OF_STRING`            |
| `ToKebab(s)`                | `any-kind-of-string`            |
| `ToScreamingKebab(s)`       | `ANY-KIND-OF-STRING`            |
| `ToCamel(s)`                | `AnyKindOfString`               |
| `ToLowerCamel(s)`           | `anyKindOfString`               |


Functions cannot be passed as parameters, they must be known at the time when the expression is parsed, and are unchangeable after parsing.

### Accessors

 Assuming `foo` has a field called "Length":

	"foo.Length > 9000"

Accessors can be nested to any depth, like the following

	"foo.Bar.Baz.SomeVar"


###  Operators

* Modifiers: `+` `-` `/` `*` `&` `|` `^` `**` `%` `>>` `<<`
* Comparators: `>` `>=` `<` `<=` `==` `!=` `=~` `!~`
* Logical ops: `||` `&&`
* Numeric constants, as 64-bit floating point (`12345.678`)
* String constants (single quotes: `'foobar'`)
* Date constants (single quotes, using any permutation of RFC3339, ISO8601, ruby date, or unix date; date parsing is automatically tried with any string constant)
* Boolean constants: `true` `false`
* Parenthesis to control order of evaluation `(` `)`
* Arrays (anything separated by `,` within parenthesis: `(1, 2, 'foo')`)
* Prefixes: `!` `-` `~`
* Ternary conditional: `?` `:`
* Null coalescence: `??`

See [MANUAL.md](https://github.com/Knetic/govaluate/blob/master/MANUAL.md) for exacting details on what types each operator supports.
