document.addEventListener('DOMContentLoaded', function() {
  if (document.URL.match(/edit/)) {
    var status = document.getElementById("status").innerHTML;
    var todoButton = document.getElementById("status-todo");
    var progressButton = document.getElementById("status-progress");
    var doneButton = document.getElementsByClassName("status")

    if (status == 0) {
      todoButton.classList.remove('btn-outline-warning');
      todoButton.classList.add('btn-warning');
    } else if (status == 1) {
      progressButton.classList.remove('btn-outline-info');
      progressButton.classList.add('btn-info');
    } else if (status == 2) {
      doneButton.classList.remove('btn-outline-success');
      doneButton.classList.add('btn-success');
    }
  }
  
  // status sort
  if (document.URL.match("/todos/todo/0")){
    var sortTodo = document.getElementById("sort-btn-todo")
    sortTodo.classList.remove('btn-outline-secondary');
    sortTodo.classList.remove('text-secondary');
    sortTodo.classList.add('btn-warning');
    sortTodo.classList.add('text-dark');
  } else if (document.URL.match("/todos/inprogress/1")){
    var sortProgress = document.getElementById("sort-btn-in-progress")
    sortProgress.classList.remove('btn-outline-secondary');
    sortProgress.classList.remove('text-secondary');
    sortProgress.classList.add('btn-info');
    sortProgress.classList.add('text-white');
  } else if (document.URL.match("/todos/done/2")){
    var sortDone = document.getElementById("sort-btn-done")
    sortDone.classList.remove('btn-outline-secondary');
    sortDone.classList.remove('text-secondary');
    sortDone.classList.add('btn-success');
    sortDone.classList.add('text-white');
  } else if (document.URL.match("/todos")){
    var sortAll = document.getElementById("sort-btn-all")
    sortAll.classList.remove('btn-outline-secondary');
    sortAll.classList.remove('text-secondary');
    sortAll.classList.add('btn-secondary');
    sortAll.classList.add('text-light');
  }

  if (document.URL.match("/todos")){
    var todoIds = document.getElementsByClassName("todo-id");
    var statusEls = document.getElementsByClassName("status");
    var todoButtons = document.getElementsByClassName("status-todo");
    var progressButtons = document.getElementsByClassName("status-progress");
    var doneButtons = document.getElementsByClassName("status-done");

    for (var i = 0; i < todoIds.length; i++) {
      var status = statusEls[i].innerHTML;
      var todoButton = todoButtons[i];
      var progressButton = progressButtons[i];
      var doneButton = doneButtons[i];

      if (status == 0) {
        todoButton.classList.remove('btn-outline-warning');
        todoButton.classList.add('btn-warning');
      } else if (status == 1) {
        progressButton.classList.remove('btn-outline-info');
        progressButton.classList.add('btn-info');
      } else if (status == 2) {
        doneButton.classList.remove('btn-outline-success');
        doneButton.classList.add('btn-success');
      }
    }
  }

  // if (document.URL.match(/edit/) || document.URL.match(/new/)){
  //   todoButton.addEventListener('click', function() {
  //     submitStatus("0"); // ボタンがクリックされたらパラメータとして0を送信
  //   });

  //   progressButton.addEventListener('click', function() {
  //     submitStatus("1");
  //   });

  //   doneButton.addEventListener('click', function() {
  //     submitStatus("2");
  //   });

  //   function submitStatus(status) {
  //     // パラメータを含めてリクエストを送信
  //     var xhr = new XMLHttpRequest();
  //     xhr.open('POST', `/todos/update/${todoId}`, true); // リクエスト先のURLを適切に指定する
  //     var formData = new FormData();
  //     formData.append("status", status);
  //     xhr.send(formData);
  //   }
  // };
});
