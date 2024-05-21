package com.project.oop.Services;

import com.project.oop.Repositories.GroupsRepository;
import com.project.oop.Repositories.TeachersRepository;
import com.project.oop.Tables.Groups;
import com.project.oop.Tables.Teachers;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class TeachersService {
    @Autowired
    private TeachersRepository repo;

    public List<Teachers> getAllTeachers(){
        return repo.findAll();
    }

    public Teachers getTeacherById(int id){
        Optional<Teachers> tempTeacher = repo.findById(id);

        if(tempTeacher.isEmpty()){
            throw new RuntimeException("Teacher with id {" + id + "} not found");
        }
        return tempTeacher.get();
    }

    public Teachers addTeacher(Teachers teacher){
        return repo.save(teacher);
    }

    public Teachers updateTeacherId(int id, Teachers teacher){
        Optional<Teachers> tempTeacher = repo.findById(id);

        if(tempTeacher.isEmpty()){
            throw new RuntimeException("Teacher with id {" + id + "} not found");
        }
        teacher.setId(id);
        return repo.save(teacher);
    }

    public void deleteTeacherById(int id){
        Optional<Teachers> tempTeacher = repo.findById(id);

        if(tempTeacher.isEmpty()){
            throw new RuntimeException("Teacher with id {" + id + "} not found");
        }
        repo.deleteById(id);
    }

}
