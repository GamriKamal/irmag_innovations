package Category_ConsumablesAndAccessories

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryPaperAndPhotoPaper(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Бумага, фотобумага")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryLaserCartridges(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Лазерные картриджи")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryInkCartridgesAndInks(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Струйные картриджи и чернила")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryTonersForLaserPrinters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Тонеры для лазерных принтеров")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryConsumablesForLaserPrinters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Расходные материалы для лазерных принтеров")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryDotMatrixCartridges(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Матричные картриджи")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryConsumablesForCopiers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Расходные материалы для копиров")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryFilmsForLaminators(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Пленки для ламинаторов")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryAccessoriesForBindingMachines(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аксессуары для переплетных машин")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryPlasticFor3DPrinting(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Пластик для 3D печати")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryCablesAndAdapters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кабели, переходники")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}

func GetDataCategoryCleaningProducts(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чистящие средства")
	Kafka.SendMessageToBroker(dataChannel, "Department_OfficeEquipmentAndSupplies")
}
