package asyncfilereader

type CSVFile struct {
	Header     string
	Rows       []string
	Name       string
	ImportedAt string
}

type ExcelFile struct {
	Header string
	Rows   []string
	Name   string
}
