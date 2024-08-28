package irmag.io.auth_service.DTO;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import irmag.io.auth_service.Models.User;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.List;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class YandexUserInfo {

    @JsonProperty("id")
    private String id;

    @JsonProperty("login")
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

    @JsonProperty("sex")
    private String sex;

    @JsonProperty("default_email")
    private String defaultEmail;

    @JsonProperty("emails")
    private List<String> emails;

    @JsonProperty("birthday")
    private String birthday;

    @JsonProperty("default_phone")
    private Phone defaultPhone;

    @JsonProperty("psuid")
    private String psuid;

    public static class Phone {

        @JsonProperty("id")
        private int id;

        @JsonProperty("number")
        private String number;

        public int getId() {
            return id;
        }

        public void setId(int id) {
            this.id = id;
        }

        public String getNumber() {
            return number;
        }

        public void setNumber(String number) {
            this.number = number;
        }
    }

}

