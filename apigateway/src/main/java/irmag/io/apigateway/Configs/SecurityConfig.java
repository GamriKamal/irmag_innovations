package irmag.io.apigateway.Configs;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.crypto.password.PasswordEncoder;

@Configuration
@EnableWebSecurity
@EnableGlobalMethodSecurity(prePostEnabled = true)
public class SecurityConfig extends WebSecurityConfigurerAdapter{

    @Autowired
    private PasswordEncoder encoder;

    @Value("${security.users.username}")
    private String username;

    @Value("${security.users.password}")
    private String password;

    @Autowired
    private AppUserDetailsService appUserDetailsService;


    @Override
    protected void configure(HttpSecurity http)
            throws Exception {
        http.csrf().disable()
                .headers().frameOptions().disable()
                .and()
                .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS)
                .and()
                .addFilterBefore(new JWTVerifierFilter(), UsernamePasswordAuthenticationFilter.class)
                .authorizeRequests()
                .antMatchers(HttpMethod.GET, "/api/v1/users").permitAll()
                .anyRequest()
                .authenticated()
                .and().httpBasic();

    }
}