package logic

import (
	"github.com/louisevanderlith/mango/control"
)

func GetMenu(path string) *control.Menu {
	return getItems(path)
}

func getItems(path string) *control.Menu {
	result := control.NewMenu(path)

	result.AddItem("#", "Clients", "fa-users", clientChlidren(path))
	result.AddItem("#", "Inventory", "fa-box", inventoryChlidren(path))
	result.AddItem("#", "Services", "fa-question", servicesChlidren(path))

	return result
}

func clientChlidren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/client/create", "Create New", "fa-user", nil)

	return children
}

func inventoryChlidren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/parts/view/A10", "View", "fa-glass", nil)
	children.AddItem("/parts/create", "Create New", "fa-box", nil)

	return children
}

func servicesChlidren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/services/view/A10", "View", "fa-glass", nil)
	children.AddItem("/services/create", "Create New", "fa-box", nil)

	return children
}
