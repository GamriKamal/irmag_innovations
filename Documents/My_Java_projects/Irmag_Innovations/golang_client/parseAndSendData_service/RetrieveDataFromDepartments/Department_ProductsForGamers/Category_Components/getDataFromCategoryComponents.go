package Category_Components

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryGamingMotherboards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые материнские платы")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingRAM(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровая оперативная память")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingGPUs(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Видеокарты для игр")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingSSD(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "SSD для игр")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}

func GetDataCategoryGamingPowerSupplies(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые блоки питания")
	Kafka.SendMessageToBroker(dataChannel, "Department_ProductsForGamers")
}
