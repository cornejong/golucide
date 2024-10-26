package golucide

import (
	"fmt"
)

var DefaultIconAttributes = iconAttributes{
	Width:          24,
	Height:         24,
	Fill:           "none",
	Stroke:         "currentColor",
	StrokeWidth:    2,
	StrokeLineCap:  "round",
	StrokeLineJoin: "round",
}

type iconAttributes struct {
	Width          int
	Height         int
	Fill           string
	Stroke         string
	StrokeWidth    int
	StrokeLineCap  string
	StrokeLineJoin string
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

func GetIcon(iconName string, opts ...iconOption) string {
	template, exists := icons[iconName]
	if !exists {
		return ""
	}

	attrs := DefaultIconAttributes
	for _, opt := range opts {
		opt(&attrs)
	}

	return fmt.Sprintf(template,
		attrs.Width,
		attrs.Height,
		attrs.Fill,
		attrs.Stroke,
		attrs.StrokeWidth,
		attrs.StrokeLineCap,
		attrs.StrokeLineJoin,
	)
}

func GetAvailableIconNames() []string {
	names := []string{}
	for name := range icons {
		names = append(names, name)
	}

	return names
}
