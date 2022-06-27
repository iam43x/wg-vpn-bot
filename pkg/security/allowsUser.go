package security

import (
	"golang.org/x/exp/slices"
	"iam43x.vpn.bot/pkg/config"
)

func AllowUser(name string) bool {
	return slices.Contains(config.Config.Users.Allows, name)
}