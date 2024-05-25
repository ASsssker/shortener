package models

type (
	UrlResponseModels struct {
		ResultUrl string `json:"result"`
	}

	RequestDataModels struct {
		Url string `json:"url"`
	}

	BatchRequestModel struct {
		CorrelationId string `json:"correlation_id"`
		OriginalUrl   string `json:"original_url"`
	}

	BatchResponseModel struct {
		CorrelationId string `json:"correlation_id"`
		ShortUrl      string `json:"short_url"`
	}
)
