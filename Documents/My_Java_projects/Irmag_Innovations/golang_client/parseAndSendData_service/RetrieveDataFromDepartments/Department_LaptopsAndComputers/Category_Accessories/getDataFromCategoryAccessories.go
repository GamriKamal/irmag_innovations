package Category_Accessories

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryCleaningProducts(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чистящие средства")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryCablesAdapterss(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кабели, переходники")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryDockingStationsHubs(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Док-станции, концентраторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryOtherAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Прочие аксессуары")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}
