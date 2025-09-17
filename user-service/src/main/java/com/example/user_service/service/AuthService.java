package com.example.user_service.service;

import com.example.user_service.dto.LoginRequest;
import com.example.user_service.dto.RegisterRequest;
import com.example.user_service.dto.UserDto;
import com.example.user_service.entity.User;
import com.example.user_service.repository.UserRepository;
import com.example.user_service.security.JwtUtil;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jws;

import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;
import java.util.*;

@Service
public class AuthService {

    private final UserRepository userRepository;
    private final JwtUtil jwtUtil;
    private final AuthenticationManager authenticationManager;
    private final BCryptPasswordEncoder passwordEncoder = new BCryptPasswordEncoder();

    public AuthService(UserRepository userRepository, JwtUtil jwtUtil, AuthenticationManager authenticationManager) {
        this.userRepository = userRepository;
        this.jwtUtil = jwtUtil;
        this.authenticationManager = authenticationManager;
    }

    public UserDto register(RegisterRequest req) {
        if (userRepository.existsByEmailAndDeletedFalse(req.getEmail())) {
            throw new RuntimeException("Email sudah terdaftar");
        }
        User user = User.builder()
                .name(req.getName())
                .email(req.getEmail())
                .password(passwordEncoder.encode(req.getPassword()))
                .role(req.getRole())
                .active(true)
                .deleted(false)
                .build();
        return toDto(userRepository.save(user));
    }

    public Map<String, Object> login(LoginRequest req) {
        try {
            // cek ke Spring Security
            authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(req.getEmail(), req.getPassword())
            );
        } catch (Exception e) {
            throw new RuntimeException("Email/password salah");
        }

        User user = userRepository.findByEmailAndDeletedFalse(req.getEmail())
                .orElseThrow(() -> new RuntimeException("Email/password salah"));

        Map<String, Object> claims = new HashMap<>();
        claims.put("id", user.getId());
        claims.put("name", user.getName());
        claims.put("roles", List.of(user.getRole().name()));

        String token = jwtUtil.generateToken(claims, user.getEmail());

        return Map.of(
            "token", token,
            "user", toDto(user)
        );
    }

    public UserDto validateToken(String token) {
        try {
            Jws<Claims> claimsJws = jwtUtil.validateToken(token);
            Long id = claimsJws.getBody().get("id", Long.class);
            return userRepository.findByIdAndDeletedFalse(id)
                    .map(this::toDto)
                    .orElseThrow(() -> new RuntimeException("User not found"));
        } catch (Exception ex) {
            throw new RuntimeException("Token invalid: " + ex.getMessage());
        }
    }

    private UserDto toDto(User u) {
        return UserDto.builder()
                .id(u.getId())
                .name(u.getName())
                .email(u.getEmail())
                .role(u.getRole())
                .active(u.isActive())
                .build();
    }
}

