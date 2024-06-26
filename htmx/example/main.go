// Just some basic example usage of the library
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ryanwclark1/ui2/htmx"
)

func main() {
	fmt.Println(htmx.SwapBeforeEnd.
		Scroll(htmx.Bottom).
		SettleAfter(time.Millisecond * 500),
	)
	r := htmx.NewResponse().
		Reswap(htmx.SwapAfterBegin.Scroll(htmx.Top)).
		AddTrigger(
			htmx.TriggerObject("hello", "world"),
			htmx.TriggerObject("myEvent", map[string]string{
				"level":   "info",
				"message": "Here is a Message",
			}),
		)

	fmt.Println(r)
	fmt.Println(r.Headers())
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	if !htmx.IsHTMX(r) {
		w.Write([]byte("only HTMX requests allowed"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writer := htmx.NewResponse().
		Reswap(htmx.SwapBeforeBegin).
		Redirect("/cats").
		LocationWithContext("/hello", htmx.LocationContext{
			Target: "#testdiv",
			Source: "HELLO",
		}).
		Refresh(false)

	writer.Write(w)
}
