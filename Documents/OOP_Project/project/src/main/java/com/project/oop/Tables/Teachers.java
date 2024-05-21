package com.project.oop.Tables;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Entity
@AllArgsConstructor
@Getter
@Setter
public class Teachers {
    @Id
    private int id;

    private String firstName, lastName, subject, typeOfLesson;

    public Teachers() {

    }
}
