package Category_PrintingDevices

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryLaserMFPs(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "МФУ Лазерные")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}
func GetDataCategoryLaserPrinters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Принтеры лазерные")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryInkjetPrinters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Принтеры струйные")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryCopiers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Копировальные аппараты")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryPlotters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Плоттеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategory3DPrinters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "3D принтеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategory3DPens(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "3D ручки")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}
