function addelement(s) {
    var elem1 = $('<li><span></span> <button class="delete-button">Удалить</button></li>');
    $('span', elem1).text(s);
    $('#root ul').append(elem1);
    $('.delete-button', elem1).click(function(ev){$(this).parent().remove()});
}
$(function(){
    $('#root').append('<ul></ul>');
    $('#root').append('<input id="add_task_input"> <button id="add_task">Добавить</button>');
    $('#add_task').click(function(){addelement($('#add_task_input').val())});
    addelement('Сделать задание #3 по web-программированию');
}) 