package com.example.user_service.repository;

import com.example.user_service.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;
import java.util.List;

public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByEmailAndDeletedFalse(String email);
    Optional<User> findByIdAndDeletedFalse(Long id);
    List<User> findAllByDeletedFalse();
    boolean existsByEmailAndDeletedFalse(String email);
}
