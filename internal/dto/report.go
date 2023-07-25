package dto

type ReportDTO struct {
	ReportName string
	Date       string
	Header     []string
	Body       [][]string
	Footer     []string
	Template   ReportTemplate
}

type ReportTemplate struct {
	TemplateName string
	TableHeader  string
	TableBody    string
	TableFooter  string
}
