package Category_OfficeAndCommercialEquipment

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryTelephones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Телефоны")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryCutters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Резаки")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryLaminators(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Ламинаторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryShredders(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Уничтожители документов")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryBindingMachines(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Переплетное оборудование")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryTradeEquipment(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Торговое оборудование")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}
