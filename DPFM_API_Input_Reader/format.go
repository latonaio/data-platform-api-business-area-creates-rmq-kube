package dpfm_api_input_reader

import (
	"data-platform-api-business-area-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBusinessArea() *requests.BusinessArea {
	data := sdc.BusinessArea
	return &requests.BusinessArea{
		BusinessArea: data.BusinessArea,
	}
}
