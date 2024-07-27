package Category_Peripherals

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryGamingMice(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые мыши")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingMousePads(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые коврики для мышей")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingKeyboards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые клавиатуры")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingHeadsets(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые наушники")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingControllers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Джойстики, рули, геймпады")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryStreamingControllers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Контроллеры для стриминга")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingMonitors(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые мониторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}
