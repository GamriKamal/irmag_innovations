package irmag.io.auth_service.Controllers;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class TestController {

    @GetMapping("/page")
    public String getPage(){
        return "page";
    }
}
