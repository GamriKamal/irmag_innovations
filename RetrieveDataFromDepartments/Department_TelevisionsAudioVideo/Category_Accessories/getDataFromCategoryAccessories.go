package Category_Accessories

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryMemoryCards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Карты памяти")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryTripodsMonopods(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Штативы, моноподы")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryChargers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Зарядные устройства")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryCleaningSupplies(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чистящие средства")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryHeadphoneCases(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чехлы для наушников")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}

func GetDataCategoryCablesAdapters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кабели, переходники")
	Kafka.SendMessageToBroker(dataChannel, "Department_TelevisionsAudioVideo")
}
