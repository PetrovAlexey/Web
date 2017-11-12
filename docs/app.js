$(document).on('click','.delete', function(){
    $(this).parent().remove();
});

$(document).on('click','#add_task', function(){
    $('#root').append('<ul><li><span>'+$('#add_task_input').val()+'</span><button class="delete">Удалить</button></li></ul>');
});

$(document).ready(function () {
    $('#root').append('<ul><li><span>Сделать задание #3 по web-программированию</span><button class="delete">Удалить</button></li></ul>');
    $('#root').prepend('<input id = "add_task_input"></input> <button id = "add_task">Добавить</button>');
});

