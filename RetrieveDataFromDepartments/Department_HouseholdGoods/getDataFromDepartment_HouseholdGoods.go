package Department_HouseholdGoods

import (
	"context"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_HouseholdGoods/Category_BeautyHealth"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_HouseholdGoods/Category_KarcherProducts"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_HouseholdGoods/Category_Kitchenware"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_HouseholdGoods/Category_Other"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_HouseholdGoods/Category_SmartHome"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_HouseholdGoods/Category_ToolsAndPowerTools"
)

func GetDataFromDepartmentHouseholdGoods(ctx context.Context) {
	//Category_BeautyHealth
	Category_BeautyHealth.GetDataCategoryScales(ctx)
	Category_BeautyHealth.GetDataCategoryShaversEpilators(ctx)
	Category_BeautyHealth.GetDataCategoryClippersTrimmers(ctx)
	Category_BeautyHealth.GetDataCategoryHairDryersCurlers(ctx)
	Category_BeautyHealth.GetDataCategoryHealthHygiene(ctx)
	Category_BeautyHealth.GetDataCategoryFacialCleaningDevices(ctx)
	Category_BeautyHealth.GetDataCategoryOralCare(ctx)
	Category_BeautyHealth.GetDataCategoryMassagers(ctx)
	Category_BeautyHealth.GetDataCategoryTreadmills(ctx)

	//Category_KarcherProducts
	Category_KarcherProducts.GetDataCategoryPressureWashersKarcher(ctx)
	Category_KarcherProducts.GetDataCategoryVacuumCleanersKarcher(ctx)
	Category_KarcherProducts.GetDataCategoryHouseholdChemistryKarcher(ctx)
	Category_KarcherProducts.GetDataCategoryGardenEquipmentKarcher(ctx)

	//Category_Kitchenware
	Category_Kitchenware.GetDataCategoryFryingPans(ctx)
	Category_Kitchenware.GetDataCategoryPots(ctx)
	Category_Kitchenware.GetDataCategoryKitchenAccessories(ctx)
	Category_Kitchenware.GetDataCategoryKnives(ctx)
	Category_Kitchenware.GetDataCategoryServingWare(ctx)

	//Category_SmartHome
	Category_SmartHome.GetDataCategorySurveillanceSystems(ctx)
	Category_SmartHome.GetDataCategorySmartHomeDevices(ctx)
	Category_SmartHome.GetDataCategoryPetSupplies(ctx)
	Category_SmartHome.GetDataCategorySmartLighting(ctx)
	Category_SmartHome.GetDataCategoryRobotVacuums(ctx)
	Category_SmartHome.GetDataCategoryYandexStation(ctx)

	//Category_ToolsAndPowerTools
	Category_ToolsAndPowerTools.GetDataCategoryDrillsScrewdrivers(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryRotaryHammers(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryGrindingMachines(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryElectricJigsawsSaws(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryToolBatteries(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryPressureWashers(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryToolSets(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryCrimpingTools(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryLaserRangeFinders(ctx)
	Category_ToolsAndPowerTools.GetDataCategorySolderingEquipment(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryWeldingEquipment(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryHandTools(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryScrewdrivers(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryGardenTools(ctx)
	Category_ToolsAndPowerTools.GetDataCategoryOtherToolsEquipment(ctx)

	//Category_Other
	Category_Other.GetDataCategoryWallClocks(ctx)
	Category_Other.GetDataCategorySewingMachines(ctx)
	Category_Other.GetDataCategoryLamps(ctx)
	Category_Other.GetDataCategoryFlashlights(ctx)
	Category_Other.GetDataCategoryBatteriesChargers(ctx)
	Category_Other.GetDataCategoryCleaningSupplies(ctx)
}
