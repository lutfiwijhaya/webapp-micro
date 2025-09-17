package com.example.user_service.dto;

import com.example.user_service.entity.Role;
import lombok.*;

@Getter @Setter @NoArgsConstructor @AllArgsConstructor @Builder
public class UserDto {
    private Long id;
    private String name;
    private String email;
    private Role role;
    private boolean active;
}
