package irmag.io.auth_service.Services;

import com.fasterxml.jackson.databind.ObjectMapper;
import irmag.io.auth_service.DTO.JwtAuthenticationResponse;
import irmag.io.auth_service.DTO.SignInRequest;
import irmag.io.auth_service.DTO.SignUpRequest;
import irmag.io.auth_service.DTO.YandexUserInfo;
import irmag.io.auth_service.Models.User;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.net.URI;
import java.net.URLDecoder;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;

@Service
public class FetchYandexUserInfo {

    private final AuthenticationService authenticationService;
    private static Logger logger = LoggerFactory.getLogger(FetchYandexUserInfo.class);
    private final UserService userService;
    private final JwtService jwtService;

    @Autowired
    public FetchYandexUserInfo(AuthenticationService authenticationService, UserService userService, JwtService jwtService) {
        this.authenticationService = authenticationService;
        this.userService = userService;
        this.jwtService = jwtService;
    }

    public JwtAuthenticationResponse fetchUserInfo(String oauthToken) {
        String apiUrl = "https://login.yandex.ru/info?format=json";

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(apiUrl))
                .header("Authorization", "OAuth " + oauthToken)
                .GET()
                .build();

        try {
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

            if (response.statusCode() == 200) {
                logger.info("Response Body: {}", response.body());
                ObjectMapper objectMapper = new ObjectMapper();
                YandexUserInfo userInfo = objectMapper.readValue(response.body(), YandexUserInfo.class);

                if (userInfo != null && userInfo.getDefaultEmail() != null) {
                    logger.info(userInfo.getDefaultEmail());
                    User existingUser = userService.findByEmail(userInfo.getDefaultEmail());
                    if (existingUser == null) {
                        SignUpRequest yandexInfoRequest = new SignUpRequest();
                        yandexInfoRequest.setUsername(userInfo.getDisplayName());
                        yandexInfoRequest.setPassword(generateRandomPassword());
                        yandexInfoRequest.setEmail(userInfo.getDefaultEmail());

                        return authenticationService.signUp(yandexInfoRequest);
                    } else {
                        return authenticationService.signIn(new SignInRequest(userInfo.getDisplayName(), userInfo.getDefaultEmail()));
                    }

                } else {
                    logger.warn("User info or email is missing.");
                    return null;
                }

            } else {
                logger.error("GET request failed. Response Code: {}", response.statusCode());
                return null;
            }

        } catch (Exception e) {
            logger.error("HTTP request failed: {}", e.getMessage());
            return null;
        }
    }


    public static String extractAccessToken(String url) {
        Logger logger = LoggerFactory.getLogger(FetchYandexUserInfo.class);
        String accessToken = null;
        String[] parts = url.split("#|&");
        for (String part : parts) {
            if (part.startsWith("access_token=")) {
                accessToken = part.substring("access_token=".length());
                accessToken = URLDecoder.decode(accessToken, StandardCharsets.UTF_8);
                break;
            }
        }

        return accessToken;
    }

    private String generateRandomPassword() {
        return "randomPassword123";
    }

}
