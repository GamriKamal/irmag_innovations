package com.project.oop.Services;

import com.project.oop.Repositories.GroupsRepository;
import com.project.oop.Tables.Groups;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class GroupsService {
    @Autowired
    private GroupsRepository repo;

    public List<Groups> getAllGroups(){
        return repo.findAll();
    }

    public Groups getGroupById(int id){
        Optional<Groups> tempGroup = repo.findById(id);

        if(tempGroup.isEmpty()){
            throw new RuntimeException("Group with id {" + id + "} not found");
        }
        return tempGroup.get();
    }

    public Groups addGroup(Groups group){
        return repo.save(group);
    }

    public Groups updateGroupId(int id, Groups group){
        Optional<Groups> tempGroup = repo.findById(id);

        if(tempGroup.isEmpty()){
            throw new RuntimeException("Group with id {" + id + "} not found");
        }
        group.setGroupId(id);
        return repo.save(group);
    }

    public void deleteGroupById(int id){
        Optional<Groups> tempGroup = repo.findById(id);

        if(tempGroup.isEmpty()){
            throw new RuntimeException("Group with id {" + id + "} not found");
        }
        repo.deleteById(id);
    }
}
