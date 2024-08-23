package Category_Other

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryWallClocks(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Часы настенные")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategorySewingMachines(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Швейные машинки, оверлоки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryLamps(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Лампы, светильники")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryFlashlights(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Фонари")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryBatteriesChargers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Батарейки, зарядные устройства")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryCleaningSupplies(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чистящие средства")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}
