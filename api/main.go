package api

import (
	"server-application/server"
)

func Start(r *server.Router) {
	r.OPTIONS("/{rest:.*}", server.ConstructRequest)
	handleFiles(r.Group("/files"))
	r.PutToAllRoutes(server.ConstructRequest)
}
