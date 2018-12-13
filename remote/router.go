/*
Copyright (C) 2018 KIM KeepInMind GmbH/srl

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package remote

import (
	"net/http"

	"github.com/booster-proj/booster"
	"github.com/booster-proj/booster/store"
	"github.com/gorilla/mux"
)

type Router struct {
	r *mux.Router

	Config booster.Config
	Store  *store.SourceStore
}

func NewRouter() *Router {
	return &Router{r: mux.NewRouter()}
}

// SetupRoutes adds the routes available to the router. Make sure
// to fill the public fields of the Router before calling this
// function, otherwise the handlers will not be able to work
// properly.
func (r *Router) SetupRoutes() {
	router := r.r
	router.HandleFunc("/health", makeHealthCheckHandler(r.Config))
	router.HandleFunc("/sources", makeSourcesHandler(r.Store))
	router.HandleFunc("/sources/{name}/block", makeBlockHandler(r.Store)).Methods("POST", "DELETE")
	router.HandleFunc("/policies", makePoliciesHandler(r.Store))
	router.Use(loggingMiddleware)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.r.ServeHTTP(w, req)
}