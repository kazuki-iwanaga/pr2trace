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
	EtlOutput struct{}

	// EtlInteractor is an interactor for ETL usecase.
	EtlInteractor struct {
		prExtractor domain.PrExtractor
		tlmExporter domain.TelemetryExporter
		presenter   EtlPresenter
	}
)

// NewEtlInput creates a new EtlInteractor.
func NewEtlInteractor(
	prExtractor domain.PrExtractor,
	tlmExporter domain.TelemetryExporter,
	presenter EtlPresenter,
) *EtlInteractor {
	return &EtlInteractor{
		prExtractor: prExtractor,
		tlmExporter: tlmExporter,
		presenter:   presenter,
	}
}

func (i *EtlInteractor) Execute(input *EtlInput) (*EtlOutput, error) {
	// Extract PRs
	_, err := i.prExtractor.Search(input.Query)
	if err != nil {
		return nil, err
	}

	// Transform PRs to Telemetries
	tlms := []*domain.Telemetry{}

	// Export Telemetries
	if err := i.tlmExporter.Export(tlms); err != nil {
		return nil, err
	}

	return &EtlOutput{}, nil
}
