package wasm

import (
	"syscall/js"
)

// IndexeDB es un objeto que permite realizar consultas en la API de navegador web del mismo nombre.
// Para más información revisar la documentación sobre indexeDB en:
//  https://developer.mozilla.org/es/docs/Web/API/IndexedDB_API/Usando_IndexedDB.
type IndexeDB struct {
	Name string
}

// IndexConsult realiza la consulta a una indexeDB y permite manipular el objeto encontrado. Esta función
//	sigue el mismo esquema que consulta realizada en javascript vanilla, por lo tanto es necesario
// 	ingresar el nombre del objectStore y el índice para realizar la consulta. Así mismo, para poder manipular
//	los datos obtenidos, es necesario pasar una función a la medida.
func (iDB IndexeDB) IndexConsult(objectStore, index string, onsuccessFunction func(this js.Value)) {
	db := js.Global().Get("window").Get("indexedDB").Call("open", iDB.Name, 1)

	db.Set("onsuccess", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		DB := db.Get("result")
		data := DB.Call("transaction", iDB.Name, "readonly")
		objSt := data.Call("objectStore", objectStore)
		request := objSt.Call("get", index)

		request.Set("onsuccess", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			onsuccessFunction(request)

			return nil
		}))

		return nil
	}))

}

// CustomConsult realiza la consulta a una indexeDB y permite manipular el objeto encontrado. La búsqueda
// se realiza en el índice que coincida con el valor del parámetro index y el valor del parámetro indexValue.
// El parámetro onsuccessFunction permite manipular el objeto encontrado.
func (iDB IndexeDB) CustomConsult(objectStore, index, indexValue string, onsuccessFunction func(this js.Value)) {
	db := js.Global().Get("window").Get("indexedDB").Call("open", iDB.Name, 1)

	db.Set("onsuccess", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		DB := db.Get("result")
		data := DB.Call("transaction", iDB.Name, "readonly")
		objSt := data.Call("objectStore", objectStore)
		indexDB := objSt.Call("index", index)
		request := indexDB.Call("get", indexValue)

		request.Set("onsuccess", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			onsuccessFunction(request)

			return nil
		}))

		return nil
	}))
}

// GetItems permite manipular los objetos dentro de una indexeDB. Sigue el mismo esquema que una consulta
// con js, por lo que es necesario indicar el objectStore. El parámetro onsuccessFunction recibe como argumento
// el cursor resultante (el resultado del evento onsuccess), sobre el mismo se puede acceder a cada uno de
// los objetos guardados en la indexeDB.
func (iDB IndexeDB) GetItems(objectStore string, onsuccessFunction func(e js.Value)) {
	db := js.Global().Get("window").Get("indexedDB").Call("open", iDB.Name, 1)
	db.Set("onsuccess", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		DB := db.Get("result")
		objSt := DB.Call("transaction", iDB.Name).Call("objectStore", objectStore)
		objSt.Call("openCursor").Set("onsuccess", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			cursor := args[0].Get("target").Get("result") //<-- args[0] Es el evento
			onsuccessFunction(cursor)
			return nil
		}))

		return nil
	}))
}
