package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func CloseWhenDone(out Out, done In) Out {
	out_ := make(Bi)
	go func() {
		for v := range out {
			select {
			case <- done:
				close(out_)
				return
			default:
			}
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
			out = CloseWhenDone(stage(in), done)
		} else {
			out = CloseWhenDone(stage(out), done)
		}
	}
	return out
}
