# Playground
This is a proof of concept for generating WASM applications from mostly HTML.
*REQUIRES Go 1.11+ at TIP (July 15, 2018)*

## Running

```
go get ./...
make run
```
OR
```
docker build -t you/playground .
docker run -p 80:80 you/playground
```

## Directories

* app - contains the HTML and Javascript bootstrap
* assets - page assets (js,images, etc), served as "/assets/*"
* client - the WASM main() that gets run in the browser
* components - reusable HTML components
* layouts - FUTURE/not used yet
* pages - individual pages, indexed by route
* server - Go HTTP server that serves the "app" folder and assets
* cmd/parse - the `parse` command that transforms ".ghtml" files into Go source

## Component Format
```
<div class="blue">
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
}
```

The first part of the `.ghtml` file is the HTML template.
The template is followed by several tags:

* Imports - a list of imports required by Methods
* Parameters - a list of fields that will be added to the component struct
* Methods - (this tag must be LAST in the template) a list of methods declared on the component struct

## Page Format
```
<main role="main" class="container">
    <components:Nav Name="Stuff" />
    <div class="starter-template">
        <h1>Bootstrap starter template</h1>
        <p class="lead">Use this document as a way to quickly start any new project.<br/> All you get is this text and a mostly barebones HTML document.</p>
    </div>
</main>

@Imports {
    "strings"
}

@Route "/"
@Title "Page Title!"

@Parameters {
    Name string
}

@Methods 
    func (p *Index) LowerName() string {
        return strings.ToLower(p.Name)
}
```


The first part of the `.ghtml` file is the HTML template. It may include reference tags to components in the "./components" package
The template is followed by several tags:

* Imports - a list of imports required by Methods
* Route - the URL that this page will respond for 
* Title - the page title
* Parameters - a list of fields that will be added to the page's component struct
* Methods - (this tag must be LAST in the template) a list of methods declared on the page's component struct
