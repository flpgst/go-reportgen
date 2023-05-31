package entity

type ReportRepositoryInterface interface {
	Save(report *Report) error
	GetReport(name, date string) (*Report, error)
}
