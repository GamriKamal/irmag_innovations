package Category_HomeAppliances

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryVacuumCleaners(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Пылесосы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryVacuumCleanerAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аксессуары для пылесосов")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryWashingMachines(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Стиральные машины")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryWaterHeaters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Водонагреватели")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryFaucetWaterHeaters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Краны-водонагреватели")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryIronsSteamers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Утюги, отпариватели")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryClothingCare(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Уход за одеждой")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryHomeCleaning(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Уборка дома")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategorySteamCleaners(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Пароочистители")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryWindowCleaners(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Стеклоочистители")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryPhones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Телефоны")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryStabilizers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Стабилизаторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}
