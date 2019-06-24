package env

import (
	rootenv "github.com/ian-howell/airshipctl/pkg/environment"
)

type Settings struct {
	*rootenv.AirshipCTLSettings

	Message string
}
