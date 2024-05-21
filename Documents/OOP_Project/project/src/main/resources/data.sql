CREATE TABLE TEACHERS (
    id INT NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    course_type VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE GROUPS (
    id INT NOT NULL AUTO_INCREMENT,
    groupId INT NOT NULL,
    nameOfGroup VARCHAR(255) NOT NULL,
    major VARCHAR(255) NOT NULL,
    numberOfStudents INT NOT NULL,
    classroomId INT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE CLASSROOMS (
    id INT NOT NULL AUTO_INCREMENT,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    typeOfLesson VARCHAR(255) NOT NULL,
    numberOfClassroom INT NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO TEACHERS VALUES (DEFAULT, 'John', 'Doe', 'English', 'Lecture');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Jane', 'Doe', 'History', 'Practise');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Peter', 'Parker', 'Physics', 'Lecture');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Mary', 'Jane', 'Chemistry', 'Practise');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Bruce', 'Wayne', 'Biology', 'Lecture');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Clark', 'Kent', 'Geography', 'Practise');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Tony', 'Stark', 'Computer Science', 'Lecture');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Natasha', 'Romanoff', 'Foreign Language', 'Practise');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Steve', 'Rogers', 'Physical Education', 'Lecture');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Wanda', 'Maximoff', 'Arts and Crafts', 'Practise');
INSERT INTO TEACHERS VALUES (DEFAULT, 'Vision', 'AI', 'Ethics', 'Lecture');

INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (1, 'Group 1', 'Computer Science', 20, 101);
INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (2, 'Group 2', 'Mathematics', 25, 102);
INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (3, 'Group 3', 'Physics', 30, 103);
INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (4, 'Group 4', 'Chemistry', 35, 104);
INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (5, 'Group 5', 'Biology', 40, 105);
INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (6, 'Group 6', 'History', 45, 106);
INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (7, 'Group 7', 'English', 50, 107);
INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (8, 'Group 8', 'Foreign Language', 55, 108);
INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (9, 'Group 9', 'Arts and Crafts', 60, 109);
INSERT INTO GROUPS (groupId, nameOfGroup, major, numberOfStudents, classroomId)
VALUES (10, 'Group 10', 'Physical Education', 65, 110);


INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Alice', 'Smith', 'Math', 'Lecture', 1);
INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Bob', 'Jones', 'Science', 'Discussion', 2);
INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Carol', 'Williams', 'English', 'Lab', 3);
INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Dave', 'Brown', 'History', 'Tutorial', 4);
INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Eve', 'Johnson', 'Art', 'Project', 5);

INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Alice', 'Jones', 'Math', 'Discussion', 6);
INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Bob', 'Williams', 'Science', 'Lab', 7);
INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Carol', 'Brown', 'English', 'Tutorial', 8);
INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Dave', 'Johnson', 'History', 'Project', 9);
INSERT INTO CLASSROOMS (firstName, lastName, subject, typeOfLesson, numberOfClassroom) VALUES ('Eve', 'Smith', 'Art', 'Lecture', 10);
