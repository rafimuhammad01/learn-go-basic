package main

import "fmt"

func main() {
	database := ToDoArray{}

	//Create
	create("kamu", "cobain ini", &database)
	create("kamu", "cobain itu", &database)
	create("dia", "cobain itu", &database)
	database.Println()

	//Read
	cobaread := read("kamu", database)
	fmt.Println(cobaread)

	//Update
	update("dia", "saya", "cobain", &database)
	database.Println()

	//Delete
	delete("kamu", &database)
	database.Println()
}

type ToDo struct {
	name, description string
}

type ToDoArray struct {
	arr []*ToDo
}

func (todoarray *ToDoArray) Println() {
	s := "{"
	for j, i := range todoarray.arr {
		s += i.name
		s += " "
		s += i.description
		if len(todoarray.arr) > 1 && j != len(todoarray.arr)-1 {
			s += ", "
		}
	}
	s += "}"

	fmt.Println(s)
}

func (todoarray *ToDoArray) getElement(name string) (*ToDo, int) {
	for j, i := range todoarray.arr {
		if i.name == name {
			return i, j
		}
	}

	return &ToDo{}, -1
}

func (todoarray *ToDoArray) addToArray(todo *ToDo) {
	isFind := false
	element, _ := todoarray.getElement(todo.name)
	if element.name != "" {
		isFind = true
	}

	if !isFind {
		todoarray.arr = append(todoarray.arr, todo)
	}

}

// Create
func create(name, description string, database *ToDoArray) {
	newTodo := ToDo{
		name:        name,
		description: description,
	}

	database.addToArray(&newTodo)
}

// Read
func read(name string, database ToDoArray) *ToDo {
	element, _ := database.getElement(name)
	return element
}

//Update
func update(name, newName, newDescription string, database *ToDoArray) {
	element, _ := database.getElement(name)
	if element != (&ToDo{}) {
		element.name = newName
		element.description = newDescription
	}
}

//Delete
func delete(name string, database *ToDoArray) {
	_, index := database.getElement(name)
	if index != -1 {
		database.arr = append(database.arr[:index], database.arr[index+1:]...)
	}
}
