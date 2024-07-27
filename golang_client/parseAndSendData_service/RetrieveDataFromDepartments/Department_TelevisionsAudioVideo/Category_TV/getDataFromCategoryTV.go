package Category_TV

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryTelevisions(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Телевизоры")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategorySoundbars(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Саундбары")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryMediaPlayers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Медиаплееры")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryTVTuners(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "ТВ-тюнеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryTVWallMounts(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кронштейны для телевизоров")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryRemoteControls(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Пульты")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryTVAntennas(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Антенны для телевизора")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}
