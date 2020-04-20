package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func CloseWhenDone(out Out, done In) Out {
	out1 := make(Bi)
	go func(out Out, out1 Bi) {
		for v := range out {
			select {
			case <-done:
				close(out1)
				return
			default:
			}
			out1 <- v
		}
		close(out1)
	}(out, out1)
	return out1
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
