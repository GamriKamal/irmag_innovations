package irmag.io.auth_service.Controllers;

import irmag.io.auth_service.DTO.SignInRequest;
import irmag.io.auth_service.DTO.SignUpRequest;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("/auth")
public class AuthController {
    @GetMapping("/sign-up")
    public String showSignUpPage(Model model) {
        model.addAttribute("signUpRequest", new SignUpRequest());
        return "sign-up";
    }

    @GetMapping("/sign-in")
    public String showSignInPage(Model model) {
        model.addAttribute("signInRequest", new SignInRequest());
        return "sign-in";
    }

    @GetMapping("/getToken")
    public String getToken(HttpServletRequest request) {
        return "getToken";
    }


}
