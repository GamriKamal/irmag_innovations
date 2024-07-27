package Category_Periphery

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryMonitors(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Мониторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryMonitorMounts(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Крепления для мониторов")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryMouses(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Мыши")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryMousePads(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Коврики для мышей")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryKeyboards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Клавиатуры")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryKeyboardAccessories(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аксессуары для клавиатур")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryStickersForKeyboards(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Наклейки для клавиатур")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryAcousticSystemsAndSpeakers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Акустические системы и колонки")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryHeadphonesAndHeadsets(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Наушники и гарнитуры")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryMicrophones(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Микрофоны")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryWebcams(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Веб-камеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryUPS(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "ИБП (UPS)")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryBatteriesForUPS(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Аккумуляторы для ИБП")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryStabilizers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Стабилизаторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryNetworkFiltersAdapters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сетевые фильтры, адаптеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryJoysticksHandlebarsGamepads(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Джойстики, рули, геймпады")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryControllersForStreaming(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Контроллеры для стриминга")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryGraphicTablets(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Графические планшеты")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryUSBHub(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "USB-hub")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}

func GetDataCategoryCardReaders(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Картридеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_ComputerPeripherals")
}
