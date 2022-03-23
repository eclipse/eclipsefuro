# furops

## variables:

```yaml
variables:
- varname: Name
  default: null
  prompt: "Enter the component name"
  inputKind: "string"
  regexp: ".{3,}"
  regexpText: "At least 3 characters"
```

### varname
### default
### prompt
### inputKind
### regexp / regexpText


## Supported Input Kinds

### string
### directory
### types
### services


## Expressions


### String Fu
```go
"AnyKind of_string"
```

| Function                    | Result               |
|-----------------------------|----------------------|
| `ToSnake(s)`                | `any_kind_of_string` |
| `ToSnakeWithIgnore(s, '.')` | `any_kind.of_string` |
| `ToScreamingSnake(s)`       | `ANY_KIND_OF_STRING` |
| `ToKebab(s)`                | `any-kind-of-string` |
| `ToScreamingKebab(s)`       | `ANY-KIND-OF-STRING` |
| `ToCamel(s)`                | `AnyKindOfString`    |
| `ToLowerCamel(s)`           | `anyKindOfString`    |
| `Strlen(s)`                 | 17                   |

