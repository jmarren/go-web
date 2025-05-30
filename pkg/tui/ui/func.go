package ui

import "fmt"

func GetById[T EasyPrimitive](node EasyPrimitive, name ID) (*T, error) {
	found := getById(node, name)
	if found == nil {
		return nil, fmt.Errorf("node '%s' not found", name.String())
	}
	if textArea, ok := found.(T); !ok {
		return nil, fmt.Errorf("node '%s' is not of the specified type", name.String())
	} else {
		return &textArea, nil
	}
}

func getById(node EasyPrimitive, name ID) EasyPrimitive {
	if node == nil {
		return nil
	}
	if node.Id() == name {
		return node
	}

	for _, child := range node.Children() {

		childItem := getById(child, name)
		if childItem != nil {
			return childItem
		}
	}
	return nil
}
