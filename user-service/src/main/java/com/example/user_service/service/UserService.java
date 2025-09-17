package com.example.user_service.service;

import com.example.user_service.dto.UserDto;
import com.example.user_service.entity.User;
import com.example.user_service.repository.UserRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class UserService {
  private final UserRepository repo;

  public UserService(UserRepository repo) {
    this.repo = repo;
  }

  public List<UserDto> getAll() {
    return repo.findAllByDeletedFalse().stream().map(this::toDto).collect(Collectors.toList());
  }

  public UserDto getById(Long id) {
    User u = repo.findByIdAndDeletedFalse(id).orElseThrow(() -> new RuntimeException("User not found"));
    return toDto(u);
  }

  public UserDto softDelete(Long id) {
    User u = repo.findByIdAndDeletedFalse(id).orElseThrow(() -> new RuntimeException("User not found"));
    u.setDeleted(true);
    u.setActive(false);
    repo.save(u);
    return toDto(u);
  }

  public UserDto update(Long id, UserDto dto) {
    User u = repo.findByIdAndDeletedFalse(id).orElseThrow(() -> new RuntimeException("User not found"));
    u.setName(dto.getName());
    u.setActive(dto.isActive());
    repo.save(u);
    return toDto(u);
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

  public UserDto findByEmail(String email) {
    User user = repo.findByEmailAndDeletedFalse(email)
        .orElseThrow(() -> new RuntimeException("User not found"));
    return toDto(user);
  }

}
