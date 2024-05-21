package com.project.oop.Services;

import com.project.oop.Repositories.ClassroomsRepository;
import com.project.oop.Tables.Classrooms;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;


@Service
public class ClassroomsService {
    @Autowired
    private ClassroomsRepository repo;

    public List<Classrooms> getAllClassrooms(){
        return repo.findAll();
    }

    public Classrooms getClassroomById(int id){
        Optional<Classrooms> tempRoom = repo.findById(id);

        if(tempRoom.isEmpty()){
            throw new RuntimeException("Classroom with id {" + id + "} not found");
        }
        return tempRoom.get();
    }

    public Classrooms addClassroom(Classrooms classroom){
        return repo.save(classroom);
    }


    public void deleteClassroomById(int id){
        Optional<Classrooms> tempClassroom = repo.findById(id);

        if(tempClassroom.isEmpty()){
            throw new RuntimeException("Classroom with id {" + id + "} not found");
        }
        repo.deleteById(id);
    }
}
