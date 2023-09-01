package sheetdb

type GetContentParams struct {
	Sheet          string `json:"sheet"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	SortBy         string `json:"sort_by"`
	SortOrder      string `json:"sort_order"`
	SortMethod     string `json:"sort_method"`
	SortDateFormat string `json:"sort_date_format"`
	CastNumbers    string `json:"cast_numbers"`
	SingleObject   bool   `json:"single_object"`
	Mode           string `json:"mode"`
}

type Content map[string]interface{}

type Contents []Content

type TableKeys []map[string]interface{}

type DocumentName struct {
	Name string `json:"name"`
}

type Count struct {
	Rows int `json:"rows"`
}
