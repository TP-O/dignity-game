package http

import (
	"chat-server/configs"
	"chat-server/internal/apps/apis"
	"chat-server/internal/adapters/server/gingonic"
	"chat-server/internal/adapters/databases/sqlc"
)

func Cmd(conf *configs.Configs) {
	db := sqlc.NewDatabase(conf)
	application := apis.NewApplication(db)
	server := gingonic.NewAdapter(conf, application)
	
	server.Run()
}