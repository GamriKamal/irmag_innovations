package irmag.io.auth_service.Controllers;

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import irmag.io.auth_service.DTO.JwtAuthenticationResponse;
import irmag.io.auth_service.DTO.SignInRequest;
import irmag.io.auth_service.DTO.SignUpRequest;
import irmag.io.auth_service.Models.User;
import irmag.io.auth_service.Services.AuthenticationService;
import irmag.io.auth_service.Services.FetchYandexUserInfo;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping("/auth")
@Tag(name = "Authentication")
public class RestAuthController {
    private AuthenticationService authenticationService;
    private final FetchYandexUserInfo fetchYandexUserInfo;
    private static final Logger logger = LoggerFactory.getLogger(RestAuthController.class);

    @Autowired
    public RestAuthController(AuthenticationService authenticationService, FetchYandexUserInfo fetchYandexUserInfo) {
        this.authenticationService = authenticationService;
        this.fetchYandexUserInfo = fetchYandexUserInfo;
    }

    @Autowired
    public void setAuthenticationService(AuthenticationService authenticationService) {
        this.authenticationService = authenticationService;
    }

    @Operation(summary = "User registration")
    @PostMapping("/sign-up")
    public JwtAuthenticationResponse signUp(@Valid SignUpRequest request, HttpServletResponse response) throws IOException {
        response.sendRedirect("/page");
        return authenticationService.signUp(request);
    }

    @Operation(summary = "User sign in")
    @PostMapping("/sign-in")
    public JwtAuthenticationResponse signIn(@Valid SignInRequest request, HttpServletResponse response) throws IOException {
        response.sendRedirect("/page");
        return authenticationService.signIn(request);
    }


    @PostMapping("/receiveToken")
    public ResponseEntity<Map<String, String>> receiveToken(@RequestBody Map<String, String> requestData) {
        String url = requestData.get("url");
        try {
            String token = FetchYandexUserInfo.extractAccessToken(url);
            logger.info("Received token: " + token);
            JwtAuthenticationResponse response = fetchYandexUserInfo.fetchUserInfo(token);
            if (response != null) {
                logger.info("Token processed successfully.");
                Map<String, String> responseBody = new HashMap<>();
                responseBody.put("token", response.getToken());
                return ResponseEntity.ok(responseBody);
            } else {
                logger.error("Error processing token.");
                throw new RuntimeException("Authentication failed.");
            }
        } catch (Exception e) {
            logger.error("Error processing token: ", e);
            throw new RuntimeException("Error processing token.", e);
        }
    }

}