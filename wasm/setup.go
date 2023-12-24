package wasm

import (
	"fmt"
	"html/template"
)

func makeScript(s string) template.HTML {
	script := `<script>%s</script>`
	return template.HTML(fmt.Sprintf(script, s))
}

func MakeWasmScript(tag, s string) template.HTML {
	t := fmt.Sprintf(wasmScript, tag, s)
	return makeScript(t)
}

var wasmScript = `document.addEventListener("DOMContentLoaded", function() {
            const go = new Go();
  WebAssembly.instantiateStreaming(fetch("/assets/other/json.wasm.gz?id=%s"), go.importObject).then((result) => {
                go.run(result.instance);
                WasmReady('%s');
            });
});`
