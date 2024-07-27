package Category_Toys

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryRCToys(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Радиоуправляемые игрушки")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryToyCars(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игрушечные машинки")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryConstructors(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Конструкторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryLEGO(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "LEGO")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryEducationalToys(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Развивающие игрушки и робототехника")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategory3DPens(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "3D ручки")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}
