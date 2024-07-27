package RetrieveDataFromDepartments

import (
	"context"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_Accessories"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_ComputerPeripherals"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_EntertainmentAndRecreation"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_HomeAppliances"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_HouseholdGoods"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_LaptopsAndComputers"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_NetworkAndServerEquipment"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_OfficeEquipmentAndSupplies"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_ProductsForGamers"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_SmartphonesAndGadgets"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_TelevisionsAudioVideo"
)

import (
	"sync"
)

func SendAllDataToKafka(ctx context.Context) error {
	var wg sync.WaitGroup

	wg.Add(11)
	go func() {
		defer wg.Done()
		Department_Accessories.GetDataFromDepartmentAccessories(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_ComputerPeripherals.GetDataFromDepartment_ComputerPeripherals(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_EntertainmentAndRecreation.GetDataFromDepartment_EntertainmentAndRecreation(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_HomeAppliances.GetDataFromDepartmentHomeAppliances(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_HouseholdGoods.GetDataFromDepartmentHouseholdGoods(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_LaptopsAndComputers.GetDataFromDepartment_LaptopsAndComputers(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_NetworkAndServerEquipment.GetDataFromDepartment_NetworkAndServerEquipment(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_OfficeEquipmentAndSupplies.GetDataFromDepartment_OfficeEquipmentAndSupplies(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_ProductsForGamers.GetDataFromDepartment_ProductsForGamers(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_SmartphonesAndGadgets.GetDataFromDepartment_SmartphonesAndGadgets(ctx)
	}()

	go func() {
		defer wg.Done()
		Department_TelevisionsAudioVideo.GetDataFromDepartment_TelevisionsAudioVideo(ctx)
	}()

	wg.Wait()

	return nil
}
