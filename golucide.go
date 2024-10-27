package golucide

import (
	"fmt"
	"strings"
)

const svgTagTemplate string = `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" fill="%s" stroke="%s" stroke-width="%d" stroke-linecap="%s" stroke-linejoin="%s" class="%s" %s viewBox="0 0 24 24" >%s</svg>`

var DefaultIconAttributes = iconAttributes{
	Width:          24,
	Height:         24,
	Fill:           "none",
	Stroke:         "currentColor",
	StrokeWidth:    2,
	StrokeLineCap:  "round",
	StrokeLineJoin: "round",
	Class:          "lucide",
}

type iconAttributes struct {
	Width          int
	Height         int
	Fill           string
	Stroke         string
	StrokeWidth    int
	StrokeLineCap  string
	StrokeLineJoin string
	Class          string
	Extra          []KeyValuePair
}

type KeyValuePair struct {
	Key   string
	Value string
}

type iconOption func(*iconAttributes)

func Size(size int) iconOption {
	return func(a *iconAttributes) {
		a.Width = size
		a.Height = size
	}
}

func SizeHW(width, height int) iconOption {
	return func(a *iconAttributes) {
		a.Width = width
		a.Height = height
	}
}

func StrokeWidth(strokeWidth int) iconOption {
	return func(a *iconAttributes) {
		a.StrokeWidth = strokeWidth
	}
}

func Fill(fill string) iconOption {
	return func(a *iconAttributes) {
		a.Fill = fill
	}
}

func StrokeColor(stroke string) iconOption {
	return func(a *iconAttributes) {
		a.Stroke = stroke
	}
}

func StrokeLineCap(lineCap string) iconOption {
	return func(a *iconAttributes) {
		a.StrokeLineCap = lineCap
	}
}

func StrokeLineJoin(lineJoin string) iconOption {
	return func(a *iconAttributes) {
		a.StrokeLineJoin = lineJoin
	}
}

func WithClasses(classes ...string) iconOption {
	return func(a *iconAttributes) {
		a.Class = fmt.Sprintf("%s %s", a.Class, strings.Join(classes, " "))
	}
}

func WithID(id string) iconOption {
	return WithAttribute("id", id)
}

func WithAttribute(name string, value string) iconOption {
	return func(a *iconAttributes) {
		a.Extra = append(a.Extra, KeyValuePair{
			Key:   name,
			Value: value,
		})
	}
}

func GetIcon(iconName string, opts ...iconOption) string {
	paths, exists := icons[iconName]
	if !exists {
		return ""
	}

	attrs := DefaultIconAttributes
	for _, opt := range opts {
		opt(&attrs)
	}

	var extraAttributes strings.Builder
	if len(attrs.Extra) > 0 {
		for _, pair := range attrs.Extra {
			extraAttributes.WriteString(fmt.Sprintf(`%s="%s"`, pair.Key, pair.Value))
		}
	}

	return fmt.Sprintf(svgTagTemplate,
		attrs.Width,
		attrs.Height,
		attrs.Fill,
		attrs.Stroke,
		attrs.StrokeWidth,
		attrs.StrokeLineCap,
		attrs.StrokeLineJoin,
		attrs.Class,
		extraAttributes.String(),
		paths,
	)
}

func GetAvailableIconNames() []string {
	names := []string{}
	for name := range icons {
		names = append(names, name)
	}

	return names
}
