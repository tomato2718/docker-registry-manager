package env

import (
	"os"
)

var (
	RegistryUrl = os.Getenv("DRM_REGISTRY_URL")
)
