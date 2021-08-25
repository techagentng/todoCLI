
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TodoListCmd represents the TodoList command
var TodoListCmd = &cobra.Command{
	Use:   "TodoList",
	Short: "This represents a short and callable description of todo list application ",
	Long: `A longer description comes in handy here.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TodoList called")
	},
}

func init() {
	rootCmd.AddCommand(TodoListCmd)
}

// TodoApp for base struct of TodoList
type TodoApp struct {
	TodoList map[string]bool
	unDoneTask  []string
	doneTask []string
}

//addTask receiver method, setting up the structure of map with valid params
func (todo *TodoApp) addTodo(todoTask string) *TodoApp{
	//CreateFile()  //Write file function may take below as argument
	todo.TodoList[todoTask] = false
	//writeFile()
	//I am returning a todo struct here
	return todo
}