module "module"
[ident] "choice-x"
{ "{"
container "container"
[ident] "x"
{ "{"
choice "choice"
[ident] "how"
{ "{"
config "config"
false "false"
; ";"
case "case"
[ident] "one"
{ "{"
leaf "leaf"
[ident] "y"
{ "{"
type "type"
[ident] "uint16"
; ";"
} "}"
} "}"
case "case"
[ident] "two"
{ "{"
leaf "leaf"
[ident] "z"
{ "{"
type "type"
[ident] "string"
; ";"
} "}"
} "}"
} "}"
} "}"
} "}"
