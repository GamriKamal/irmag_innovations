package Category_BeautyHealth

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryScales(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Весы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryShaversEpilators(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Электробритвы, эпиляторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryClippersTrimmers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Машинки для стрижки, триммеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryHairDryersCurlers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Фены, электрощипцы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryHealthHygiene(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Здоровье и гигиена")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryFacialCleaningDevices(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Приборы для чистки лица")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryOralCare(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Уход за полостью рта")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryMassagers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Массажеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryTreadmills(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Беговые дорожки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}
