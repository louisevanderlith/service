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
	//result.AddItem("#", "Comment API", "fa-chat", commentChildren(path))

	return result
}

func clientChlidren(path string) *control.Menu {
	children := control.NewMenu(path)
	children.AddItem("/client/create", "Create New", "fa-user", nil)

	return children
}
