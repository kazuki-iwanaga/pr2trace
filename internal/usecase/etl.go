package usecase

import (
	"github.com/kazuki-iwanaga/pr2trace/internal/domain"
)

type (
	// EtlUsecase is an input port for ETL usecase.
	EtlUsecase interface {
		Execute(i *EtlInput) (*EtlOutput, error)
	}

	// EtlInput is an input data for ETL usecase.
	EtlInput struct {
		Query string
	}

	// EtlPresenter is an output port for ETL usecase.
	EtlPresenter interface {
		Output(o *EtlOutput) EtlOutput
	}

	// EtlOutput is an output data for ETL usecase.
	EtlOutput struct {
		Result string
	}

	// EtlInteractor is an interactor for ETL usecase.
	EtlInteractor struct {
		prGateway    domain.PrGateway
		traceGateway domain.TraceGateway
		presenter    EtlPresenter
	}
)

// NewEtlInput creates a new EtlInteractor.
func NewEtlInteractor(
	prGateway domain.PrGateway,
	traceGateway domain.TraceGateway,
	presenter EtlPresenter,
) *EtlInteractor {
	return &EtlInteractor{
		prGateway:    prGateway,
		traceGateway: traceGateway,
		presenter:    presenter,
	}
}

func (i *EtlInteractor) Execute(input *EtlInput) (*EtlOutput, error) {
	// Extract PRs
	prs, err := i.prGateway.Search(input.Query)
	if err != nil {
		return nil, err
	}

	// Transform PRs to Traces
	prRootSpans := make([]*domain.Span, 0, len(prs))

	for _, pr := range prs {
		prRootSpan, err := domain.PrRootSpan("PrRootSpan", pr)
		if err != nil {
			return nil, err
		}

		// prRootSpan.AddChildTelemetry(foo)

		prRootSpans = append(prRootSpans, prRootSpan)
	}

	if len(prs) == 1 {
		trace := domain.NewTrace(prRootSpans[0])

		// Export Traces
		if err := i.traceGateway.Export(trace); err != nil {
			return nil, err
		}
	} else {
		multiPrRootSpan, err := domain.MultiPrRootSpan("MultiPrRootSpan", prs)
		if err != nil {
			return nil, err
		}

		multiPrRootSpan.AddChildSpans(prRootSpans...)

		trace := domain.NewTrace(multiPrRootSpan)
		if err := i.traceGateway.Export(trace); err != nil {
			return nil, err
		}
	}

	return &EtlOutput{Result: "success"}, nil
}
