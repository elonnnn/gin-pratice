package global

import (
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
)

var (
	GVA_VP                  *viper.Viper
	GVA_Concurrency_Control = &singleflight.Group{}
)
