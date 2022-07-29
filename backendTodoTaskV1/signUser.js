const allUsers = require('./allUsers');
const userCheck = require('./userCheck');

module.exports.createSignUser = function (login, password) {
    if (!userCheck.check(login, password)) {
        allUsers.users.push({
            login: login,
            password: password,
            notesList: [],
            diaryList: [],
        });
        console.log('Пользователь создан');
    } else {
        console.log('Пользователь с таким именем уже есть');
        return false;
    }
};
