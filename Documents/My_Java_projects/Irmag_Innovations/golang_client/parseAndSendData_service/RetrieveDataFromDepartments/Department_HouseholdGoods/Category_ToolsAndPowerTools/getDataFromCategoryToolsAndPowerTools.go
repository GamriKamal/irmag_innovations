package Category_ToolsAndPowerTools

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryDrillsScrewdrivers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Дрели, шуруповерты")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryRotaryHammers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Перфораторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryGrindingMachines(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Шлифовальные машины")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryElectricJigsawsSaws(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Электролобзики, пилы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryToolBatteries(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аккумуляторы для электроинструментов")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryPressureWashers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Мойки высокого давления")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryToolSets(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Наборы инструментов")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryCrimpingTools(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Обжимной инструмент")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryLaserRangeFinders(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Лазерные дальномеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategorySolderingEquipment(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Паяльное оборудование")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryWeldingEquipment(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сварочное оборудование")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryHandTools(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Инструменты")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryScrewdrivers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Отвертки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryGardenTools(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Садовая техника")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}

func GetDataCategoryOtherToolsEquipment(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Прочие инструменты, оборудование")
	Kafka.SendMessageToBroker(dataChannel, "Department_HouseholdGoods")
}
