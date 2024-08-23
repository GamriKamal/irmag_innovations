package Category_DemoEquipment

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryProjectors(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Проекторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryProjectionScreens(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Экраны проекционные")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryProjectorMounts(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Крепления для проекторов")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}
