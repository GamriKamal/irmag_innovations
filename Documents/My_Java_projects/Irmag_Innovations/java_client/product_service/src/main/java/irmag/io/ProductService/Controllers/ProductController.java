package irmag.io.ProductService.Controllers;

import irmag.io.ProductService.Entity.Category;
import irmag.io.ProductService.Entity.Department;
import irmag.io.ProductService.Entity.InfoAboutDepartmentsAndCategories;
import irmag.io.ProductService.MongoDB.MongoService;
import org.bson.Document;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;

import java.io.IOException;
import java.time.LocalDate;
import java.util.List;

@Controller
public class ProductController {
    private static Long totalDocuments = 0L;
    private final Logger logger = LoggerFactory.getLogger(ProductController.class);
    InfoAboutDepartmentsAndCategories info = new InfoAboutDepartmentsAndCategories();
    private MongoService mongoService;

    @Autowired
    public void setMongoService(MongoService mongoService) {
        this.mongoService = mongoService;
    }

    @GetMapping("/")
    public String getWelcomePage(Model model){
        model.addAttribute("featuredProducts", mongoService.getRandomDocuments());
        return "welcome-page";
    }

    @GetMapping("/categories")
    public String getMainPage(Model model) {
        model.addAttribute("departments", info.getDepartments());
        return "main-page";
    }

    @GetMapping("/categories/{departmentName}")
    public String getCategoriesPage(@PathVariable String departmentName, Model model) {
        Department department = info.getDepartments()
                .stream()
                .filter(d -> d.getName().equals(departmentName))
                .findFirst()
                .orElse(null);

        model.addAttribute("department", department);
        model.addAttribute("departmentName", departmentName);
        return "category";
    }


    @GetMapping("/categories/{departmentName}/{categoryName}")
    public String getCategoryPage(
            @PathVariable String departmentName,
            @PathVariable String categoryName,
            Model model) {
        Department department = info.getDepartments()
                .stream()
                .filter(d -> d.getName().equals(departmentName))
                .findFirst()
                .orElse(null);

        Category category = department.getCategories()
                .stream()
                .filter(c -> c.getName().equals(categoryName))
                .findFirst()
                .orElse(null);

        model.addAttribute("categoryName", categoryName);
        model.addAttribute("category", category);
        model.addAttribute("departmentName", departmentName);
        return "subcategory";
    }

    @GetMapping("/categories/{departmentName}/{categoryName}/{subcategoryName}")
    public String getProductsByCategoryPage(@PathVariable String departmentName,
                                            @PathVariable String categoryName,
                                            @PathVariable String subcategoryName,
                                            @RequestParam(defaultValue = "1") int page,
                                            @RequestParam(defaultValue = "15") int pageSize,
                                            Model model) {
        List<Document> productList = mongoService.getDataByCategory(subcategoryName, page, pageSize);
        model.addAttribute("productList", productList);

        totalDocuments = Math.max(totalDocuments, mongoService.getTotalDocuments(subcategoryName));

        int totalPages = (int) Math.ceil(totalDocuments / pageSize);

        logger.info("totalDocuments: " + totalDocuments);
        model.addAttribute("categoryName", categoryName);
        model.addAttribute("subcategoryName", subcategoryName);
        model.addAttribute("departmentName", departmentName);
        model.addAttribute("currentPage", page);
        model.addAttribute("totalPages", totalPages);

        return "listOfProductPage";
    }

    @GetMapping("/product/{id_item}")
    public String getProductByID(@PathVariable("id_item") String id, Model model) {
        Document tempDoc = mongoService.getProductByID(Long.parseLong(id));
        model.addAttribute("product", tempDoc);
        return "product-page";
    }

    @GetMapping("/search")
    public String search(@RequestParam String text, Model model){
        model.addAttribute("resultList", mongoService.searchDocuments(text));
        model.addAttribute("searchInput", text);
        return "search-page";
    }

    @GetMapping("/privacy")
    public String privacyPolicy(Model model) {
        model.addAttribute("effectiveDate", LocalDate.now());
        return "privacy";
    }

    @GetMapping("/about")
    public String aboutUs(Model model) {
        model.addAttribute("foundationDate", "1 июня 2005 года");
        return "about";
    }

    @GetMapping("/terms")
    public String termsOfUse(Model model) {
        model.addAttribute("effectiveDate", "1 июня 2005 года");
        return "terms";
    }

    @GetMapping("/contacts")
    public String contactPage(Model model) {
        return "contacts";
    }

    @PostMapping("/send-message")
    public String sendMessage(@RequestParam String name,
                              @RequestParam String email,
                              @RequestParam String message,
                              Model model) {
        model.addAttribute("message", "Ваше сообщение было отправлено успешно!");
        return "contacts";
    }
}
