package com.project.oop.Tables;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NonNull;
import lombok.Setter;

@Entity
@Setter
@Getter
@AllArgsConstructor
public class Classrooms {
    @Id
    private int id;
    @NonNull
    private String firstName, lastName, subject, typeOfLesson;
    private int numberOfClassroom;

    public Classrooms() {

    }

}
