package Category_Video

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryActionCameras(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Экшн-камеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryActionCameraAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аксессуары для экшн-камер")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryDroneAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аксессуары для квадрокоптеров")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategorySurveillanceSystems(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Системы видеонаблюдения")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryWebCameras(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Веб-камеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}
