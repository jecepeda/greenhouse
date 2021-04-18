package runtime

import "github.com/spf13/viper"

type RuntimeVars struct {
	JWTSeedKey string
}

// Vars are the environment variables that can be used anywhere without the
// need of calling viper or other environment variable manager
var Vars RuntimeVars

func init() {
	Vars = RuntimeVars{
		JWTSeedKey: viper.GetString("jwt_seed_key"),
	}
}
