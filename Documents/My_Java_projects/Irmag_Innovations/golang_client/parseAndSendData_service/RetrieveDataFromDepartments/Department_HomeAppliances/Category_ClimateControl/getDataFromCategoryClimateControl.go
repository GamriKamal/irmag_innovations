package Category_ClimateControl

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryAirConditioners(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кондиционеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryHeaters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Обогреватели")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryFans(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Вентиляторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryHumidifiersAndAirPurifiers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Увлажнители и очистители воздуха")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}
