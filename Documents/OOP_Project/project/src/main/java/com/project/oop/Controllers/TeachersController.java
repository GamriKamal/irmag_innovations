package com.project.oop.Controllers;

import com.project.oop.Services.GroupsService;
import com.project.oop.Services.TeachersService;
import com.project.oop.Tables.Groups;
import com.project.oop.Tables.Teachers;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Controller
@RequestMapping("/teachers")
public class TeachersController {
    @Autowired
    TeachersService service;

    @GetMapping
    public List<Teachers> getAllTeachers(){
        return service.getAllTeachers();
    }

    @GetMapping("/{id}")
    public Teachers getTeacherById(@PathVariable int id){
        return service.getTeacherById(id);
    }

    @PostMapping
    public Teachers addTeacher(@RequestBody Teachers teachers){
        return service.addTeacher(teachers);
    }

    @PutMapping("/{id}")
    public Teachers updateTeacher(@PathVariable int id, @RequestBody Teachers teachers){
        return service.updateTeacherId(id, teachers);
    }

    @DeleteMapping("/{id}")
    public void deleteTeacherById(@PathVariable int id){
        service.deleteTeacherById(id);
    }
}
