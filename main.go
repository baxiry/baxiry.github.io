package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}
		go func() {
			defer conn.Close()

			var (
				state  = ws.StateServerSide
				reader = wsutil.NewReader(conn, state)
				writer = wsutil.NewWriter(conn, state, ws.OpText)
			)
			for {
				header, err := reader.NextFrame()
				if err != nil {
					// handle error
					fmt.Println("\n\t", "Next frame error", err)
					//return

				}
				fmt.Println(header)

				// Reset writer to write frame with right operation code.
				writer.Reset(conn, state, header.OpCode)

				if _, err = io.Copy(writer, reader); err != nil {
					// handle error
					fmt.Println("\n\t", "Copy error", err)
					return
				}
				if err = writer.Flush(); err != nil {
					// handle error
					fmt.Println("\n\t", "Flush error:", err)
					return
				}
			}
		}()
	}))
}
