package wasm

import "syscall/js"

// Devuelve un js.Value que corresponde al objeto HTML con el id señalado.
func GetElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

// Devuelve el valor de un input en un formulario HTML.
func InputValue(id string) string {
	return js.Global().Get("document").Call("getElementById", id).Get("value").String()
}

// Lanza una alerta del navegador
func Alert(msj string) {
	js.Global().Call("alert", msj)
}

// Selecciona un DOMObject y le inserta un contenido.
//
// Example. La función
//
//  InnerHTML("#my_div", "<p>Hello Word</p>")
//
// Es equivalente al siguiente script js:
//
//  document.querySelector('#my_div').innerHTML = '<p>Hello Word</p>'
//
func InnerHTML(selector, content string) {
	js.Global().Get("document").Call("querySelector", selector).Set("innerHTML", content)
}
