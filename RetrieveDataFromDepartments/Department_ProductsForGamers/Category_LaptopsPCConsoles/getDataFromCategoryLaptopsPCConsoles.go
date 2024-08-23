package Category_LaptopsPCConsoles

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryGamingLaptops(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые ноутбуки")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingPCs(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые компьютеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}
