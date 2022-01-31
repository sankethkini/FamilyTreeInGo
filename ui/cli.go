package ui

import (
	"fmt"
	"os"

	"github.com/sankethkini/FamilyTreeInGo/application"
)

func Start() {
	app := application.NewApp()
	for {
		MenuForUser(app)
	}
}

func MenuForUser(app *application.MyApp) {
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
	selectAPI(option, app)
}

func selectAPI(option int, app *application.MyApp) {
	switch option {
	case 1:
		addNewNode(app)

	case 2:
		addNewDependency(app)

	case 3:
		deleteNode(app)

	case 4:
		deleteDependency(app)

	case 5:
		getParents(app)

	case 6:
		getChildren(app)

	case 7:
		getAncestors(app)

	case 8:
		getDescendants(app)

	case 9:
		os.Exit(1)

	default:
		fmt.Println("enter the correct option")
	}
}

func displayMessage(msg ...map[string]interface{}) {
	for _, val := range msg {
		for k, v := range val {
			fmt.Printf("%v---------->%v\n", k, v)
		}
	}
}

func nodeIDInput() (id string) {
	fmt.Println("enter the node id")
	fmt.Scanf("%s", &id)
	return
}

func dependencyInput() (parentID string, childID string) {
	fmt.Println("Enter the parent id")
	fmt.Scanf("%s", &parentID)
	fmt.Println("Enter the child id")
	fmt.Scanf("%s", &childID)
	return
}

func addNewNode(app *application.MyApp) {
	var name string
	id := nodeIDInput()
	fmt.Println("enter the name")
	fmt.Scanf("%s", &name)
	fmt.Println(name)
	msg, err := app.AddNode(name, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	displayMessage(msg...)
}

func addNewDependency(app *application.MyApp) {
	parentID, childID := dependencyInput()
	msg, err := app.AddDependency(parentID, childID)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func deleteNode(app *application.MyApp) {
	id := nodeIDInput()
	msg, err := app.DeleteNode(id)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func deleteDependency(app *application.MyApp) {
	parentID, childID := dependencyInput()
	msg, err := app.DeleteDependency(parentID, childID)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func getParents(app *application.MyApp) {
	id := nodeIDInput()
	msg, err := app.Parents(id)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func getChildren(app *application.MyApp) {
	id := nodeIDInput()
	msg, err := app.Children(id)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func getAncestors(app *application.MyApp) {
	id := nodeIDInput()
	msg, err := app.Ancestors(id)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}

func getDescendants(app *application.MyApp) {
	id := nodeIDInput()
	msg, err := app.Descendants(id)
	if err != nil {
		fmt.Println(err)
	}
	displayMessage(msg...)
}
