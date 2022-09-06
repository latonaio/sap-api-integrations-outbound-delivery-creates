package sap_api_caller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-outbound-delivery-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-outbound-delivery-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostOutboundDelivery(
	header *requests.Header,
	item *requests.Item,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(header)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(item)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(header *requests.Header) {
	outputDataHeader, err := c.callOutboundDeliverySrvAPIRequirementHeader("A_OutbDeliveryHeader", header)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataHeader)
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementHeader(api string, header *requests.Header) (*sap_api_output_formatter.Header, error) {
	body, err := json.Marshal(header)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(item *requests.Item) {
	url := fmt.Sprintf("A_OutbDeliveryHeader('%s')/to_Item", item.DeliveryDocument)
	outputDataItem, err := c.callOutboundDeliverySrvAPIRequirementItem(url, item)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataItem)
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementItem(api string, item *requests.Item) (*sap_api_output_formatter.Item, error) {
	body, err := json.Marshal(item)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}

	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
