package configs

import "github.com/vanbien2402/first-web-demo/internal/pkg/rds"

//Config system configs
type Config struct {
	DB rds.Config
}
