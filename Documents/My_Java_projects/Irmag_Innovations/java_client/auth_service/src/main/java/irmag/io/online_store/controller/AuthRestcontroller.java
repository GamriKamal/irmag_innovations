package irmag.io.online_store.controller;

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import irmag.io.online_store.domain.dto.JwtAuthenticationResponse;
import irmag.io.online_store.domain.dto.SignInRequest;
import irmag.io.online_store.domain.dto.SignUpRequest;
import irmag.io.online_store.service.AuthenticationService;
import irmag.io.online_store.service.FetchYandexUserInfo;
import jakarta.validation.Valid;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@RestController
@RequestMapping("/auth")
@Tag(name = "Аутентификация")
public class AuthRestcontroller {
    private AuthenticationService authenticationService;

    @Autowired
    public void setAuthenticationService(AuthenticationService authenticationService) {
        this.authenticationService = authenticationService;
    }

    @Operation(summary = "Регистрация пользователя")
    @PostMapping("/sign-up")
    public JwtAuthenticationResponse signUp(@RequestBody @Valid SignUpRequest request) {
        return authenticationService.signUp(request);
    }


    @Operation(summary = "Авторизация пользователя")
    @PostMapping("/sign-in")
    public JwtAuthenticationResponse signIn(@ModelAttribute @Valid SignInRequest request) {
        return authenticationService.signIn(request);
    }

    @PostMapping("/receiveToken")
    public void receiveToken(@RequestBody Map<String, String> requestData) {
        Logger logger = LoggerFactory.getLogger(YandexCallbackController.class);
        String url = requestData.get("url");
        logger.info("Received URL: " + url);
        try {
            String token = FetchYandexUserInfo.extractAccessToken(url);
            FetchYandexUserInfo.fetchUserInfo(token);
            logger.info("Token processed successfully.");
        } catch (Exception e) {
            logger.error("Error processing token: ", e);
        }
    }
}

