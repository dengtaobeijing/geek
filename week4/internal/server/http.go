package server

import "github.com/google/wire"

var WireSet = wire.NewSet(NewHttpServer0, NewHttpServer1)
