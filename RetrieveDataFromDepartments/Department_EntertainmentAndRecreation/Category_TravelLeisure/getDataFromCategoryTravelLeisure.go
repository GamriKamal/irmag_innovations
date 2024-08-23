package Category_TravelLeisure

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategorySuitcases(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чемоданы")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryBagsBackpacks(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сумки и рюкзаки")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryTourism(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Туризм и отдых на природе")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryAirMattresses(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Надувные матрасы")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryBeachHoliday(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Пляжный отдых")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryPools(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Бассейны")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryPoolAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аксессуары для бассейнов")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryTrampolines(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Батуты")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}

func GetDataCategoryInflatableBoats(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Надувные лодки")
	Kafka.SendMessageToBroker(dataChannel, "Department_EntertainmentAndRecreation")
}
