package Category_Accessories

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryArmoredFilms(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Бронепленки")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryProtectiveGlassesAndFilms(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Защитные стекла и пленки")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryHeadphonesForPhones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Наушники для телефонов")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryWirelessHeadphonesAndHeadsets(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Беспроводные наушники и гарнитуры")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryPortableSpeakers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Портативные колонки")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategorySmartphoneCases(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чехлы для смартфонов")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryEbookCases(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чехлы для электронных книг")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryMountsStandsForSmartphones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Крепления, подставки для смартфонов")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryPowerBank(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Power bank (мобильные аккумуляторы)")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryMonopods(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Моноподы")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryChargers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Зарядные устройства")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryCablesAndAdapters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кабели и переходники")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryMemoryCards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Карты памяти")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryStrapsForSmartWatchesAndFitnessBracelets(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Ремешки для смарт-часов и фитнес-браслетов")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryHeadphoneCases(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чехлы для наушников")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryRingLamps(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кольцевые лампы")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryCleaningProducts(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чистящие средства")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}

func GetDataCategoryOtherAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Прочие аксессуары")
	Kafka.SendMessageToBroker(dataChannel, "Department_SmartphonesAndGadgets")
}
