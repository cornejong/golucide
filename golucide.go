package golucide

import (
	"fmt"
)

var DefaultAttributes = Attributes{
	Width:          24,
	Height:         24,
	Fill:           "none",
	Stroke:         "currentColor",
	StrokeWidth:    2,
	StrokeLineCap:  "round",
	StrokeLineJoin: "round",
}

type Attributes struct {
	Width          int
	Height         int
	Fill           string
	Stroke         string
	StrokeWidth    int
	StrokeLineCap  string
	StrokeLineJoin string
}

type Option func(*Attributes)

func Size(size int) Option {
	return func(a *Attributes) {
		a.Width = size
		a.Height = size
	}
}

func SizeHW(width, height int) Option {
	return func(a *Attributes) {
		a.Width = width
		a.Height = height
	}
}

func StrokeWidth(strokeWidth int) Option {
	return func(a *Attributes) {
		a.StrokeWidth = strokeWidth
	}
}

func Fill(fill string) Option {
	return func(a *Attributes) {
		a.Fill = fill
	}
}

func StrokeColor(stroke string) Option {
	return func(a *Attributes) {
		a.Stroke = stroke
	}
}

func StrokeLineCap(lineCap string) Option {
	return func(a *Attributes) {
		a.StrokeLineCap = lineCap
	}
}

func StrokeLineJoin(lineJoin string) Option {
	return func(a *Attributes) {
		a.StrokeLineJoin = lineJoin
	}
}

func GetIcon(iconName string, opts ...Option) string {
	template, exists := icons[iconName]
	if !exists {
		return ""
	}

	attrs := DefaultAttributes
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
