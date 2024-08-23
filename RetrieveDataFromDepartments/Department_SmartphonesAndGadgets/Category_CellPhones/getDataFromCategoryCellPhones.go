package Category_CellPhones

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategorySmartphones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Смартфоны")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryPushButtonPhones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кнопочные телефоны")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}
