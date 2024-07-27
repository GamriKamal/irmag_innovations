package Category_Photography

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryCameras(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Фотоаппараты")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryCameraAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аксессуары для фотоаппаратов")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryBags(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сумки для фото и видео")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}
