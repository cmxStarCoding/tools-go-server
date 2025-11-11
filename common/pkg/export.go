package pkg

type Export struct {
	Headers     map[string]string
	ColumnWidth map[string]int
}

func NewExport(headers map[string]string, columnWidth map[string]int) *Export {
	return &Export{
		Headers:     headers,
		ColumnWidth: columnWidth,
	}
}
