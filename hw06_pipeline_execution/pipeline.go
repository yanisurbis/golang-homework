package hw06_pipeline_execution //nolint:golint,stylecheck
import "fmt"

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)


func ExecutePipeline(in In, done In, stages ...Stage) Out {
	in_ := make(Bi)
	go func() {
		for v := range in {
			select {
				case <- done:
					fmt.Println("stopped")
					close(in_)
					return
				default:
					fmt.Println("not stopped")
			}
			fmt.Println(v)
			in_ <- v
		}
		close(in_)
	}()

	out := make(Out)
	for i, stage := range stages {
		if i == 0 {
			out = stage(in_)
		} else {
			out = stage(out)
		}
	}

	return out
}
