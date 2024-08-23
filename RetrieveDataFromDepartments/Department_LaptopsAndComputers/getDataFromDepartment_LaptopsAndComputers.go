package Department_LaptopsAndComputers

import (
	"context"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_LaptopsAndComputers/Category_Accessories"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_LaptopsAndComputers/Category_AccessoriesForLaptops"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_LaptopsAndComputers/Category_Computers"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_LaptopsAndComputers/Category_Laptops"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_LaptopsAndComputers/Category_Software"
)

func GetDataFromDepartment_LaptopsAndComputers(ctx context.Context) {
	//Category_Accessories
	Category_Accessories.GetDataCategoryCleaningProducts(ctx)
	Category_Accessories.GetDataCategoryCablesAdapterss(ctx)
	Category_Accessories.GetDataCategoryDockingStationsHubs(ctx)
	Category_Accessories.GetDataCategoryOtherAccessories(ctx)

	//Category_AccessoriesForLaptops
	Category_AccessoriesForLaptops.GetDataCategorySSDDrives(ctx)
	Category_AccessoriesForLaptops.GetDataCategoryLaptopHardDrives(ctx)
	Category_AccessoriesForLaptops.GetDataCategoryRAMForLaptops(ctx)
	Category_AccessoriesForLaptops.GetDataCategoryOtherAccessories(ctx)

	//Category_Computers
	Category_Computers.GetDataCategoryMonoblocks(ctx)
	Category_Computers.GetDataCategoryDesktopComputers(ctx)
	Category_Computers.GetDataCategoryNettops(ctx)
	Category_Computers.GetDataCategoryMonitors(ctx)

	//Category_Laptops
	Category_Laptops.GetDataCategoryLaptops(ctx)
	Category_Laptops.GetDataCategoryUltrabooks(ctx)
	Category_Laptops.GetDataCategoryLaptopBags(ctx)
	Category_Laptops.GetDataCategoryCoolingStands(ctx)
	Category_Laptops.GetDataCategoryLaptopChargers(ctx)

	//Category_Software
	Category_Software.GetDataCategoryAntiviruses(ctx)
	Category_Software.GetDataCategoryOperatingSystemss(ctx)
	Category_Software.GetDataCategoryOfficePrograms(ctx)
	Category_Software.GetDataCategoryDigitalSubscriptions(ctx)
}
