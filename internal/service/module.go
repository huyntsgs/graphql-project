package service

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewGitlabService,
	// adding other services below
)
