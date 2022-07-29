/*проверяет есть ли пользователь */
const allUsers = require('./allUsers');

module.exports.check = function (login, password) {
    for (let i = 0; i < allUsers.users.length; i++) {
        if (
            allUsers.users[i].login == login &&
            allUsers.users[i].password == password
        ) {
            console.log('Пользователь найден');
            // return allUsers.users[i];
            return i;
        }
    }
    console.log('Ползователь не найден');
    return false;
};
