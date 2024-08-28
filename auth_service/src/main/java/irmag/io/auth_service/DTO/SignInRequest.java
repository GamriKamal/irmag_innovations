package irmag.io.auth_service.DTO;

import io.swagger.v3.oas.annotations.media.Schema;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Schema(description = "Sign in request")
@AllArgsConstructor
@NoArgsConstructor
public class SignInRequest {

    @Schema(description = "Username", example = "Kamal")
    @Size(min = 3, max = 50, message = "The user name must contain from 3 to 50 characters")
    @NotBlank(message = "The username cannot be empty")
    private String username;

    @Schema(description = "Password", example = "my_1secret1_password")
    @Size(max = 255, message = "The password must be no more than 255 characters long")
    @NotBlank(message = "The password cannot be empty")
    private String password;
}
