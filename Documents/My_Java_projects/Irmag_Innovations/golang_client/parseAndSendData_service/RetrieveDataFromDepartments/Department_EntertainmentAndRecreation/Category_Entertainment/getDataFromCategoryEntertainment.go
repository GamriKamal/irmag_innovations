package Category_Entertainment

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryBinoculars(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Бинокли")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryEBooks(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Электронные книги")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}
