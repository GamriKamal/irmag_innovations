package Category_EverythingForBuilding_a_Computer

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryCPU(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Процессоры")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryMotherboards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Материнские платы")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryRAM(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Оперативная память")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryGPU(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Видеокарты")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryHardDrives(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Жесткие диски")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategorySDDrives(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "SSD диски")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryPowerSupplyUnits(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Блоки питания")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryEnclosures(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Корпуса")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryCPUCoolers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кулеры для процессора")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryWaterCoolingSystems(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Системы водяного охлаждения")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryThermalPaste(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Термопаста")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}

func GetDataCategoryDrives(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Оптические дисководы (приводы)")
	Kafka.SendMessageToBroker(dataChannel, "Department_Accessories")
}
