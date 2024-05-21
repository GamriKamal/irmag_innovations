package com.project.oop.Controllers;

import com.project.oop.Services.GroupsService;
import com.project.oop.Tables.Groups;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/groups")
public class GroupsController {
    @Autowired
    GroupsService service;

    @GetMapping
    public List<Groups> getAllGroups(){
        return service.getAllGroups();
    }

    @GetMapping("/{id}")
    public Groups getGroupById(@PathVariable int id){
        return service.getGroupById(id);
    }

    @PostMapping
    public Groups addGroups(@RequestBody Groups group){
        return service.addGroup(group);
    }

    @PutMapping("/{id}")
    public Groups updateGroup(@PathVariable int id, @RequestBody Groups groups){
        return service.updateGroupId(id, groups);
    }

    @DeleteMapping("/{id}")
    public void deleteGroupById(@PathVariable int id){
        service.deleteGroupById(id);
    }

}
