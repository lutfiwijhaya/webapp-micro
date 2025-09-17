package com.example.user_service.entity;

import jakarta.persistence.*;
import lombok.*;

import java.time.LocalDateTime;

@Entity
@Table(name = "users")
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable=false)
    private String name;

    @Column(unique = true, nullable=false)
    private String email;

    @Column(nullable=false)
    private String password; // BCrypt hashed

    @Enumerated(EnumType.STRING)
    @Column(nullable=false)
    private Role role;

    @Column(nullable=false)
    private boolean active = true;

    @Column(nullable=false)
    private boolean deleted = false; // soft delete flag

    private LocalDateTime createdAt = LocalDateTime.now();
    private LocalDateTime updatedAt = LocalDateTime.now();

    @PreUpdate
    public void preUpdate() {
        this.updatedAt = LocalDateTime.now();
    }
}
