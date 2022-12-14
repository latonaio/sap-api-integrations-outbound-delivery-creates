package sap_api_input_reader

import (
	"sap-api-integrations-outbound-delivery-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.OutboundDelivery
	return &requests.Header{
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
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataOutboundDelivery := sdc.OutboundDelivery
	data := sdc.OutboundDelivery.DeliveryDocumentItem
	return &requests.Item{
		DeliveryDocument:               dataOutboundDelivery.DeliveryDocument,
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
}
