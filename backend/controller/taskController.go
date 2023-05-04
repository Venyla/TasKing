package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"webservice/service"
)

//var tasks = []data.Task{
//	{TaskId: uuid.MustParse("db9c48a8-f491-4120-a11f-21ec27335c2a"), Title: "Drink Beer", IconUrl: "https://imageresizer.static9.net.au/mAbtmTO6BX05IdEILplNAgXv_Wc=/1200x675/https://prod.static9.net.au/fs/ff4238d6-65e7-4f73-afa2-537d3f64378e", XCoordinates: 20, YCoordinates: 35},
//	{TaskId: uuid.MustParse("a3958663-e7c4-4b19-a43b-b3a06fec33b5"), Title: "Pet Elvis", IconUrl: "https://scontent.fzrh3-1.fna.fbcdn.net/v/t39.30808-6/309430860_474589338041453_5993993981776996852_n.jpg?_nc_cat=108&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=7Tkot0fTAsoAX-FYOaX&_nc_ht=scontent.fzrh3-1.fna&oh=00_AfAsy-DH7-AJZmeKHnIMzUDN24XHC4BrAXHZ-SXf_ys_Kg&oe=643AC639", XCoordinates: 43, YCoordinates: 51},
//	{TaskId: uuid.MustParse("27504507-1ad7-48cf-b049-84163e5de019"), Title: "Visit Erika", IconUrl: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse3.mm.bing.net%2Fth%3Fid%3DOIP.9YTVRgjKM5V4xxRDVrMBcwHaJ4%26pid%3DApi&f=1&ipt=d917c56f3d343346bbafd198f4fb180c0dcecfd6ab290b62d2a57902489c2fcb&ipo=images", XCoordinates: 40, YCoordinates: 45},
//	{TaskId: uuid.MustParse("94e07a0c-5e5a-4af5-895c-98709be46c9d"), Title: "Get Coffee", IconUrl: "https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fpngimg.com%2Fuploads%2Fmug_coffee%2Fmug_coffee_PNG16824.png&f=1&nofb=1&ipt=b69c4e760e4328689109ac952ad7db0a8a02b86678be54356fea6609aa07b699&ipo=images", XCoordinates: 50, YCoordinates: 50},
//}

type TaskController interface {
	GetTasks(context *gin.Context)
	GetTaskById(context *gin.Context)
	//PostTask(context *gin.Context)
}

type taskController struct {
	taskService service.TaskService
}

func NewTaskController(taskServ service.TaskService) TaskController {
	return &taskController{
		taskService: taskServ,
	}
}

func (c *taskController) GetTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, c.taskService.GetAllTasks())
}

func (c *taskController) GetTaskById(context *gin.Context) {
	id := uuid.Must(uuid.Parse(context.Param("id")))

	//for _, t := range tasks {
	//	if t.TaskId == id {
	//		context.IndentedJSON(http.StatusOK, t)
	//		return
	//	}
	//}
	//
	//context.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})

	context.IndentedJSON(http.StatusOK, c.taskService.GetTask(id))
}

//func (c *taskController) PostTask(context *gin.Context) {
//	var newTask data.Task
//
//	if err := context.BindJSON(&newTask); err != nil {
//		return
//	}
//
//	tasks = append(tasks, newTask)
//	context.IndentedJSON(http.StatusCreated, newTask)
//}
