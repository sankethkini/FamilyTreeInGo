package ui

import (
	"fmt"
	"os"

	"github.com/sankethkini/FamilyTreeInGo/application"
)

func MenuForUser() {
	fmt.Println("<--------------------------->")
	fmt.Println("choose the operation")
	fmt.Println("1. Add new node")
	fmt.Println("2. Add new dependency")
	fmt.Println("3. Delete a node")
	fmt.Println("4. Delete a dependency")
	fmt.Println("5. Get parents of a node")
	fmt.Println("6. Get children of a node")
	fmt.Println("7. Get ancestors of a node")
	fmt.Println("8. Get descendents of a node")
	fmt.Println("9. exit")
	fmt.Println("enter your option")
	var option int
	fmt.Scanf("%d", &option)
	selectApi(option)
}

func selectApi(option int) {

	switch option {
	case 1:
		addNewNode()

	case 2:
		addNewDependency()

	case 3:
		deleteNode()

	case 4:
		deleteDependency()

	case 5:
		getParents()

	case 6:
		getChildren()

	case 7:
		getAncestors()

	case 8:
		getDescendants()

	case 9:
		os.Exit(1)

	default:
		MenuForUser()
	}
}

func displayMessage(msg ...map[string]interface{}) {
	for _, val := range msg {
		for k, v := range val {
			fmt.Printf("%v---------->%v\n", k, v)
		}
	}
}

func nodeIdInput() (id string) {
	fmt.Println("enter the node id")
	fmt.Scanf("%s", &id)
	return
}

func dependencyInput() (parentId string, childId string) {
	fmt.Println("Enter the parent id")
	fmt.Scanf("%s", &parentId)
	fmt.Println("Enter the child id")
	fmt.Scanf("%s", &childId)
	return
}

func addNewNode() {
	var name string
	id := nodeIdInput()
	fmt.Println("enter the name")
	fmt.Scanf("%s", &name)
	fmt.Println(name)
	msg, err := application.AddNode(name, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	displayMessage(msg...)
}

func addNewDependency() {
	parentId, childId := dependencyInput()
	msg, err := application.AddDependency(parentId, childId)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func deleteNode() {
	id := nodeIdInput()
	msg := application.DeleteNode(id)
	displayMessage(msg...)
}

func deleteDependency() {
	parentId, childId := dependencyInput()
	msg := application.DeleteDependency(parentId, childId)
	displayMessage(msg...)
}

func getParents() {
	id := nodeIdInput()
	msg, err := application.Parents(id)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func getChildren() {
	id := nodeIdInput()
	msg, err := application.Children(id)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func getAncestors() {
	id := nodeIdInput()
	msg, err := application.Ancestors(id)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func getDescendants() {
	id := nodeIdInput()
	msg, err := application.Descendants(id)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}
