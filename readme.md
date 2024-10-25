# GoLucide: Lucide Icons for Go
**GoLucide** is [lucide](https://lucide.dev/) icons for go. Currently based on ```v0.453.0```.

# Installation
To install GoLucide, use go get:

```bash
go get github.com/cornejong/golucide
```

# Usage
GoLucide offers a straightforward API with default icon attributes, customizable via options. Here is a quick start guide:

### Basic Usage
Use GetIcon to generate an icon's SVG string:

```go
package main

import (
	"fmt"
	"github.com/cornejong/golucide"
)

func main() {
	fmt.Println(golucide.GetIcon("star"))
}
```

### Customizing Icon Attributes
To change the icon’s size, color, or stroke properties, pass options to GetIcon:

```go
icon := golucide.GetIcon("star", 
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

### Default Attributes
Icons have default attributes defined as:

```go
var DefaultAttributes = Attributes{
	Width:          24,
	Height:         24,
	Fill:           "none",
	Stroke:         "currentColor",
	StrokeWidth:    2,
	StrokeLineCap:  "round",
	StrokeLineJoin: "round",
}
```
### Example

```go
icon := golucide.GetIcon("star", 
	golucide.Size(40), 
	golucide.StrokeColor("#FF5733"), 
	golucide.Fill("#FFF"),
)

fmt.Println(icon)
```

## Contributing
Contributions are welcome! If you'd like to add new features or report issues, feel free to open an issue or pull request on the GitHub repository.

## License
GoLucide is licensed under the MIT License.