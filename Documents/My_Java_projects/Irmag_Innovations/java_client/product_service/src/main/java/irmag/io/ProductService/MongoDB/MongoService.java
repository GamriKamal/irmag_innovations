package irmag.io.ProductService.MongoDB;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.mongodb.client.FindIterable;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.model.Indexes;
import irmag.io.ProductService.Entity.Category;
import irmag.io.ProductService.Entity.InfoAboutDepartmentsAndCategories;
import irmag.io.ProductService.Entity.KeywordCategoryMap;
import org.bson.Document;
import org.bson.conversions.Bson;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Service;
import org.w3c.dom.stylesheets.DocumentStyle;

import javax.print.Doc;
import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;

import static com.mongodb.client.model.Filters.eq;

@Service
public class MongoService {

    private final ObjectMapper objectMapper = new ObjectMapper();
    Logger logger = LoggerFactory.getLogger(MongoService.class);
    private MongoTemplate mongoTemplate;

    @Autowired
    public void setMongoTemplate(MongoTemplate mongoTemplate) {
        this.mongoTemplate = mongoTemplate;
    }


    public MongoCollection<Document> getNameOfCollection() {
        return mongoTemplate.getCollection("products");
    }

    public void insertDataToMongo(String data) throws JsonProcessingException {
        Map<String, Object> messageMap = objectMapper.readValue(data, Map.class);
        if (messageMap == null || messageMap.isEmpty()) {
            logger.error("Parsed message map is null or empty.");
            return;
        }

        if(!existItemByTitle(messageMap.get("title").toString())){
            messageMap.remove("_id");
            try {
                getNameOfCollection().insertOne(new Document(messageMap));
            } catch (Exception e) {
                e.printStackTrace();
            }
            ensureIndexExists();
        }
    }

    public List<Document> getDataByCategory(String category, int page, int pageSize) {
        page -= 1;
        if (category == null || category.isEmpty()) {
            logger.error("Category is null or empty.");
            return new ArrayList<>();
        }

        Bson filter = eq("category", category);
        FindIterable<Document> documents = getNameOfCollection().find(filter)
                .skip(page * pageSize)
                .limit(pageSize);

        return StreamSupport.stream(documents.spliterator(), false)
                .toList();
    }

    public Long getTotalDocuments(String category) {
        Bson filter = eq("category", category);
        return getNameOfCollection().countDocuments(filter);
    }

    public Document getProductByID(long id) {
        Bson filter = eq("id_item", id);
        Document tempDoc = getNameOfCollection().find(filter).first();
        if (tempDoc == null) {
            throw new NullPointerException("There is no item with this id!");
        } else {
            return tempDoc;
        }
    }

    public boolean existItemByUID(String uid){
        Bson filter = eq("UID товара", uid);
        Document tempDoc = getNameOfCollection().find(filter).first();
        if (tempDoc == null) {
            return false;
        } else {
            return true;
        }
    }

    public boolean existItemByTitle(String title){
        Bson filter = eq("title", title);
        Document tempDoc = getNameOfCollection().find(filter).first();
        if (tempDoc == null) {
            return false;
        } else {
            return true;
        }
    }

    public List<Document> searchDocuments(String searchText) {
        List<Document> results = searchDocumentsByField(searchText);
        if(results.isEmpty()){
            results = searchDocumentsFullText(searchText);
        }

        for(Document document : results){
            if(document != null){
                logger.info("class: " + document.get("photos").getClass());
            }
        }

        return results;
    }

    private List<Document> searchDocumentsByField(String fieldValue){
        String[] allSearchWords = fieldValue.split("\\s+");
        List<Document> orQueries = new ArrayList<>();
        for (String str : allSearchWords) {
            orQueries.add(new Document("title", str));
            orQueries.add(new Document("category", str));
            orQueries.add(new Document("Производитель", str));
            orQueries.add(new Document("Модель", str));
            orQueries.add(new Document("UID товара", str));
            orQueries.add(new Document("description", str));
        }
        Document query = new Document("$or", orQueries);

        List<Document> results = new ArrayList<>();
        mongoTemplate.getCollection("products").find(query).into(results);
        return results;
    }

    private List<Document> searchDocumentsFullText(String fullText){
        String[] parts = fullText.toLowerCase().split("\\s+");

        Map<String, List<String>> keywordCategoryMap = KeywordCategoryMap.getKeywordMap();

        List<String> allSearchWords = new ArrayList<>(Arrays.asList(parts));
        for (String part : parts) {
            if (keywordCategoryMap.containsKey(part)) {
                allSearchWords.addAll(keywordCategoryMap.get(part));
            }
        }

        List<Document> orQueries = new ArrayList<>();
        for (String str : allSearchWords) {
            orQueries.add(new Document("title", str));
            orQueries.add(new Document("category", str));
            orQueries.add(new Document("Производитель", str));
            orQueries.add(new Document("Модель", str));
            orQueries.add(new Document("UID товара", str));
            orQueries.add(new Document("description", str));
        }
        Document query = new Document("$or", orQueries);

        List<Document> results = new ArrayList<>();
        getNameOfCollection().find(query).into(results);

        return results;
    }

    public List<Document> getRandomDocuments(){
        List<Category> list = InfoAboutDepartmentsAndCategories.getCategories();
        List<Document> result = new ArrayList<>();
        for(Category category : list){
            Random random = new Random();
            Document query = new Document("category", category.getName());
            getNameOfCollection().find(query).limit(random.nextInt(2)).into(result);
        }

        Collections.shuffle(result);
        return result;
    }

    private void ensureIndexExists() {
        String collectionName = getNameOfCollection().toString();
        List<String> existingTextIndexes = getTextIndexes(collectionName);

        List<String> requiredIndexes = List.of("category", "Производитель", "Модель", "UID товара", "description", "title");

        if (!existingTextIndexes.containsAll(requiredIndexes) || existingTextIndexes.size() != requiredIndexes.size()) {
            existingTextIndexes.forEach(index -> dropIndex(collectionName, index));

            requiredIndexes.forEach(field -> getNameOfCollection().createIndex(Indexes.text(field)));
        }
    }

    private List<String> getTextIndexes(String collectionName) {
        return mongoTemplate.getCollection(collectionName).listIndexes().into(new ArrayList<>()).stream()
                .filter(index -> index.get("key").toString().contains("text"))
                .map(index -> index.get("key").toString())
                .collect(Collectors.toList());
    }

    private void dropIndex(String collectionName, String index) {
        mongoTemplate.getCollection(collectionName).dropIndex(index);
    }


//    private void validateDocumentFields(Document document) {
//        List<String> requiredFields = List.of("category", "Производитель", "title", "UID товара", "Модель");
//        for (String field : requiredFields) {
//            if (!document.containsKey(field)) {
//                throw new IllegalArgumentException("Document must contain the field '" + field + "'");
//            }
//        }
//    }
}
