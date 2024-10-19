package micro_queue

import (
	"fmt"
	"github.com/ysw-jingzhe/go-toolbox/examples/micro_queue/routine"
	"time"
)

func main() {
	routine.Initialize()

	routine.PushQueue(&routine.Task{TID: 1001})
	routine.PushQueue(&routine.Task{TID: 1002})
	routine.PushQueue(&routine.Task{TID: 1003})

	time.Sleep(5 * time.Second)
	fmt.Println("*** debug finish ***")
}
