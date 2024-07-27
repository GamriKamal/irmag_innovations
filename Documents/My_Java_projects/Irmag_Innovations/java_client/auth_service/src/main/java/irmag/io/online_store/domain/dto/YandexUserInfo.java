package irmag.io.online_store.domain.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

@Getter
@NoArgsConstructor
@AllArgsConstructor
public class YandexUserInfo {
    private String id;
    private String login;
    @JsonProperty("client_id")
    private String clientId;
    @JsonProperty("display_name")
    private String displayName;
    @JsonProperty("real_name")
    private String realName;
    @JsonProperty("first_name")
    private String firstName;
    @JsonProperty("last_name")
    private String lastName;
    private String sex;
    @JsonProperty("default_email")
    private String defaultEmail;
    private String[] emails;
    private String birthday;
    private String psuid;
}

