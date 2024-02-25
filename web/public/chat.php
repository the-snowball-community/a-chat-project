<?php
$chatLine = '
  <div class="d-flex">
    <div class="flex-shrink-0">
      <img class="rounded-circle" src="https://via.placeholder.com/50/FA6F57/fff&text=ME" alt="...">
    </div>
    <div class="col-md-9 ms-3">
      <p class="fw-bold mb-0" id="userName">asdff</p>
      <p>
      This is some content from a media component. You can replace this with any content and adjust it as needed.
      </p>
    </div>
  </div>
';
$chatLineMe = '';
$chatLineMe .= '<div class="d-flex flex-row-reverse">';
$chatLineMe .= '<div class="flex-shrink-0">';
$chatLineMe .= '<img class="rounded-circle" src="https://via.placeholder.com/50/4287f5/fff&text=En" alt="...">';
$chatLineMe .= '</div>';
$chatLineMe .= '<div class="col-md-9 me-3">';
$chatLineMe .= '<p class="text-end fw-bold mb-0" id="userName">ddddddddd</p>';
$chatLineMe .= '<p>This is some content from a media component. You can replace this with any content and adjust it as needed.</p>';
$chatLineMe .= '</div>';
$chatLineMe .= '</div>';
?>
<!DOCTYPE html>
<html lang="ko">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Snowball Chat</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.min.js" integrity="sha384-BBtl+eGJRgqQAUMxJ7pMwbEyER4l1g+O15P+16Ep7Q9Q+zqX6gSbd85u4mG4QzX+" crossorigin="anonymous"></script>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css">
  <link rel="stylesheet" as="style" crossorigin
    href="https://cdn.jsdelivr.net/gh/orioncactus/pretendard@v1.3.9/dist/web/variable/pretendardvariable.min.css" />
  </script>
  <script src="https://code.jquery.com/jquery-3.7.1.min.js"
    integrity="sha256-/JqT3SQfawRcv/BIHPThkBvs0OEvtFFmqPF/lYI/Cxo=" crossorigin="anonymous"></script>
  <style>
  /* html,
  body {
    font-family: 'Pretendard Variable';
  } */
  </style>
</head>

<body>
  <div id="content"></div>
  <div class="app">
    <header class="mx-3 fixed-top">
      <nav class="navbar  bg-body-tertiary">
        <div class="container-fluid">
          <a class="navbar-brand" href="/public/index.php">
            <i class="bi bi-house-door-fill"></i>
            Snowball Chat
          </a>
          <p class="navbar-text my-0" id="userName"> 님</p>
        </div>
      </nav>
    </header>
    <!-- .container-fluid, which is width: 100% at all breakpoints -->
    <div class="container-fluid py-5" style="background-color: gray;">
      <div class="row py-5">
        <!-- chat contain -->
        <div class="col-md-8 mx-auto">
          <!-- left, -->
          <?php for($i=0; $i<3; ++$i) echo $chatLine ?>
            
          <!-- right, -->
          <?php for($i=0; $i<3; ++$i) echo $chatLineMe ?>
          <?php for($i=0; $i<3; ++$i) echo $chatLine ?>
        </div>
      </div>
        
      <!-- write message -->
      <div class="col-md-8 mx-auto py-5 fixed-bottom">
        <form method="POST" id="contentForm" autocomplete="off">
          <input type="hidden" name="userId" value="1234">
          <input type="hidden" name="userName" value="무뚝뚝한여우">
          <input type="hidden" name="roomId" value="102939">
          <label for="content" class="visually-hidden">채팅을 입력하세요.</label>
          <div class="input-group">
            <input type="text" class="form-control" name="content" placeholder="채팅을 입력하세요." aria-label="채팅을 입력하세요." aria-describedby="sendMessage">
            <div class="input-group-append">
              <button class="btn btn-outline-secondary" id="sendMessage" type="submit" >전송</button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
<script>
$(document).ready(function() {
  getRandUser();
  getMessages();
  // alert('반갑습니다. 채팅방에 참여하였습니다.');

  $("form").on("submit", function(event) {
    sendMessage();

    event.preventDefault();
  });
});

const getRandUser = () => {
  console.log("get Rand Username");
  $.ajax({
    type: "GET",
    url: "",
    dataType: "json",
    timeout: 5000
  })
  .done(function(result) {
    console.dir(result);
    $("#userName").prepend(result.name);
  })
  .fail(function(jqXHR, textStatus) {
    console.log(textStatus);
  })
}
// const sendMessage = async() => {
//   console.log('start send message');
//   const contents = $("#contentForm").serializeArray();
//   let param = {};

//   contents.map(function(data,index){
//     param[data.name] = data.value;
//   })
//   console.dir(param);
//   let url = "";
//   const response = await fetch(url, {
//     method: "POST", // *GET, POST, PUT, DELETE, etc.
//     headers: {
//       "Content-Type": "application/json",
//       // 'Content-Type': 'application/x-www-form-urlencoded',
//     },
//     body: JSON.stringify( {
//       "userId": param.userId,
//       "roomId": param.roomId,
//       "userName": param.userName,
//       "content": param.content
//     }) // body data type must match "Content-Type" header
//   });
//   let test = response.json();
//   console.log(test);
// }


const sendMessage = () => {
  console.log('start send message');
  const contents = $("#contentForm").serializeArray();
  let param = {};

  contents.map(function(data,index){
    param[data.name] = data.value;
  })
  console.dir(param);
  $.ajax({
    type: "POST",
    url: "",
    dataType: "json",
    contentType: 'application/json',
    timeout: 5000,
    data:JSON.stringify( {
      "userId": param.userId,
      "roomId": param.roomId,
      "userName": param.userName,
      "content": param.content
    })
  })
  .done(function(result) {
    console.log(result);
  })
  .fail(function(jqXHR, textStatus) {
    console.log(textStatus);
  })
}

const getMessages = () => {
  console.log('get start');
  $.ajax({
    type: "GET",
    url: "",
    dataType: "json",
    timeout: 5000
  })
  .done(function(result) {
    console.dir(result);
    // $("#userName").html(result);
    // console.log(result);
  })
  .fail(function(jqXHR, textStatus) {
    console.log(textStatus);
  })
}
</script>
</body>
</html>
