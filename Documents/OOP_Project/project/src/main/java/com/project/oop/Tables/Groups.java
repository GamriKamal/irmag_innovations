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
public class Groups {
    @Id
    private int id;

    private int groupId, numberOfStudents, classroomId;
    private String nameOfGroup, major;

    public Groups() {

    }
}
