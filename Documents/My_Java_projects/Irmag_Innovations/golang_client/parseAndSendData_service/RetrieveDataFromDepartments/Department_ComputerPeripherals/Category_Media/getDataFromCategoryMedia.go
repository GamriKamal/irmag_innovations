package Category_Media

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryFlashDrives(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Флешки")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryMemoryCards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Карты памяти")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryUSB(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "USB внешние накопители")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryNetworkStorage(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сетевые хранилища")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}
