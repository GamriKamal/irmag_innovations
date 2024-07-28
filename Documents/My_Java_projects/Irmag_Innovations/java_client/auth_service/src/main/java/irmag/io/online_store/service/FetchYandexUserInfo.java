package irmag.io.online_store.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import irmag.io.online_store.domain.dto.SignUpRequest;
import irmag.io.online_store.domain.dto.YandexUserInfo;
import irmag.io.online_store.domain.model.User;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

import java.io.UnsupportedEncodingException;
import java.net.URI;
import java.net.URLDecoder;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;

@Service
public class FetchYandexUserInfo {

    private static AuthenticationService authenticationService;

    public static User fetchUserInfo(String oauthToken) {
        Logger logger = LoggerFactory.getLogger(FetchYandexUserInfo.class);

        String apiUrl = "https://login.yandex.ru/info?format=json";

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(apiUrl))
                .header("Authorization", "OAuth " + oauthToken)
                .GET()
                .build();

        try {
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
            ObjectMapper objectMapper = new ObjectMapper();

            if (response.statusCode() == 200) {
                try {
                    YandexUserInfo userInfo = objectMapper.readValue(response.body(), YandexUserInfo.class);
                    SignUpRequest yandexInfoRequest = new SignUpRequest();
                    yandexInfoRequest.setUsername(userInfo.getDisplayName());
                    yandexInfoRequest.setPassword("");
                    yandexInfoRequest.setEmail(userInfo.getDefaultEmail());
                    yandexInfoRequest.setPhoneNumber("");
                    authenticationService.signUp(yandexInfoRequest);
                } catch (Exception e) {
                    e.printStackTrace();
                }

            } else {
                logger.error("GET request failed. Response Code: " + response.statusCode());
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return new User();
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

        logger.info("\n\n\n" + accessToken);
        return accessToken;
    }

    @Autowired
    public void setAuthenticationService(AuthenticationService authenticationService) {
        FetchYandexUserInfo.authenticationService = authenticationService;
    }
}
