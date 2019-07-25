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
}

func main() {
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "canvas")

	bodyW := doc.Get("body").Get("clientWidth").Float()
	bodyH := doc.Get("body").Get("clientHeight").Float()

	canvasEl.Set("width", bodyW)
	canvasEl.Set("height", bodyH)

	app := App{}
	app.ctx = canvasEl.Call("getContext", "2d")

	startPaint := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		app.isPainting = true

		app.x = e.Get("pageX").Float() - canvasEl.Get("offsetLeft").Float()
		app.y = e.Get("pageY").Float() - canvasEl.Get("offsetTop").Float()
		return nil
	})

	canvasEl.Call("addEventListener", "mousedown", startPaint)

	paint := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !app.isPainting {
			return nil
		}

		e := args[0]
		nx := e.Get("pageX").Float() - canvasEl.Get("offsetLeft").Float()
		ny := e.Get("pageY").Float() - canvasEl.Get("offsetTop").Float()

		app.ctx.Set("strokeStyle", app.color)
		app.ctx.Set("lineJoin", "round")
		app.ctx.Set("lineWidth", "5")

		app.ctx.Call("beginPath")
		app.ctx.Call("moveTo", nx, ny)
		app.ctx.Call("lineTo", app.x, app.y)
		app.ctx.Call("closePath")

		app.ctx.Call("stroke")

		app.x = nx
		app.y = ny

		return nil
	})

	canvasEl.Call("addEventListener", "mousemove", paint)

	exit := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		app.isPainting = false
		return nil
	})

	canvasEl.Call("addEventListener", "mouseup", exit)

	colorsEl := doc.Call("getElementById", "colors")
	colors := [...]string{"#F4908E", "#F2F097", "#88B0DC", "#F7B5D1", "#53C4AF", "#FDE38C"}

	for _, c := range colors {
		node := doc.Call("createElement", "div")
		node.Call("setAttribute", "class", "color")
		node.Call("setAttribute", "id", c)
		node.Call("setAttribute", "style", fmt.Sprintf("background-color: %s", c))

		setColor := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			e := args[0]
			app.color = e.Get("target").Get("id").String()
			return nil
		})

		node.Call("addEventListener", "click", setColor)
		colorsEl.Call("appendChild", node)
	}

	println("Hello Wasm")
	stop := make(chan struct{})
	<-stop
}
