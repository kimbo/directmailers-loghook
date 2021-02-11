package api

type RecipientDetails struct {
	Name         string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Zip          string
}

type SenderDetails RecipientDetails

// https://jsapi.apiary.io/apis/directmailers/reference/letter/create-letter/create-letter.html
type LetterRequest struct {
	Description     string              // Friendly object description.
	Size            string              // Size of postcard to be mailed. Allowed values are '8.5x11' and '8.5x14'
	Duplex          bool                // Set to true to enable two sided printing.
	DryRun          bool                // Set to true for testing. This will suppress printing and mailing and will result in no unit cost.
	WaitForRender   bool                // Set to true to disable asynchronous thumbnail rendering. Useful for forcing the API to not respond until all thumbnails and renders are available.
	BlankFirstPage  bool                // Set to true to insert a blank first page.
	PostalClass     string              // Mail class of letter to be sent. Allowed values are 'First Class' and 'Marketing Mail'
	Data            string              // Creative to be used for the letter. Value must be either a HTML Template id, public PDF URL, or HTML string less than 30,000 characters
	VariablePayload []map[string]string // Key Value Array with variable names and values to replace in front and back object data HTML string or HTML Template
	To              RecipientDetails
	From            SenderDetails
}

type PostcardRequest struct {
	Description     string
	Size            string // Size of postcard to be mailed. Allowed values are '4.25x6'
	Front           string // Creative to be used for the front (address side) of the postcard. Value must be either a HTML Template id, public PDF or Image URL, or HTML string less than 30,000 characters
	Back            string // Creative to be used for the back (non address side) of the postcard. Value must be either a HTML Template id, public PDF or Image URL, or HTML string less than 30,000 characters
	VariablePayload []map[string]string
	DryRun          bool
	WaitForRender   bool
	To              RecipientDetails
	From            SenderDetails
}

type LetterResponse struct {
	PrintRecord     string `json:"PrintRecord"`
	Created         string `json:"Created"`
	MailingDate     string `json:"MailingDate"`
	Canceled        bool   `json:"Canceled"`
	Status          string `json:"Status"`
	Description     string `json:"Description"`
	Medium          string `json:"Medium"`
	Size            string `json:"Size"`
	VariablePayload map[string]string
	To              RecipientDetails
	From            SenderDetails
	Cost            float64 `json:"Cost"`
	PdfPages        int     `json:"PdfPages"`
	PrintPages      int     `json:"PrintPages"`
	DryRun          bool    `json:"DryRun"`
	Duplex          bool    `json:"Duplex"`
	BlankFirstPage  bool    `json:"BlankFirstPage"`
	RenderedPdf     string  `json:"RenderedPdf"`
	PostalCarrier   string  `json:"PostalCarrier"`
	PostalClass     string  `json:"PostalClass"`
	Data            string  `json:"Data"`
	Thumbnails      struct {
		Small  string `json:"Small"`
		Medium string `json:"Medium"`
		Large  string `json:"Large"`
	} `json:"Thumbnails"`
	TrackingEvents []struct {
		UspsScanTime          string `json:"UspsScanTime"`
		TrackingOperationCode string `json:"TrackingOperationCode"`
		Description           string `json:"Description"`
		City                  string `json:"City"`
		State                 string `json:"State"`
		Latitude              string `json:"Latitude"`
		Longitude             string `json:"Longitude"`
	} `json:"TrackingEvents"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate"`
	ActualDeliveryDate    string `json:"ActualDeliveryDate"`
}

type PostcardResponse struct {
	PrintRecord     string `json:"PrintRecord"`
	Created         string `json:"Created"`
	MailingDate     string `json:"MailingDate"`
	Canceled        bool   `json:"Canceled"`
	Status          string `json:"Status"`
	Description     string `json:"Description"`
	Medium          string `json:"Medium"`
	Size            string `json:"Size"`
	Front           string `json:"Front"`
	Back            string `json:"Back"`
	VariablePayload map[string]string
	Cost            float64 `json:"Cost"`
	DryRun          bool    `json:"DryRun"`
	RenderedPdf     string  `json:"RenderedPdf"`
	PostalCarrier   string  `json:"PostalCarrier"`
	PostalClass     string  `json:"PostalClass"`
	To              RecipientDetails
	From            SenderDetails
	FrontThumbnails struct {
		Small  string `json:"Small"`
		Medium string `json:"Medium"`
		Large  string `json:"Large"`
	} `json:"FrontThumbnails"`
	BackThumbnails struct {
		Small  string `json:"Small"`
		Medium string `json:"Medium"`
		Large  string `json:"Large"`
	} `json:"BackThumbnails"`
	TrackingEvents        []interface{} `json:"TrackingEvents"`
	EstimatedDeliveryDate string        `json:"EstimatedDeliveryDate"`
	ActualDeliveryDate    interface{}   `json:"ActualDeliveryDate"`
}
