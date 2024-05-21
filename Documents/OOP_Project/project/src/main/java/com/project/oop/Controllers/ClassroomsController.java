package com.project.oop.Controllers;

import com.project.oop.Services.ClassroomsService;
import com.project.oop.Services.TeachersService;
import com.project.oop.Tables.Classrooms;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Controller
public class ClassroomsController {
    @Autowired
    ClassroomsService service;

    @GetMapping
    public List<Classrooms> getAllClassrooms(){
        return service.getAllClassrooms();
    }

    @GetMapping("/{id}")
    public Classrooms getClassroomById(@PathVariable int id){
        return service.getClassroomById(id);
    }

    @PostMapping
    public Classrooms addClassroom(@RequestBody Classrooms classroom){
        return service.addClassroom(classroom);
    }

    @DeleteMapping("/{id}")
    public void deleteClassroomById(@PathVariable int id){
        service.deleteClassroomById(id);
    }
}
