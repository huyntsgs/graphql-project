package gitlab

import "go.uber.org/fx"

var Module = fx.Provide(
	NewGitlabRepo,
)
