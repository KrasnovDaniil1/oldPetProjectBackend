/*проверяет есть ли пользователь */
const allUsers = require('./allUsers');
const userCheck = require('./userCheck');
let numId = 0;

/*добавляет задачу */
module.exports.notesListNameAdd = function (login, password, name) {
    if (userCheck.check(login, password)) {
        numId++;
        allUsers.users[userCheck.check(login, password)].notesList.push({
            id: numId,
            name: name,
            plans: [],
        });
    } else {
        console.log('Несмогли добавить новую карточку');
    }
};

/*редактирует название задачу */
module.exports.notesListNameChange = function (
    login,
    password,
    idName,
    newName
) {
    if (userCheck.check(login, password)) {
        for (let i = 0; i < allUsers.users.notesList.length; i++) {
            if (allUsers.users.notesList[i].id == idName) {
                allUsers.users.notesListх[i].name = newName;
                console.log('Задача изменена');
            }
        }
        console.log('Такой карточку нету');
        return false;
    } else {
        console.log('Несмогли изменить эту карточку');
    }
};

/*удаляет задачу */
module.exports.notesListNameDelete = function (login, password, idName) {
    console.log('начало удалениек арточки');
    if (userCheck.check(login, password)) {
        for (
            let i = 0;
            i <
            allUsers.users[userCheck.check(login, password)].notesList.length;
            i++
        ) {
            if (
                allUsers.users[userCheck.check(login, password)].notesList[i]
                    .id == idName
            ) {
                allUsers.users[
                    userCheck.check(login, password)
                ].notesList.splice(i, 1);
                console.log('карточку удалена');
            }
        }
        console.log('Такой карточку нету');
        return false;
    } else {
        console.log('Несмогли удалить эту карточку');
    }
};

/*добавляет карточку */
module.exports.notesListTaskAdd = function (login, password, idName, taskText) {
    if (userCheck.check(login, password)) {
        numId++;
        allUsers.users[userCheck.check(login, password)].notesList[
            findNotesList(login, password, idName)
        ].plans.push({
            idTask: numId,
            text: taskText,
            checked: false,
        });
    } else {
        console.log('Несмогли добавить новую задачу');
    }
};

/*изменяет карточку */
module.exports.notesListTaskTextChange = function (
    login,
    password,
    idName,
    idTask,
    taskText
) {
    if (userCheck.check(login, password)) {
        numId++;
        allUsers.users[userCheck.check(login, password)].notesList[
            findNotesList(login, password, idName)
        ].plans[findNotesListTask(login, password, idName, idTask)].text =
            taskText;
        console.log('Название задачи изменено');
    } else {
        console.log('Несмогли изменить эту карточку');
    }
};

module.exports.notesListTaskDoneChange = function (
    login,
    password,
    idName,
    idTask,
    done
) {
    if (userCheck.check(login, password)) {
        numId++;
        allUsers.users[userCheck.check(login, password)].notesList[
            findNotesList(login, password, idName)
        ].plans[findNotesListTask(login, password, idName, idTask)].checked =
            done;
        console.log('Выполнение задачи изменено');
    } else {
        console.log('Несмогли изменить эту карточку');
    }
};

module.exports.notesListTaskDelete = function (
    login,
    password,
    idName,
    idTask
) {
    if (userCheck.check(login, password)) {
        numId++;
        allUsers.users[userCheck.check(login, password)].notesList[
            findNotesList(login, password, idName)
        ].plans.splice(findNotesListTask(login, password, idName, idTask), 1);
    } else {
        console.log('Несмогли удаллить  задачу');
    }
};

function findNotesList(login, password, idName) {
    for (
        let i = 0;
        i < allUsers.users[userCheck.check(login, password)].notesList.length;
        i++
    ) {
        if (
            allUsers.users[userCheck.check(login, password)].notesList[i].id ==
            idName
        ) {
            console.log('Карточка нашлась');
            return i;
        }
    }
    console.log('Нет такой карточки');
}
function findNotesListTask(login, password, idName, idTask) {
    for (
        let i = 0;
        i < allUsers.users[userCheck.check(login, password)].notesList.length;
        i++
    ) {
        if (
            allUsers.users[userCheck.check(login, password)].notesList[i].id ==
            idName
        ) {
            for (
                let j = 0;
                j <
                allUsers.users[userCheck.check(login, password)].notesList[i]
                    .plans.length;
                j++
            ) {
                if (
                    allUsers.users[userCheck.check(login, password)].notesList[
                        i
                    ].plans[j].id == idTask
                ) {
                    console.log('Задача нашлась');
                    return j;
                }
            }
        }
    }
    console.log('Нет такой задачи');
}
