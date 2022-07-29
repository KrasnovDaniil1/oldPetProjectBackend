const http = require('http');
const allUsers = require('./allUsers');
const userCheck = require('./userCheck');
const signUser = require('./signUser');
const notesList = require('./notesList');

http.createServer(function (request, response) {
    let data = [];
    request.on('data', (chunk) => {
        data.push(chunk);
    });
    request.on('end', () => {
        data = JSON.parse(data);
        console.log(data);

        if (request.method == 'DELETE') {
            switch (request.url) {
                case '/api/notesListNameDelete':
                    notesList.notesListNameDelete(
                        data.login,
                        data.password,
                        data.idName
                    );
                    break;
                case '/api/notesListTaskDelete':
                    notesList.notesListTaskDelete(
                        data.login,
                        data.password,
                        data.idName,
                        data.idTask
                    );
                    break;
                default:
                    console.log('нет обработчика этого запроса');
            }
        } else if (request.method == 'PUT') {
            switch (request.url) {
                case '/api/notesListNameChange':
                    notesList.notesListNameChange(
                        data.login,
                        data.password,
                        data.idName,
                        data.newName
                    );
                    break;
                case '/api/notesListTaskTextChange':
                    notesList.notesListTaskTextChange(
                        data.login,
                        data.password,
                        data.idName,
                        data.idTask,
                        data.taskText
                    );
                    break;
                case '/api/notesListTaskDoneChange':
                    notesList.notesListTaskDoneChange(
                        data.login,
                        data.password,
                        data.idName,
                        data.idTask,
                        data.done
                    );
                    break;
                default:
                    console.log('нет обработчика этого запроса');
            }
        } else if (request.method == 'POST') {
            switch (request.url) {
                case '/api/login':
                    /*авторизация */
                    break;
                case '/api/sign':
                    signUser.createSignUser(data.login, data.password);
                    break;

                case '/api/notesListNameAdd':
                    notesList.notesListNameAdd(
                        data.login,
                        data.password,
                        data.name
                    );
                    break;
                case '/api/notesListTaskAdd':
                    notesList.notesListTaskAdd(
                        data.login,
                        data.password,
                        data.idName,
                        data.taskText
                    );
                    break;

                // case '/api/diaryListNameAdd':
                //     signUser.createSignUser(data.login, data.password);
                //     response.end(
                //         JSON.stringify(
                //             allUsers.users[
                //                 userCheck.check(data.login, data.password)
                //             ]
                //         )
                //     );
                //     break;

                default:
                    console.log('нет обработчика этого запроса');
            }
        }
        console.log(allUsers.users);
        response.end(
            JSON.stringify(
                allUsers.users[userCheck.check(data.login, data.password)]
            )
        );
    });
}).listen(4000);
