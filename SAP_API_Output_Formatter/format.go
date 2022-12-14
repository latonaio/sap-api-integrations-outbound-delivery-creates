package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-outbound-delivery-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
			DeliveryDocument:              data.DeliveryDocument,
			DeliveryDocumentType:          data.DeliveryDocumentType,
			DocumentDate:                  data.DocumentDate,
			ActualGoodsMovementDate:       data.ActualGoodsMovementDate,
			ActualDeliveryRoute:           data.ActualDeliveryRoute,
			Shippinglocationtimezone:      data.Shippinglocationtimezone,
			Receivinglocationtimezone:     data.Receivinglocationtimezone,
			ActualGoodsMovementTime:       data.ActualGoodsMovementTime,
			BillingDocumentDate:           data.BillingDocumentDate,
			CompleteDeliveryIsDefined:     data.CompleteDeliveryIsDefined,
			ConfirmationTime:              data.ConfirmationTime,
			CreationDate:                  data.CreationDate,
			CreationTime:                  data.CreationTime,
			CustomerGroup:                 data.CustomerGroup,
			DeliveryBlockReason:           data.DeliveryBlockReason,
			DeliveryDate:                  data.DeliveryDate,
			DeliveryDocumentBySupplier:    data.DeliveryDocumentBySupplier,
			DeliveryIsInPlant:             data.DeliveryIsInPlant,
			DeliveryPriority:              data.DeliveryPriority,
			DeliveryTime:                  data.DeliveryTime,
			GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
			GoodsIssueTime:                data.GoodsIssueTime,
			HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
			HeaderGrossWeight:             data.HeaderGrossWeight,
			HeaderNetWeight:               data.HeaderNetWeight,
			HeaderVolume:                  data.HeaderVolume,
			HeaderVolumeUnit:              data.HeaderVolumeUnit,
			HeaderWeightUnit:              data.HeaderWeightUnit,
			IncotermsClassification:       data.IncotermsClassification,
			IsExportDelivery:              data.IsExportDelivery,
			LastChangeDate:                data.LastChangeDate,
			LoadingDate:                   data.LoadingDate,
			LoadingPoint:                  data.LoadingPoint,
			LoadingTime:                   data.LoadingTime,
			MeansOfTransport:              data.MeansOfTransport,
			OrderCombinationIsAllowed:     data.OrderCombinationIsAllowed,
			OrderID:                       data.OrderID,
			OverallDelivConfStatus:        data.OverallDelivConfStatus,
			OverallDelivReltdBillgStatus:  data.OverallDelivReltdBillgStatus,
			OverallGoodsMovementStatus:    data.OverallGoodsMovementStatus,
			OverallPackingStatus:          data.OverallPackingStatus,
			OverallPickingConfStatus:      data.OverallPickingConfStatus,
			OverallPickingStatus:          data.OverallPickingStatus,
			PickingDate:                   data.PickingDate,
			PickingTime:                   data.PickingTime,
			PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
			ReceivingPlant:                data.ReceivingPlant,
			ShippingCondition:             data.ShippingCondition,
			ShippingPoint:                 data.ShippingPoint,
			ShippingType:                  data.ShippingType,
			ShipToParty:                   data.ShipToParty,
			SoldToParty:                   data.SoldToParty,
			Supplier:                      data.Supplier,
			TransportationGroup:           data.TransportationGroup,
			TransportationPlanningDate:    data.TransportationPlanningDate,
			TransportationPlanningTime:    data.TransportationPlanningTime,
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
			DeliveryDocument:               data.DeliveryDocument,
			DeliveryDocumentItem:           data.DeliveryDocumentItem,
			BaseUnit:                       data.BaseUnit,
			ActualDeliveryQuantity:         data.ActualDeliveryQuantity,
			Batch:                          data.Batch,
			BatchBySupplier:                data.BatchBySupplier,
			CostCenter:                     data.CostCenter,
			CreationDate:                   data.CreationDate,
			CreationTime:                   data.CreationTime,
			DeliveryDocumentItemCategory:   data.DeliveryDocumentItemCategory,
			DeliveryDocumentItemText:       data.DeliveryDocumentItemText,
			DeliveryQuantityUnit:           data.DeliveryQuantityUnit,
			DistributionChannel:            data.DistributionChannel,
			Division:                       data.Division,
			GLAccount:                      data.GLAccount,
			GoodsMovementStatus:            data.GoodsMovementStatus,
			GoodsMovementType:              data.GoodsMovementType,
			InternationalArticleNumber:     data.InternationalArticleNumber,
			InventorySpecialStockType:      data.InventorySpecialStockType,
			IsCompletelyDelivered:          data.IsCompletelyDelivered,
			ItemBillingBlockReason:         data.ItemBillingBlockReason,
			ItemDeliveryIncompletionStatus: data.ItemDeliveryIncompletionStatus,
			ItemGdsMvtIncompletionSts:      data.ItemGdsMvtIncompletionSts,
			ItemGrossWeight:                data.ItemGrossWeight,
			ItemNetWeight:                  data.ItemNetWeight,
			ItemWeightUnit:                 data.ItemWeightUnit,
			ItemIsBillingRelevant:          data.ItemIsBillingRelevant,
			ItemPackingIncompletionStatus:  data.ItemPackingIncompletionStatus,
			ItemPickingIncompletionStatus:  data.ItemPickingIncompletionStatus,
			ItemVolume:                     data.ItemVolume,
			LastChangeDate:                 data.LastChangeDate,
			LoadingGroup:                   data.LoadingGroup,
			Material:                       data.Material,
			MaterialByCustomer:             data.MaterialByCustomer,
			MaterialFreightGroup:           data.MaterialFreightGroup,
			NumberOfSerialNumbers:          data.NumberOfSerialNumbers,
			OrderID:                        data.OrderID,
			OrderItem:                      data.OrderItem,
			OriginalDeliveryQuantity:       data.OriginalDeliveryQuantity,
			PackingStatus:                  data.PackingStatus,
			PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
			PickingConfirmationStatus:      data.PickingConfirmationStatus,
			PickingStatus:                  data.PickingStatus,
			Plant:                          data.Plant,
			ProductAvailabilityDate:        data.ProductAvailabilityDate,
			ProductAvailabilityTime:        data.ProductAvailabilityTime,
			ProfitCenter:                   data.ProfitCenter,
			StorageLocation:                data.StorageLocation,
			TransportationGroup:            data.TransportationGroup,
	}

	return item, nil
}
