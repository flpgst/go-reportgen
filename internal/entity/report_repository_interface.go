package entity

type ReportRepositoryInterface interface {
	Save(report *Report) error
}
