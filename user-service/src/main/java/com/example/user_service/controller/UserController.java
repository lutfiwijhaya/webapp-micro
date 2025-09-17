package com.example.user_service.controller;

import com.example.user_service.dto.UserDto;
import com.example.user_service.service.UserService;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.*;
import java.util.List;

@RestController
@RequestMapping("/api/v1/users")
public class UserController {
    private final UserService svc;
    public UserController(UserService svc) { this.svc = svc; }

    // ================= ADMIN =================
    @GetMapping
    @PreAuthorize("hasAuthority('ADMIN')")
    public List<UserDto> getAll() {
        return svc.getAll();
    }

    @PutMapping("/{id}")
    @PreAuthorize("hasAuthority('ADMIN')")
    public UserDto update(@PathVariable Long id, @RequestBody UserDto dto) {
        return svc.update(id, dto);
    }

    @DeleteMapping("/{id}")
    @PreAuthorize("hasAuthority('ADMIN')")
    public UserDto softDelete(@PathVariable Long id) {
        return svc.softDelete(id);
    }

    // ================= MANAGER =================
    @GetMapping("/manager-view")
    @PreAuthorize("hasAuthority('MANAGER')")
    public List<UserDto> managerViewAll() {
        return svc.getAll();
    }

    // ================= EMPLOYEE =================
    @GetMapping("/me")
    @PreAuthorize("hasAuthority('EMPLOYEE')")
    public UserDto getMyProfile(@AuthenticationPrincipal UserDetails userDetails) {
        return svc.findByEmail(userDetails.getUsername());
    }

    @GetMapping("/{id}")
    @PreAuthorize("hasAnyAuthority('ADMIN','MANAGER')")
    public UserDto getById(@PathVariable Long id) {
        return svc.getById(id);
    }
}
