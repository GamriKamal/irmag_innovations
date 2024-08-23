package RetrieveDataFromDepartments

import (
	"context"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_Accessories"
)

import (
	"sync"
)

func SendAllDataToKafka(ctx context.Context) error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		Department_Accessories.GetDataFromDepartmentAccessories(ctx)
	}()

	//go func() {
	//	defer wg.Done()
	//	Department_ComputerPeripherals.GetDataFromDepartment_ComputerPeripherals(ctx)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	Department_EntertainmentAndRecreation.GetDataFromDepartment_EntertainmentAndRecreation(ctx)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	Department_HomeAppliances.GetDataFromDepartmentHomeAppliances(ctx)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	Department_HouseholdGoods.GetDataFromDepartmentHouseholdGoods(ctx)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	Department_LaptopsAndComputers.GetDataFromDepartment_LaptopsAndComputers(ctx)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	Department_NetworkAndServerEquipment.GetDataFromDepartment_NetworkAndServerEquipment(ctx)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	Department_OfficeEquipmentAndSupplies.GetDataFromDepartment_OfficeEquipmentAndSupplies(ctx)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	Department_ProductsForGamers.GetDataFromDepartment_ProductsForGamers(ctx)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	Department_SmartphonesAndGadgets.GetDataFromDepartment_SmartphonesAndGadgets(ctx)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	Department_TelevisionsAudioVideo.GetDataFromDepartment_TelevisionsAudioVideo(ctx)
	//}()

	wg.Wait()

	return nil
}
