package Category_InstallationEquipment

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryCablesConnectors(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кабели, коннекторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryCableChannelsAndAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кабель-каналы, фурнитура")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryTools(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Инструменты")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}
