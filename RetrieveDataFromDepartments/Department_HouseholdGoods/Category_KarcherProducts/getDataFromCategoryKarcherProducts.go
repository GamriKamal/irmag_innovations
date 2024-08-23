package Category_KarcherProducts

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryPressureWashersKarcher(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Мойки высокого давления Karcher")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryVacuumCleanersKarcher(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Пылесосы Karcher")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryHouseholdChemistryKarcher(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Бытовая химия Karcher")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryGardenEquipmentKarcher(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Оборудование для полива Karcher")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}
