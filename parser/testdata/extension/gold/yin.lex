module "module"
[ident] "x"
{ "{"
prefix "prefix"
[string] "\"x\""
; ";"
extension "extension"
[ident] "ext"
{ "{"
description "descriptio"...
[string] "\"x\""
; ";"
argument "argument"
[string] "f"
{ "{"
yin-element "yin-elemen"...
true "true"
; ";"
} "}"
} "}"
[extension] "x:ext"
[string] "bar"
{ "{"
container "container"
[ident] "c"
{ "{"
leaf "leaf"
[ident] "l"
{ "{"
type "type"
[ident] "string"
; ";"
} "}"
} "}"
} "}"
} "}"
