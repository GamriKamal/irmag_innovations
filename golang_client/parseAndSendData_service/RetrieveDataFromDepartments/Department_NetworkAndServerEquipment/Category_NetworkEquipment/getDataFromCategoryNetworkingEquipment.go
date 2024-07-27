package Category_NetworkingEquipment

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryRoutersAndModems(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Роутеры, модемы")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryNetworkCardsAndAdapters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сетевые карты и адаптеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryAccessPoints(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Точки доступа")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategorySwitches(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Коммутаторы (switch)")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryAntennasAndSignalAmplifiers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Антенны, усилители сигнала")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryPowerlineAdapters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Powerline-адаптеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryNetworkStorages(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сетевые хранилища")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryMediaConverters(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Медиаконвертеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}

func GetDataCategoryOtherEquipment(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Прочее оборудование")
	Kafka.SendMessageToBroker(dataChannel, "Department_NetworkAndServerEquipment")
}
