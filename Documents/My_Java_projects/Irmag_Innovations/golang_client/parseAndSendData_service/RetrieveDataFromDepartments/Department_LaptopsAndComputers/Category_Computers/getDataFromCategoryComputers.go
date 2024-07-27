package Category_Computers

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryMonoblocks(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Моноблоки")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryDesktopComputers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Настольные компьютеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryNettops(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Неттопы")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryMonitors(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Мониторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}
