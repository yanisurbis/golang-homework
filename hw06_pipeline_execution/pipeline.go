package hw06_pipeline_execution //nolint:golint,stylecheck
import "fmt"

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func StopWhenDone(out Out, done In) Out {
	out_ := make(Bi)
	go func() {
		for v := range out {
			select {
			case <- done:
				fmt.Println("stopped")
				close(out_)
				return
			default:
				fmt.Println("not stopped")
			}
			fmt.Println(v)
			out_ <- v
		}
		close(out_)
	}()
	return out_
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Out)
	for i, stage := range stages {
		if i == 0 {
			out = StopWhenDone(stage(in), done)
		} else {
			out = StopWhenDone(stage(out), done)
		}
	}

	return out
}
