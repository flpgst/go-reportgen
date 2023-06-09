package pdf

import (
	"os"

	"github.com/flpgst/go-reportgen/internal/dto"
)

type PDFBuilderInterface interface {
	GeneratePDF(dto *dto.ReportOutputDTO) (*os.File, error)
}
