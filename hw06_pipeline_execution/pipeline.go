package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here
	//for v := range in {
	//	fmt.Println(v)
	//}

	out1 := stages[0](in)
	out2 := stages[1](out1)
	out3 := stages[2](out2)
	out4 := stages[3](out3)

	return out4
}
