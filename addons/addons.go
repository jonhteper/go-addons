// Extra functions for golang
package addons

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

// Indica si se está trabajando sobre Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// Indica si se está trabajando sobre Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

//Indica si se está trabajando sobre WebAssembly
func IsWASM() bool {
	return runtime.GOOS == "js"
}

// Lee un archivo tipo JSON, de la ubicación indicada en el parámetro name, e intenta convertirlo a un struct. ,
// retorna el contenido del archivo ─de encontrarlo─ y un error.
func ReadJSONFile(name string, v interface{}) (content string, err error) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}

	content = string(file)
	err = json.Unmarshal(file, v)
	return
}

// Ejecuta una serie de comandos CLI según el sistema operativo.
func Command(c ...string) error {
	args := c[1:]
	cmd := exec.Command(c[0], args...)
	if IsWindows() {

		cmd = exec.Command("cmd", c...)
	}
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
