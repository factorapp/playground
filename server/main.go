// build !js,wasm
package main

import (
	"log"
	"net/http"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "./app/example.wasm")
}
func main() {
	http.HandleFunc("/app/example.wasm", wasmHandler)
	/*	cwd, err := os.Getcwd()
		if err != nil {
			panic(err)
		}
		app := filepath.Join(cwd, "app")
	*/

	http.HandleFunc("/wasm_exec.js", jsHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
func jsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./app/wasm_exec.js")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./app/index.html")
}
