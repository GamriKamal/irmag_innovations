package Category_Laptops

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryLaptops(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Ноутбуки")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryUltrabooks(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Ультрабуки")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryLaptopBags(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сумки для ноутбуков")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryCoolingStands(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Подставки для охлаждения")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryLaptopChargers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Зарядные устройства для ноутбуков")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}
