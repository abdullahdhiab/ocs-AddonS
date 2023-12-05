package utils

import (
	"addons/api"
	_ "expvar"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	recovery "github.com/bakins/net-http-recover"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/justinas/alice"
)

// modified versions of examples found at http://golang.org/pkg/net/rpc/
// https://github.com/bakins/json-rpc-example/blob/master/README.md

func GorillaRPC() {

	srv := rpc.NewServer()
	srv.RegisterCodec(json.NewCodec(), "application/json")
	srv.RegisterService(new(api.AddonS), "")

	chain := alice.New(
		func(h http.Handler) http.Handler {
			return handlers.CombinedLoggingHandler(os.Stdout, h)
		},
		handlers.CompressHandler,
		func(h http.Handler) http.Handler {
			return recovery.Handler(os.Stderr, h, true)
		})

	// httprof and expvar endpoints
	rtr := mux.NewRouter()
	rtr.PathPrefix("/debug/").Handler(chain.Then(http.DefaultServeMux))
	rtr.Handle("/rpc", chain.Then(srv))
	log.Fatal(http.ListenAndServe(":9090", rtr))

}
