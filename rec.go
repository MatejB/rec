package rec

import (
	"fmt"
	"strings"
	"time"
)

func New(name string) *Point {
	return &Point{Name: name, Data: make(map[string]string)}
}

type Point struct {
	Name     string
	Data     map[string]string
	Children []*Point
	Start    time.Time
	End      time.Time
}

func (p *Point) PassChild(name string) *Point {
	c := &Point{Name: name, Data: make(map[string]string)}
	p.Children = append(p.Children, c)
	return c
}

func (p *Point) Record() func() {
	p.Start = time.Now()
	return func() {
		p.End = time.Now()
	}
}

func (p *Point) String() string {
	s := fmt.Sprintf("%s [ %v ]\n", p.Name, p.End.Sub(p.Start))
	// s += fmt.Sprintf("data %q\n", p.Data)
	for _, c := range p.Children {
		s += strings.Replace(strings.Trim(c.String(), " "), "\n", "\n  ", -1)
	}
	return s
}
