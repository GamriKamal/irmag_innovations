package Category_SmartHome

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategorySurveillanceSystems(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Системы видеонаблюдения")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategorySmartHomeDevices(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Датчики, устройства и управление для умного дома")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryPetSupplies(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Товары для животных")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategorySmartLighting(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Умное освещение")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryRobotVacuums(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Роботы-пылесосы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryYandexStation(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Колонки Яндекс Станция (Алиса)")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}
