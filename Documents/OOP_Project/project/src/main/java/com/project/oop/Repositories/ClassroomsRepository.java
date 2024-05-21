package com.project.oop.Repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import com.project.oop.Tables.Classrooms;

public interface ClassroomsRepository extends JpaRepository<Classrooms, Integer> {
}
