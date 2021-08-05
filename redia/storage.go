package redia

import (
	"github.com/go-redis/redis"	
	"encoding/json"
	"log"
)

var(
	todo []Tasks
	Num int

)

type Redia_Storage interface {
	AddTasks(tasks Tasks, client *redis.Client) ([]Tasks, error) 
	DisplayTasks(client *redis.Client) (string, error)
}


type Tasks struct {
	Name string `json:"name",omitempty`
	Description string `json:"describe", omitempty` 
	DaysCompleted int `json:"completed", omitempty`
	Id string `json:"id", omitempty`
}


func NewRediaObject() Redia_Storage  {
	return &Tasks{}
}

func (task *Tasks) AddTasks(tasks Tasks, client *redis.Client) ([]Tasks, error)  {
	NewTodoList()
	todo = append(todo, tasks)
	data_marshal, err := json.Marshal(todo); if err != nil {
		log.Fatalln("tasks err:", err.Error())
		return nil, err
	}

	for i := range todo {
		err = client.Set(todo[i].Id, data_marshal, 0).Err(); if err != nil {
			log.Fatalln("tasks err:", err.Error())
			return todo, err
		}	
	}
	
	return todo, nil	
}

func (task *Tasks) DisplayTasks(client *redis.Client) (string, error)  {
		for i := range todo{
			value ,err := client.Get(todo[i].Id).Result(); if err != nil {
				log.Fatalln("tasks err:", err.Error())
				return value, err
			}
			return value, nil
		}
	return "", nil
}

func NewTodoList(){
	todo = make([]Tasks, Num)	
}

