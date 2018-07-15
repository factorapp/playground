# Playground
This is a proof of concept for generating WASM applications from mostly HTML.

## Running

```
go get ./...
make run
```

## Component Format
```
<div class="main">
    <p>Some Content</p>
    <p>{vecty-field:Name}</p>
    <p>{vecty-call:LowerName}</p>
</div>

@Imports {
    "strings"
}


@Parameters {
    Name string
}

@Methods
func (n *Nav) LowerName() string {
    return strings.ToLower(n.Name)
}```

The first part of the `.ghtml` file is the HTML template.
The template is followed by several tags:

* Imports - a list of imports required by Methods
* Parameters - a list of fields that will be added to the component struct
* Methods - (this tag must be LAST in the template) a list of methods declared on the component struct
