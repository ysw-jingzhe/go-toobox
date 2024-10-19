package routine

import "fmt"

func Runner(t *Task) {
	fmt.Println(fmt.Sprintf("task(%d) running ***", t.TID))
}
