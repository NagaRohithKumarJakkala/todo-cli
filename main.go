package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"slices"
	"strconv"
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
		tasks.TaskList = slices.Delete(tasks.TaskList,id-1,id)
		data,err := json.Marshal(*tasks)
		if err!=nil{
			return
		}
		err = os.WriteFile("tasks.json",data,0644)
		if err!=nil{
			return
		}
		fmt.Println("task deleted successfully")
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
		fmt.Printf("task %d is marked as %s\n",id,message)
	}else{
		fmt.Println("enter valid id")
	}
}

func main(){
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

		}else if option[0] == "show" || option[0]=="s"{
			showTasks(&tasks)
		} else if option[0] == "quit" {
			break
		} else if option[0] == "help" || option[0]=="h"{

		}else if option[0] == "add" && len(option) != 1{
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
