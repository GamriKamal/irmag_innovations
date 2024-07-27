package Category_DemonstrationEquipment

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryProjectors(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Проекторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryProjectionScreens(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Экраны проекционные")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryInteractivePanels(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Интерактивные панели")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryProjectorMounts(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Крепления для проекторов")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryWhiteboardsAndAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Маркерные доски и аксессуары")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}
