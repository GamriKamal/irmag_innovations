package irmag.io.ProductService.Kafka;

import com.fasterxml.jackson.core.JsonProcessingException;
import irmag.io.ProductService.MongoDB.MongoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
public class KafkaConsumerService {
    private MongoService mongo_service;

    @Autowired
    public void setKafka(MongoService kafka) {
        this.mongo_service = kafka;
    }

    @KafkaListener(topics = "Department_Accessories")
    public void listenDepartment_Accessories(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_ComputerPeripherals")
    public void listenDepartment_ComputerPeripherals(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_EntertainmentAndRecreation")
    public void listenDepartment_EntertainmentAndRecreation(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_HomeAppliances")
    public void listenDepartment_HomeAppliances(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_HouseholdGoods")
    public void listenDepartment_HouseholdGoods(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_LaptopsAndComputers")
    public void listenDepartment_LaptopsAndComputers(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_NetworkAndServerEquipment")
    public void listenDepartment_NetworkAndServerEquipment(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_OfficeEquipmentAndSupplies")
    public void listenDepartment_OfficeEquipmentAndSupplies(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_ProductsForGamers")
    public void listenDepartment_ProductsForGamers(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_SmartphonesAndGadgets")
    public void listenDepartment_SmartphonesAndGadgets(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

    @KafkaListener(topics = "Department_TelevisionsAudioVideo")
    public void listenDepartment_TelevisionsAudioVideo(String message) throws JsonProcessingException {
        if (message != null) {
            mongo_service.insertDataToMongo(message);
        }
    }

}