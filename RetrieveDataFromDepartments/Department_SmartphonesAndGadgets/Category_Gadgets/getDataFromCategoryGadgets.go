package Category_Gadgets

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategorySmartClocks(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Смарт-часы")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryFitnessBracelet(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Фитнес-браслет")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryTablets(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Планшеты")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryEBooks(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Электронные книги")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}
