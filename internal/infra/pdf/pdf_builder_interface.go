package pdf

import (
	"os"

	"github.com/flpgst/go-reportgen/internal/dto"
)

type PDFBuilderInterface interface {
	GeneratePDF(dto *dto.ReportDTO) (*os.File, error)
}
