package Category_AccessoriesForLaptops

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategorySSDDrives(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "SSD диски")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryLaptopHardDrives(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Жесткие диски для ноутбуков")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryRAMForLaptops(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Оперативная память для ноутбуков")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}

func GetDataCategoryOtherAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Прочие комплектующие")
	Kafka.SendMessageToBroker(dataChannel, "Department_LaptopsAndComputers")
}
