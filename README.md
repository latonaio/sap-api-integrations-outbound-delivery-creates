# sap-api-integrations-outbound-delivery-creates 
sap-api-integrations-outbound-delivery-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で出荷データ を登録するマイクロサービスです。    
sap-api-integrations-outbound-delivery-creates には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-outbound-delivery-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_OUTBOUND_DELIVERY_SRV_0002/overview  

## 動作環境  
sap-api-integrations-outbound-delivery-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-outbound-delivery-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。 

## 本レポジトリ が 対応する API サービス
sap-api-integrations-outbound-delivery-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_OUTBOUND_DELIVERY_SRV_0002/overview  
* APIサービス名(=baseURL): API_OUTBOUND_DELIVERY_SRV;v=0002

## 本レポジトリ に 含まれる API名
sap-api-integrations-outbound-delivery-creates には、次の API をコールするためのリソースが含まれています。  

* A_OutbDeliveryHeader（出荷伝票 - ヘッダ）
* A_OutbDeliveryItem（出荷伝票 - 明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "SAPOutboundDeliveryCreates",
	"accepter": ["Header"],
	"delivery_document": "",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPOutboundDeliveryCreates",
	"accepter": ["All"],
	"delivery_document": "",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## SAP API Business Hub における API サービス の バージョン と バージョン におけるデータレイアウトの相違

SAP API Business Hub における API サービス のうちの 殆どの API サービス の バージョンは、v1.X.X であり、そのため殆どの API サービス 間 の データレイアウトは統一されています。
従って、Latona および AION における リソースにおいても、これらの 揃った v1.X.X のバージョンに従い、データレイアウトが統一されています。
一方、本レポジトリ に関わる API である Outbound Delivery のサービスは、v2.X.X までバージョンが進んでおり、その結果、本レポジトリ内の一部のAPIのデータレイアウトが、他のAPIサービスのものと異なっています。

* v1.X.X のバージョンである API サービス の データレイアウト（=responses）  
v1.X.X のバージョンであるAPIサービスのデータレイアウト（=responses）は、例えば、次の通りです。

```
type PartnerFunction struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SDDocument           string `json:"SDDocument"`
			PartnerFunction      string `json:"PartnerFunction"`
			Customer             string `json:"Customer"`
			Supplier             string `json:"Supplier"`
		} `json:"results"`
	} `json:"d"`
}


```

* 本レポジトリ内の一部のAPIのデータレイアウトを含む、v2.X.X のバージョンである API サービス の データレイアウト（=responses）  
本レポジトリ内の一部のAPIのデータレイアウトを含む、v2.X.X のバージョンである API サービスのデータレイアウト（=responses）は、次の通りです。  
ここで言う本レポジトリ内の一部のAPIのデータレイアウトには、PartnerAddressが含まれ、  
本レポジトリ内の他のデータレイアウトは、v1.X.X と同様のデータレイアウトとなっています。
（つまり、本レポジトリ内においても、現時点でv2.X.X のデータレイアウトを踏襲しているのは、PartnerAddress のみで、Header、PartnerFunction、Item、などのAPIでは、v1.X.X と同様のデータレイアウトとなっています）  

```
type PartnerAddress struct {
	D *struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			    Etag string `json:"etag"`
			} `json:"__metadata"`
			AddressID              string `json:"AddressID"`
			BusinessPartnerName1   string `json:"BusinessPartnerName1"`
			Country                string `json:"Country"`
			Region                 string `json:"Region"`
			StreetName             string `json:"StreetName"`
			CityName               string `json:"CityName"`
			PostalCode             string `json:"PostalCode"`
			CorrespondenceLanguage string `json:"CorrespondenceLanguage"`
			FaxNumber              string `json:"FaxNumber"`
			PhoneNumber            string `json:"PhoneNumber"`
	} `json:"d"`
}
```

このように、v1.X.X のバージョンである API サービス の データレイアウトと、v2.X.X のバージョンである API サービス の データレイアウトは、Results の配列構造を持っているか持っていないかという点が異なります。  

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 出荷伝票 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"DeliveryDocument" ～ "to_DeliveryDocumentItem" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-outbound-delivery-creates/SAP_API_Caller/caller.go#L68",
	"function": "sap-api-integrations-outbound-delivery-creates/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"DeliveryDocument": "80000000",
			"DeliveryDocumentType": "LF",
			"DocumentDate": "2016-08-17T09:00:00+09:00",
			"ActualGoodsMovementDate": "2016-08-17T09:00:00+09:00",
			"ActualDeliveryRoute": "TR0002",
			"Shippinglocationtimezone": "PST",
			"Receivinglocationtimezone": "EST",
			"ActualGoodsMovementTime": "PT07H00M00S",
			"BillingDocumentDate": "2016-08-17T09:00:00+09:00",
			"CompleteDeliveryIsDefined": false,
			"ConfirmationTime": "PT00H00M00S",
			"CreationDate": "2016-08-17T09:00:00+09:00",
			"CreationTime": "PT08H21M06S",
			"CustomerGroup": "01",
			"DeliveryBlockReason": "",
			"DeliveryDate": "2016-08-19T09:00:00+09:00",
			"DeliveryDocumentBySupplier": "",
			"DeliveryIsInPlant": false,
			"DeliveryPriority": "00",
			"DeliveryTime": "PT04H00M00S",
			"GoodsIssueOrReceiptSlipNumber": "",
			"GoodsIssueTime": "PT07H00M00S",
			"HeaderBillingBlockReason": "",
			"HeaderGrossWeight": "3.000",
			"HeaderNetWeight": "2.700",
			"HeaderVolume": "0.000",
			"HeaderVolumeUnit": "",
			"HeaderWeightUnit": "G",
			"IncotermsClassification": "EXW",
			"IsExportDelivery": "",
			"LastChangeDate": "",
			"LoadingDate": "2016-08-17T09:00:00+09:00",
			"LoadingPoint": "",
			"LoadingTime": "PT07H00M00S",
			"MeansOfTransport": "",
			"OrderCombinationIsAllowed": false,
			"OrderID": "",
			"OverallDelivConfStatus": "",
			"OverallDelivReltdBillgStatus": "C",
			"OverallGoodsMovementStatus": "C",
			"OverallPackingStatus": "",
			"OverallPickingConfStatus": "",
			"OverallPickingStatus": "C",
			"PickingDate": "2016-08-16T09:00:00+09:00",
			"PickingTime": "PT07H00M00S",
			"PlannedGoodsIssueDate": "2016-08-17T09:00:00+09:00",
			"ReceivingPlant": "",
			"ShippingCondition": "01",
			"ShippingPoint": "1710",
			"ShippingType": "",
			"ShipToParty": "17100001",
			"SoldToParty": "17100001",
			"Supplier": "",
			"TransportationGroup": "0001",
			"TransportationPlanningDate": "2016-08-17T09:00:00+09:00",
			"TransportationPlanningTime": "PT07H00M00S",
			"to_DeliveryDocumentPartner": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_OUTBOUND_DELIVERY_SRV;v=0002/A_OutbDeliveryHeader('80000000')/to_DeliveryDocumentPartner",
			"to_DeliveryDocumentItem": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_OUTBOUND_DELIVERY_SRV;v=0002/A_OutbDeliveryHeader('80000000')/to_DeliveryDocumentItem"
		}
	],
	"time": "2022-01-27T22:51:12+09:00"
}
```
