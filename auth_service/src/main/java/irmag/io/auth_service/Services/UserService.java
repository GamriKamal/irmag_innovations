package irmag.io.auth_service.Services;

import irmag.io.auth_service.Exceptions.EmailAlreadyExistsException;
import irmag.io.auth_service.Exceptions.UserAlreadyExistsException;
import irmag.io.auth_service.Models.Role;
import irmag.io.auth_service.Models.User;
import irmag.io.auth_service.Repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

@Service
public class UserService {
    private UserRepository userRepository;

    @Autowired
    public void setUserRepository(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public User save(User user) {
        return userRepository.save(user);
    }

    public User create(User user) {
        System.out.println(user.getUsername());
        if (userRepository.existsByUsername(user.getUsername())) {
            throw new UserAlreadyExistsException("User with that name already exists");
        }

        if (userRepository.existsByEmail(user.getEmail())) {
            throw new EmailAlreadyExistsException("User with this email already exists");
        }

        return save(user);
    }

    public User getByUsername(String username) {
        return userRepository.findByUsername(username)
                .orElseThrow(() -> new UsernameNotFoundException("User not found"));

    }

    public User findByEmail(String email) {
        return userRepository.findByEmail(email).orElse(null);
    }

    public UserDetailsService userDetailsService() {
        return this::getByUsername;
    }

    public User getCurrentUser() {
        var username = SecurityContextHolder.getContext().getAuthentication().getName();
        return getByUsername(username);
    }

    public void getAdmin() {
        var user = getCurrentUser();
        user.setRole(Role.ROLE_ADMIN);
        save(user);
    }
}
