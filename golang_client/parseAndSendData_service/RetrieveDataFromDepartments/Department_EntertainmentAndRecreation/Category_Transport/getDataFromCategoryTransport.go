package Category_Transport

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryHoverboards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Гироскутеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategorySegways(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сигвеи")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryElectricScootersBikes(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Электросамокаты и велосипеды")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}
