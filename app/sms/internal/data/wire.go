package data

import (
	"akita/quantum_cat/app/sms/internal/data/sms"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(sms.NewUserRepo)
