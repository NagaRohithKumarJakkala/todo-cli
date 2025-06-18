package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Task struct{
	// Id int `json:"id"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
}

type Tasks struct{
	TaskList []Task `json:"taskList"`
}

func showTasks(tasks *Tasks){
	if len(tasks.TaskList) == 0{
		fmt.Println("no tasks found")
		return
	}
	for id ,task:=range tasks.TaskList{
		var marker byte;
		if task.IsCompleted{
			marker = 'x'
		}else{
			marker = ' '
		}
		fmt.Printf("%d. [%c] %s\n", id+1,marker,task.Description)
	}
}

func addTask(description string ,tasks * Tasks){
	newTask := Task{Description: description,IsCompleted: false}	
	tasks.TaskList = append(tasks.TaskList, newTask)

	data,err := json.Marshal(*tasks)
	if err!=nil{
		return
	}
	err = os.WriteFile("tasks.json",data,0644)
	if err!=nil{
		return
	}
	fmt.Println("task added successfully")
}

func deleteTask(id int,tasks * Tasks){
	if id > 0 && id <= len(tasks.TaskList){
		description :=tasks.TaskList[id-1].Description
		tasks.TaskList = slices.Delete(tasks.TaskList,id-1,id)
		data,err := json.Marshal(*tasks)
		if err!=nil{
			return
		}
		err = os.WriteFile("tasks.json",data,0644)
		if err!=nil{
			return
		}
		fmt.Printf("task '%s' deleted successfully\n",description)
	}else{
		fmt.Println("enter valid id")
	}
}

func toggleTask(id int,tasks * Tasks){
	if id > 0 && id <= len(tasks.TaskList){
		tasks.TaskList[id-1].IsCompleted = !tasks.TaskList[id-1].IsCompleted;
		data,err := json.Marshal(*tasks)
		if err!=nil{
			return
		}
		err = os.WriteFile("tasks.json",data,0644)
		if err!=nil{
			return
		}

		var message string

		if(tasks.TaskList[id-1].IsCompleted){
			message = "completed"
		}else{
			message = "not completed"
		}
		fmt.Printf("task '%s' is marked as %s\n",tasks.TaskList[id-1].Description,message)
	}else{
		fmt.Println("enter valid id")
	}
}

func help(){
	fmt.Println("The following are valid commands:")
	fmt.Println("list or l - to show the list")
	fmt.Println("add <text> or a <text> - to add a task")
	fmt.Println("delete <id> or a <id> - to delete a task")
	fmt.Println("toggle <id> or t <id> - to toggle a task")
	fmt.Println("quit or q - to exit the program")
}

func main(){
	_,err := os.Stat("tasks.json")
	if err !=nil{
		data := []byte("{\"taskList\":[]}")
		err = os.WriteFile("tasks.json",data,0644)
		if err!=nil{
			fmt.Println("error: ",err)
		}
	}

	data, err := os.ReadFile("tasks.json")

	if err != nil {
		fmt.Println("Error: ",err)
		return
	}


	var tasks Tasks

	err = json.Unmarshal(data,&tasks)
	if err != nil {
		fmt.Println("Error: ",err)
		return
		
	}

	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		option := strings.Fields(input)
		if(len(option)==0){

		}else if (option[0] == "list" || option[0]=="l")&& len(option)==1{

			showTasks(&tasks)

		} else if (option[0] == "quit" || option[0]=="q") && len(option)==1{

			break

		} else if option[0] == "help" || option[0]=="h"{

			help()

		}else if (option[0] == "add" || option[0]=="a") && len(option) != 1{

			description:=option[1]
			for i:=2; i<len(option); i++{
				description+=" "+option[i]
			}
			addTask(description,&tasks)

		}else if (option[0] == "delete" || option[0]=="d")&& len(option)== 2{

			id,err:= strconv.Atoi(option[1])
			if err!=nil{
				fmt.Println("enter integer as id")
			}
			deleteTask(id,&tasks)

		}else if (option[0]=="toggle" || option[0]=="t") && len(option) ==2{

			id,err:= strconv.Atoi(option[1])
			if err!=nil{
				fmt.Println("enter integer as id")
			}
			toggleTask(id,&tasks)

		}else{

			fmt.Println("invalid command")
			fmt.Println("type h or help for help")

		}
	}
}
