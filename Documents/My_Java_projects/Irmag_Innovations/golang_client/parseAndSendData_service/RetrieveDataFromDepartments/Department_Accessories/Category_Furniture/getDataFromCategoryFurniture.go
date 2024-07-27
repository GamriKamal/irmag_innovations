package Category_Furniture

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryGamingChairs(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Игровые кресла")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryTablesForGamers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Столы для геймеров")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}
