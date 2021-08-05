package main

import (
	 ."github.com/ali2210/expert-rotary-phone/redia"
	"github.com/go-redis/redis"
	"log"
	"fmt"
)


func main() {
	redia_client := NewRediaObject()
	client := redis.NewClient(
		&redis.Options{
			Addr : "localhost:6379",
			Password : "",
			DB : 0,
		})
	ping, err := client.Ping().Result(); if err != nil {
		log.Fatalln("Client connection :", err.Error())
		return 
	}
	
	fmt.Println("@ping:", ping)
	tasks := []Tasks{ 
		Tasks{
			Name : "Application Dockerize",
			Id : "1",
			Description : "Redis container active in another container",
			DaysCompleted : 1,
		},
	}

	Num = len(tasks)
	tasks_list, err := redia_client.AddTasks(tasks[0], client); if err != nil {
		log.Fatalln("tasks :", err.Error())
		return 
	}
	

	fmt.Println("@list:", tasks_list)
	dtask, err := redia_client.DisplayTasks(client); if err != nil {
		log.Fatalln("tasks :", err.Error())
		return
	}
	fmt.Println("@get_Data:", dtask)
}