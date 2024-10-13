package usecase

type (
	// Pr2TraceUsecase is an input port for Pr2Trace usecase.
	Pr2TraceUsecase interface {
		Execute(i *Pr2TraceInput) (*Pr2TraceOutput, error)
	}

	// Pr2TraceInput is an input data for Pr2Trace usecase.
	Pr2TraceInput struct {
		Query string
	}

	// Pr2TracePresenter is an output port for Pr2Trace usecase.
	Pr2TracePresenter interface {
		Output(o *Pr2TraceOutput) Pr2TraceOutput
	}

	// Pr2TraceOutput is an output data for Pr2Trace usecase.
	Pr2TraceOutput struct{}
)
