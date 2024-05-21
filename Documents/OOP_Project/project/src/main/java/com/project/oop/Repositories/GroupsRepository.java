package com.project.oop.Repositories;

import com.project.oop.Tables.Groups;
import org.springframework.data.jpa.repository.JpaRepository;

public interface GroupsRepository extends JpaRepository<Groups, Integer> {
}
