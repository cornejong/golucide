# GoLucide: Lucide Icons for Go
**GoLucide** is [lucide](https://lucide.dev/) icons for go. Currently based on ```v0.453.0```.

[![Go Reference](https://pkg.go.dev/badge/github.com/cornejong/golucide.svg)](https://pkg.go.dev/github.com/cornejong/golucide) 
[![Go Report Card](https://goreportcard.com/badge/github.com/cornejong/golucide)](https://goreportcard.com/report/github.com/cornejong/golucide)

# Installation
To install GoLucide, use go get:

```bash
go get github.com/cornejong/golucide
```

# Usage
GoLucide offers a straightforward API with default icon attributes, customizable via options. Here is a quick start guide:

### Basic Usage
Access an icon directly by it's function:

```go
package main

import (
	"fmt"
	"github.com/cornejong/golucide"
)

func main() {
    fmt.Println(golucide.Star())
}
```
Or use GetIcon to generate an icon's SVG string:
```go
package main

import (
	"fmt"
	"github.com/cornejong/golucide"
)

func main() {
	fmt.Println(golucide.GetIcon("Star"))
}
```

### Customizing Icon Attributes
To change the icon’s size, color, or stroke properties, pass options to GetIcon:

```go
icon := golucide.Star(
    golucide.Size(32), 
	golucide.StrokeWidth(3), 
	golucide.Fill("#FFD700"), 
	golucide.StrokeColor("black"),
)

icon = golucide.GetIcon("Star", 
    golucide.Size(32), 
    golucide.StrokeWidth(3), 
    golucide.Fill("#FFD700"), 
    golucide.StrokeColor("black"),
)
```

### Available Options
- ```Size(size int)```: Set the width and height to the same value.
- ```SizeHW(width, height int)```: Set width and height individually.
- ```StrokeWidth(strokeWidth int)```: Set the width of the stroke.
- ```Fill(fill string)```: Set the fill color.
- ```StrokeColor(stroke string)```: Set the stroke color.
- ```StrokeLineCap(lineCap string)```: Set the stroke line cap.
- ```StrokeLineJoin(lineJoin string)```: Set the stroke line join.
- ```WithClasses(class...string)```: Set the classes for the icon.
- ```WithID(id string)```: Set the ID for the icon.
- ```WithAttribute(name string, value string)```: Set any attribute on the icon.

### Default Attributes
Icons have default attributes defined as:

```go
var DefaultIconAttributes = iconAttributes{
	Width:          24,
	Height:         24,
	Fill:           "none",
	Stroke:         "currentColor",
	StrokeWidth:    2,
	StrokeLineCap:  "round",
	StrokeLineJoin: "round",
    Class:          "lucide",
    Extra          []KeyValuePair
}
```

### Get all available icon names
```go
names := golucide.GetAvailableIconNames()
```

## Contributing
Contributions are welcome! If you'd like to add new features or report issues, feel free to open an issue or pull request on the GitHub repository.

## License
GoLucide is licensed under the MIT License.