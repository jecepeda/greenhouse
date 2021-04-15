package runtime

import "github.com/spf13/viper"

type RuntimeVars struct {
	JWTSeedKey string
}

var Vars RuntimeVars

func init() {
	Vars = RuntimeVars{
		JWTSeedKey: viper.GetString("jwt_seed_key"),
	}
}
