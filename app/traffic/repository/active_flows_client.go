package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	services "github.com/PaoGRodrigues/tfi-backend/app/services"
	"github.com/PaoGRodrigues/tfi-backend/app/traffic/domains"
)

type ActiveFlowsClient struct {
	tool     *services.Tool
	endpoint string
}

type HttpResponse struct {
	Rc    int
	RcStr string
	Rsp   RspData
}

type RspData struct {
	Data        []domains.ActiveFlow
	CurrentPage int
	PerPage     int
}

func NewActiveFlowClient(tool *services.Tool, endpoint string) *ActiveFlowsClient {
	return &ActiveFlowsClient{
		tool:     tool,
		endpoint: endpoint,
	}
}

func (actF *ActiveFlowsClient) GetAllActiveTraffic() ([]domains.ActiveFlow, error) {
	activeFlows, err := actF.getActiveFlows()
	if err != nil {
		return nil, err
	}
	return activeFlows, nil
}

func (actF *ActiveFlowsClient) getActiveFlows() ([]domains.ActiveFlow, error) {
	activeFlows := []domains.ActiveFlow{}
	resp, err := actF.getActiveFlowsSinglePage(1)
	if err != nil {
		return nil, err
	}
	for len(resp.Rsp.Data) > resp.Rsp.PerPage {
		activeFlows = append(activeFlows, resp.Rsp.Data...)
		resp, err = actF.getActiveFlowsSinglePage(resp.Rsp.CurrentPage + 1)
		if err != nil {
			return nil, err
		}
	}
	activeFlows = append(activeFlows, resp.Rsp.Data...)
	return activeFlows, nil
}

func (actF *ActiveFlowsClient) getActiveFlowsSinglePage(page int) (HttpResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", actF.tool.UrlClient+actF.endpoint, nil)
	if err != nil {
		return HttpResponse{}, err
	}
	req.SetBasicAuth(actF.tool.Usr, actF.tool.Pass)
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("ifid", strconv.Itoa(actF.tool.InterfaceId))
	query.Add("currentPage", strconv.Itoa(page))

	req.URL.RawQuery = query.Encode()

	response, err := client.Do(req)
	if err != nil {
		return HttpResponse{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return HttpResponse{}, err
	}

	var resp HttpResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return HttpResponse{}, err
	}
	return resp, nil
}
