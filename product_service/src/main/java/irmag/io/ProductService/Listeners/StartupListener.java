package irmag.io.ProductService.Listeners;

import com.mongodb.client.MongoCollection;
import com.mongodb.client.model.Indexes;
import irmag.io.ProductService.MongoDB.MongoService;
import org.bson.Document;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.event.EventListener;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Component;

import java.util.ArrayList;

@Component
public class StartupListener {

    private MongoTemplate mongoTemplate;
    private MongoService mongoService;


    @Autowired
    public void setMongoTemplate(MongoTemplate mongoTemplate) {
        this.mongoTemplate = mongoTemplate;
    }

    @Autowired
    public void setMongoService(MongoService mongoService) {
        this.mongoService = mongoService;
    }

    @EventListener(ApplicationReadyEvent.class)
    public void init() {
        ensureTextIndexExists();
    }

    private void ensureTextIndexExists() {
        if (!isTextIndexExists(mongoService.getNameOfCollection())) {
            mongoService.getNameOfCollection().createIndex(Indexes.text("category"));
            mongoService.getNameOfCollection().createIndex(Indexes.text("Производитель"));
            mongoService.getNameOfCollection().createIndex(Indexes.text("Модель"));
            mongoService.getNameOfCollection().createIndex(Indexes.text("UID товара"));
            mongoService.getNameOfCollection().createIndex(Indexes.text("description"));
        }
    }

    private boolean isTextIndexExists(MongoCollection<Document> collection) {
        return collection.listIndexes().into(new ArrayList<>()).stream()
                .anyMatch(index -> index.get("key").toString().contains("text"));
    }

}

