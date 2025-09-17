package com.example.user_service.dto;

import lombok.*;

@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class TokenRequest {
    private String token;
}
