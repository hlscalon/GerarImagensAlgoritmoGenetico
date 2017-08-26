package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

type Circle struct {
	x, y, r, red, green, blue uint8
}

type Rectangle struct {
	x, y, w, h, red, green, blue uint8
}

type Triangle struct {
	p1, p2, p3, red, green, blue uint8
}

// preparado para 0 até 255
type Data struct {
	circles []Circle
	rectangles []Rectangle
	triangles []Triangle
	nCircles, nRectangles, nTriangles int
}

func NewData(circles []Circle, rectangles []Rectangle, triangles []Triangle, nCircles, nRectangles, nTriangles int) *Data {
	return &Data{circles, rectangles, triangles, nCircles, nRectangles, nTriangles}
}

func (d *Data) toString() string {
	hex := ""

	for _, c := range d.circles {
		hex += fmt.Sprintf("%02x%02x%02x%02x%02x%02x", c.x, c.y, c.r, c.red, c.green, c.blue)
	}

	for _, r := range d.rectangles {
		hex += fmt.Sprintf("%02x%02x%02x%02x%02x%02x%02x", r.x, r.y, r.w, r.h, r.red, r.green, r.blue)
	}

	for _, t := range d.triangles {
		hex += fmt.Sprintf("%02x%02x%02x%02x%02x%02x", t.p1, t.p2, t.p3, t.red, t.green, t.blue)
	}

	return hex
}

func (d *Data) hexToUint(hex string) uint8 {
	n, _ := strconv.ParseUint(hex, 16, 32)
	return uint8(n)
}

func (d *Data) fromString(data string, nCircles, nRectangles, nTriangles int) {
	d.nCircles, d.nRectangles, d.nTriangles = nCircles, nRectangles, nTriangles
	d.circles = make([]Circle, nCircles)
	base := 0
	for i := 0; i < nCircles; i++ {
		x := d.hexToUint(data[base:base+2])
		base += 2
		y := d.hexToUint(data[base:base+2])
		base += 2
		r := d.hexToUint(data[base:base+2])
		base += 2
		red := d.hexToUint(data[base:base+2])
		base += 2
		green := d.hexToUint(data[base:base+2])
		base += 2
		blue := d.hexToUint(data[base:base+2])
		base += 2

		d.circles[i] = Circle{x, y, r, red, green, blue}
	}

	d.rectangles = make([]Rectangle, nRectangles)
	for i := 0; i < nRectangles; i++ {
		x := d.hexToUint(data[base:base+2])
		base += 2
		y := d.hexToUint(data[base:base+2])
		base += 2
		w := d.hexToUint(data[base:base+2])
		base += 2
		h := d.hexToUint(data[base:base+2])
		base += 2
		red := d.hexToUint(data[base:base+2])
		base += 2
		green := d.hexToUint(data[base:base+2])
		base += 2
		blue := d.hexToUint(data[base:base+2])
		base += 2

		d.rectangles[i] = Rectangle{x, y, w, h, red, green, blue}
	}

	d.triangles = make([]Triangle, nTriangles)
	for i := 0; i < nTriangles; i++ {
		p1 := d.hexToUint(data[base:base+2])
		base += 2
		p2 := d.hexToUint(data[base:base+2])
		base += 2
		p3 := d.hexToUint(data[base:base+2])
		base += 2
		red := d.hexToUint(data[base:base+2])
		base += 2
		green := d.hexToUint(data[base:base+2])
		base += 2
		blue := d.hexToUint(data[base:base+2])
		base += 2

		triangle := Triangle{p1, p2, p3, red, green, blue}
		d.triangles[i] = triangle
	}
}

func CreateData() *Data {
	now := time.Now()
	nCircles := 2 //now.Hour()
	nRectangles := 2 // now.Minute()
	nTriangles := 2 // now.Second()
	rd := rand.New(rand.NewSource(now.UnixNano()))

    // numero de retangulos, baseado na hora
	circles := make([]Circle, nCircles)
    for i := 0; i < nCircles; i++ {
		circle := Circle{uint8(rd.Intn(X)), uint8(rd.Intn(Y)), uint8(rd.Intn(50)), uint8(rd.Intn(255)), uint8(rd.Intn(255)), uint8(rd.Intn(255))}
		circles[i] = circle
	}

    // numero de circulos, baseado nos minutos
	rectangles := make([]Rectangle, nRectangles)
    for i := 0; i < nRectangles; i++ {
		rectangle := Rectangle{uint8(rd.Intn(X)), uint8(rd.Intn(Y)), uint8(rd.Intn(X/5)), uint8(rd.Intn(Y/5)), uint8(rd.Intn(255)), uint8(rd.Intn(255)), uint8(rd.Intn(255))}
		rectangles[i] = rectangle
	}

    // numero de triangulos, baseado nos segundos
	triangles := make([]Triangle, nTriangles)
    for i := 0; i < nTriangles; i++ {
		triangle := Triangle{uint8(rd.Intn(X/10)), uint8(rd.Intn(Y/10)), uint8(rd.Intn(Y/10)), uint8(rd.Intn(255)), uint8(rd.Intn(255)), uint8(rd.Intn(255))}
		triangles[i] = triangle
	}

	return NewData(circles, rectangles, triangles, nCircles, nRectangles, nTriangles)
}

