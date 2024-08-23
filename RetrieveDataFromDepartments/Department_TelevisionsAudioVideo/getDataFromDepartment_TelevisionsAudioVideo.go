package Department_TelevisionsAudioVideo

import (
	"context"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_TelevisionsAudioVideo/Category_Accessories"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_TelevisionsAudioVideo/Category_Audio"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_TelevisionsAudioVideo/Category_DemoEquipment"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_TelevisionsAudioVideo/Category_Photography"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_TelevisionsAudioVideo/Category_TV"
	"parseAndSendData_service/RetrieveDataFromDepartments/Department_TelevisionsAudioVideo/Category_Video"
)

func GetDataFromDepartment_TelevisionsAudioVideo(ctx context.Context) {
	//Category_Accessories
	Category_Accessories.GetDataCategoryMemoryCards(ctx)
	Category_Accessories.GetDataCategoryTripodsMonopods(ctx)
	Category_Accessories.GetDataCategoryChargers(ctx)
	Category_Accessories.GetDataCategoryCleaningSupplies(ctx)
	Category_Accessories.GetDataCategoryHeadphoneCases(ctx)
	Category_Accessories.GetDataCategoryCablesAdapters(ctx)

	//Category_Audio
	Category_Audio.GetDataCategoryAudioTechnics(ctx)
	Category_Audio.GetDataCategoryPlayers(ctx)
	Category_Audio.GetDataCategoryPortableSpeakers(ctx)
	Category_Audio.GetDataCategoryAcousticSystems(ctx)
	Category_Audio.GetDataCategoryHeadphones(ctx)
	Category_Audio.GetDataCategoryPhoneHeadphones(ctx)
	Category_Audio.GetDataCategoryWirelessHeadphones(ctx)
	Category_Audio.GetDataCategoryMicrophones(ctx)
	Category_Audio.GetDataCategoryMicrophoneAccessories(ctx)
	Category_Audio.GetDataCategorySoundCards(ctx)

	//Category_DemoEquipment
	Category_DemoEquipment.GetDataCategoryProjectors(ctx)
	Category_DemoEquipment.GetDataCategoryProjectionScreens(ctx)
	Category_DemoEquipment.GetDataCategoryProjectorMounts(ctx)

	//Category_Photography
	Category_Photography.GetDataCategoryCameras(ctx)
	Category_Photography.GetDataCategoryCameraAccessories(ctx)
	Category_Photography.GetDataCategoryBags(ctx)

	//Category_TV
	Category_TV.GetDataCategoryTelevisions(ctx)
	Category_TV.GetDataCategorySoundbars(ctx)
	Category_TV.GetDataCategoryMediaPlayers(ctx)
	Category_TV.GetDataCategoryTVTuners(ctx)
	Category_TV.GetDataCategoryTVWallMounts(ctx)
	Category_TV.GetDataCategoryRemoteControls(ctx)
	Category_TV.GetDataCategoryTVAntennas(ctx)

	//Category_Video
	Category_Video.GetDataCategoryActionCameras(ctx)
	Category_Video.GetDataCategoryActionCameraAccessories(ctx)
	Category_Video.GetDataCategoryDroneAccessories(ctx)
	Category_Video.GetDataCategorySurveillanceSystems(ctx)
	Category_Video.GetDataCategoryWebCameras(ctx)
}
