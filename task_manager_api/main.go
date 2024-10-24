package main

import "Go_Rest_API/task_manager_api/router"




func main(){
	r := router.SetupRouter()
	r.Run(":8080")

}

