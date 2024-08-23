package Category_Kitchenware

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryFryingPans(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сковороды, сотейники")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryPots(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кастрюли")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryKitchenAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кухонные принадлежности")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryKnives(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Ножи")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryServingWare(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Посуда для сервировки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}
