package data

import "github.com/google/wire"

var WireSet = wire.NewSet(NewUserRepo)
