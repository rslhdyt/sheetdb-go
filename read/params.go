package read

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

func (p *GetContentParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("sheet", p.Sheet)
	urlValues.Add("limit", strconv.Itoa(p.Limit))
	urlValues.Add("offset", strconv.Itoa(p.Offset))
	urlValues.Add("sort_by", p.SortBy)
	urlValues.Add("sort_order", p.SortOrder)
	urlValues.Add("sort_method", p.SortMethod)
	urlValues.Add("sort_date_format", p.SortDateFormat)
	urlValues.Add("cast_numbers", p.CastNumbers)
	urlValues.Add("single_object", strconv.FormatBool(p.SingleObject))
	urlValues.Add("mode", p.Mode)

	return urlValues.Encode()
}