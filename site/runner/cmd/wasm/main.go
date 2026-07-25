//go:build js && wasm

// Command wasm exposes the dedupe runner to JavaScript. Build with:
//
//	GOOS=js GOARCH=wasm go build -o ../../web/dedupe.wasm .
package main

import (
	"encoding/json"
	"syscall/js"

	runner "github.com/gopher-workplace/site/runner"
)

func runDedupe(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return `{"ok":false,"error":"no source provided"}`
	}
	rep := runner.RunDedupe(args[0].String())
	b, err := json.Marshal(rep)
	if err != nil {
		return `{"ok":false,"error":"marshal: ` + err.Error() + `"}`
	}
	return string(b)
}

func runDedupeCustom(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		return `{"ok":false,"error":"need source and input"}`
	}
	res := runner.RunDedupeCustom(args[0].String(), args[1].String())
	b, err := json.Marshal(res)
	if err != nil {
		return `{"ok":false,"error":"marshal: ` + err.Error() + `"}`
	}
	return string(b)
}

func runPlanLimits(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return `{"ok":false,"error":"no source provided"}`
	}
	rep := runner.RunPlanLimits(args[0].String())
	b, err := json.Marshal(rep)
	if err != nil {
		return `{"ok":false,"error":"marshal: ` + err.Error() + `"}`
	}
	return string(b)
}

func runPlanLimitsCustom(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		return `{"ok":false,"error":"need source and input"}`
	}
	res := runner.RunPlanLimitsCustom(args[0].String(), args[1].String())
	b, err := json.Marshal(res)
	if err != nil {
		return `{"ok":false,"error":"marshal: ` + err.Error() + `"}`
	}
	return string(b)
}

func format(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return `{"ok":false,"error":"no source provided"}`
	}
	res := runner.Format(args[0].String())
	b, err := json.Marshal(res)
	if err != nil {
		return `{"ok":false,"error":"marshal: ` + err.Error() + `"}`
	}
	return string(b)
}

func main() {
	js.Global().Set("gopherFormat", js.FuncOf(format))
	js.Global().Set("gopherRunDedupe", js.FuncOf(runDedupe))
	js.Global().Set("gopherRunDedupeCustom", js.FuncOf(runDedupeCustom))
	js.Global().Set("gopherRunPlanLimits", js.FuncOf(runPlanLimits))
	js.Global().Set("gopherRunPlanLimitsCustom", js.FuncOf(runPlanLimitsCustom))
	select {} // keep the Go runtime alive for callbacks
}
