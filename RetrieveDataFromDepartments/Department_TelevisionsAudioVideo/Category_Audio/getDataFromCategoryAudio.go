package Category_Audio

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryAudioTechnics(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аудиотехника, Hi-Fi оборудование")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryPlayers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Плееры")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryPortableSpeakers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Портативные колонки")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryAcousticSystems(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Акустические системы и колонки")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryHeadphones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Наушники и гарнитуры")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryPhoneHeadphones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Наушники для телефонов")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryWirelessHeadphones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Беспроводные наушники и гарнитуры")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryMicrophones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Микрофоны")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryMicrophoneAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аксессуары для микрофонов")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategorySoundCards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Звуковые карты и усилители")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}
