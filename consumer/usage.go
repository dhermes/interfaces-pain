package consumer

import (
	"github.com/dhermes/interfaces-pain/newname"
	"github.com/dhermes/interfaces-pain/oldname"
	"github.com/dhermes/interfaces-pain/withalias"
)

// UpgradeValue upgrades a value from the old package to the new package.
func UpgradeValue(v oldname.Value) newname.Value {
	return v
}

// UpgradeContainer upgrades a container from the old package to the new package.
func UpgradeContainer(c oldname.Container) newname.Container {
	return c
}

// UpgradeValueAlias upgrades a value from the old package to the alias package.
func UpgradeValueAlias(v oldname.Value) withalias.Value {
	return v
}

// UpgradeContainerAlias upgrades a container from the old package to the alias package.
func UpgradeContainerAlias(c oldname.Container) withalias.Container {
	return c
}
