package Category_ServerEquipment

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryServers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Серверы")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryServerComponents(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Комплектующие для серверов")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryServerCabinets(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Серверные шкафы")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryCabinetComponents(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Комплектующие для шкафов")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}
