package Category_Software

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryAntiviruses(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Антивирусы")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryOperatingSystemss(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Операционные системы")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryOfficePrograms(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Офисные программы")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryDigitalSubscriptions(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Цифровые подписки")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}
