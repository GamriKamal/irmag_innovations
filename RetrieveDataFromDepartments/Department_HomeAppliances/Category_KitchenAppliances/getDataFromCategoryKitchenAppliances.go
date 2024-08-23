package Category_KitchenAppliances

import (
	"context"
	"parseAndSendData_service/Kafka"
	"parseAndSendData_service/OperationsWithMongoDB"
)

func GetDataCategoryRefrigerators(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Холодильники")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryFreezers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Морозильники")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryIceMakers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Льдогенераторы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryWineCoolers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Винные шкафы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryMicrowaves(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Микроволновые печи")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryCooktops(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Варочные поверхности")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryOvens(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Встраиваемые духовки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryElectricOvens(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Электропечи")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryStoves(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Плиты")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryInductionCookers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Настольные плиты")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryRangeHoods(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Вытяжки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryKettles(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Чайники, термопоты")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryBlendersJuicers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Блендеры, соковыжималки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryMixers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Миксеры")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryFoodProcessors(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кухонные комбайны и машины")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryMultiCookers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Мультиварки, пароварки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryIceCreamMakers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Мороженицы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryFoodDehydrators(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сушилки для продуктов")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryYogurtMakers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Йогуртницы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryBreadMakers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Хлебопечи")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryCoffeeMachines(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кофеварки, кофемашины")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryCoffeeGrinders(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Кофемолки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryTurkishCoffeePots(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Турки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryKitchenScales(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Весы кухонные")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryToastersGrills(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Тостеры, гриль")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryVacuumSealers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Вакууматоры и аксессуары")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryWaffleMakers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Вафельницы, блинницы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryDeepFryers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Фритюрницы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategorySandwichMakers(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Сэндвичницы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryMeatGrinders(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Электромясорубки")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryElectricGrills(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Электрошашлычницы")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}

func GetDataCategoryKitchenGadgets(ctx context.Context) {
	dataChannel := OperationsWithMongoDB.GetDataByField(ctx, "Гаджеты для кухни")
	Kafka.SendMessageToBroker(dataChannel, "Department_HomeAppliances")
}
