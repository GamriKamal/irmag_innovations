package irmag.io.apigateway.Configs;


import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateSerializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;
import irmag.io.apigateway.Filters.AuthenticationPrefilter;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.autoconfigure.jackson.Jackson2ObjectMapperBuilderCustomizer;
import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class RouteConfiguration {

    @Bean
    public RouteLocator routes(RouteLocatorBuilder builder, AuthenticationPrefilter authFilter) {
        return builder.routes()
                .route("auth-service-route", r -> r.path("/authentication-service/**")
                        .filters(f -> f
                                .rewritePath("/authentication-service/(?<segment>.*)", "/${segment}")
                                .filter(authFilter.apply(new AuthenticationPrefilter.Config())))
                        .uri("lb://AUTH_SERVICE"))
                .route("user-service-route", r -> r.path("/user-service/**")
                        .filters(f -> f
                                .rewritePath("/user-service/(?<segment>.*)", "/${segment}")
                                .filter(authFilter.apply(new AuthenticationPrefilter.Config())))
                        .uri("lb://user-service"))
                .route("product-service-route", r -> r.path("/categories/**", "/product/{id_item}", "/search", "/privacy", "/about", "/terms", "/contacts", "/send-message")
                        .filters(f -> f.stripPrefix(1))
                        .uri("lb://PRODUCT_SERVICE"))
                .route("golang-parser-service-route", r -> r.path("/go_client/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("lb://GOLANG_PARSER_SERVICE"))
                .route("default-route", r -> r.path("/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("lb://PRODUCT_SERVICE"))
                .route("auth-sign-in-route", r -> r.path("/auth/sign-in")
                        .uri("lb://AUTH_SERVICE"))
                .route("auth-sign-up-route", r -> r.path("/auth/sign-up")
                        .uri("lb://AUTH_SERVICE"))
                .build();
    }
}
