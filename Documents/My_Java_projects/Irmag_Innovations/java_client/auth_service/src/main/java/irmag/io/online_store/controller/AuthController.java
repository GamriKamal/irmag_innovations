package irmag.io.online_store.controller;

import io.swagger.v3.oas.annotations.Operation;
import irmag.io.online_store.domain.dto.SignInRequest;
import org.springframework.http.MediaType;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@org.springframework.stereotype.Controller
@RequestMapping(value = "/auth")
public class AuthController {

    @Operation(summary = "Авторизация пользователя (Показ страницы)")
    @GetMapping("/sign-in")
    public String signIn(Model model) {
        model.addAttribute("SignInRequest", new SignInRequest());
        return "sign-in";
    }

    @GetMapping(value = "/sign-up", consumes = {"application/xml"})
    public String signUp(Model model) {
        model.addAttribute("SignInRequest", new SignInRequest());
        return "sign-up";
    }

    @GetMapping("/getToken")
    public String getToken() {
        return "pageToGetToken";
    }

}
