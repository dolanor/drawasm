package main

import (
	"fmt"
	"syscall/js"
)

type App struct {
	isPainting bool
	x          float64
	y          float64
	ctx        js.Value
	color      string

	doc      js.Value
	canvasEl js.Value
}

func (a *App) startPaint(this js.Value, args []js.Value) interface{} {
	e := args[0]
	a.isPainting = true

	a.x = e.Get("pageX").Float() - a.canvasEl.Get("offsetLeft").Float()
	a.y = e.Get("pageY").Float() - a.canvasEl.Get("offsetTop").Float()
	return nil
}
func (a *App) paint(this js.Value, args []js.Value) interface{} {
	if !a.isPainting {
		return nil
	}

	e := args[0]
	nx := e.Get("pageX").Float() - a.canvasEl.Get("offsetLeft").Float()
	ny := e.Get("pageY").Float() - a.canvasEl.Get("offsetTop").Float()

	a.ctx.Set("strokeStyle", a.color)
	a.ctx.Set("lineJoin", "round")
	a.ctx.Set("lineWidth", "5")

	a.ctx.Call("beginPath")
	a.ctx.Call("moveTo", nx, ny)
	a.ctx.Call("lineTo", a.x, a.y)
	a.ctx.Call("closePath")

	a.ctx.Call("stroke")

	a.x = nx
	a.y = ny

	return nil
}

func (a *App) exit(this js.Value, args []js.Value) interface{} {
	a.isPainting = false
	return nil
}

func (a *App) buildColorPalette() {
	colorsEl := a.doc.Call("getElementById", "colors")
	colors := [...]string{"#F4908E", "#F2F097", "#88B0DC", "#F7B5D1", "#53C4AF", "#FDE38C"}

	for _, c := range colors {
		node := a.doc.Call("createElement", "div")
		node.Call("setAttribute", "class", "color")
		node.Call("setAttribute", "id", c)
		node.Call("setAttribute", "style", fmt.Sprintf("background-color: %s", c))

		setColor := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			e := args[0]
			a.color = e.Get("target").Get("id").String()
			return nil
		})

		node.Call("addEventListener", "click", setColor)
		colorsEl.Call("appendChild", node)
	}
}

func (a *App) setupMouseListeners() {
	a.canvasEl.Call("addEventListener", "mousedown", js.FuncOf(a.startPaint))
	a.canvasEl.Call("addEventListener", "mousemove", js.FuncOf(a.paint))
	a.canvasEl.Call("addEventListener", "mouseup", js.FuncOf(a.exit))
}

func NewApp(doc, canvasEl js.Value) *App {
	app := App{
		doc:      doc,
		ctx:      canvasEl.Call("getContext", "2d"),
		canvasEl: canvasEl,
	}
	app.setupMouseListeners()
	app.buildColorPalette()

	return &app
}

func NewCanvasEl(doc js.Value, id string, width, height float64) js.Value {
	canvasEl := doc.Call("getElementById", "canvas")
	canvasEl.Set("width", width)
	canvasEl.Set("height", height)

	return canvasEl
}

func (a *App) Run() {
	stop := make(chan struct{})
	<-stop
}

func main() {
	doc := js.Global().Get("document")
	bodyW := doc.Get("body").Get("clientWidth").Float()
	bodyH := doc.Get("body").Get("clientHeight").Float()

	canvasEl := NewCanvasEl(doc, "canvas", bodyW, bodyH)

	app := NewApp(doc, canvasEl)

	println("Hello Wasm")
	app.Run()
}
