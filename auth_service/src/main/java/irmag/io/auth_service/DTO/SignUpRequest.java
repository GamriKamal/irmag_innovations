package irmag.io.auth_service.DTO;

import io.swagger.v3.oas.annotations.media.Schema;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.Data;
import lombok.ToString;

@Data
@Schema(description = "Registration request")
@ToString
public class SignUpRequest {

    @Schema(description = "Username", example = "Kamal")
    @Size(min = 3, max = 50, message = "The user name must contain from 3 to 50 characters")
    @NotBlank(message = "The username cannot be empty")
    private String username;

    @Schema(description = "Email", example = "mrirmag.pr@gmail.com")
    @Size(min = 5, max = 255, message = "The email address must contain from 5 to 255 characters")
    @NotBlank(message = "The email address cannot be empty")
    @Email(message = "The email address must be in the format user@example.com")
    private String email;

    @Schema(description = "Password", example = "my_1secret1_password")
    @Size(max = 255, message = "The password must be no more than 255 characters long")
    @NotBlank(message = "The password cannot be empty")
    private String password;
}
