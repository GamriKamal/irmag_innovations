package com.project.oop.Repositories;

import com.project.oop.Tables.Teachers;
import org.hibernate.id.IntegralDataTypeHolder;
import org.springframework.data.jpa.repository.JpaRepository;

public interface TeachersRepository extends JpaRepository<Teachers, Integer> {
}
