<!DOCTYPE html>
<html lang="ko">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Snowball Chat</title>
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet"
    integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css">
  <link rel="stylesheet" as="style" crossorigin
    href="https://cdn.jsdelivr.net/gh/orioncactus/pretendard@v1.3.9/dist/web/variable/pretendardvariable.min.css" />
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.min.js"
    integrity="sha384-BBtl+eGJRgqQAUMxJ7pMwbEyER4l1g+O15P+16Ep7Q9Q+zqX6gSbd85u4mG4QzX+" crossorigin="anonymous">
  </script>
  <script src="https://code.jquery.com/jquery-3.7.1.min.js"
    integrity="sha256-/JqT3SQfawRcv/BIHPThkBvs0OEvtFFmqPF/lYI/Cxo=" crossorigin="anonymous"></script>
  <style>
  html,
  body {
    height: 100%;
    font-family: 'Pretendard Variable';
  }
  </style>
</head>

<body>
  <div class="app h-100">
    <header class="ms-3">
      <nav class="navbar bg-body-tertiary">
        <div class="container-fluid">
          <a class="navbar-brand" href="<?php echo $_SERVER['PHP_SELF'];  ?>">
            <i class="bi bi-house-door-fill"></i>
            Snowball Chat
          </a>
        </div>
      </nav>
    </header>
    <div class="container h-75">
      <div class="row h-100">
        <div class="col-md-8 m-auto text-center">
          <div class="col">
            <span class="display-4">반갑습니다. 채팅에 참여해보세요</span>
          </div>
          <div class="col-auto">
            <button type="button" class="btn btn-primary" id="goToChat">Enter!</button>
          </div>
        </div>
      </div>

    </div>
  </div>
</body>

</html>
<script type="text/javascript">
$( document ).ready( function() {
  $( "#goToChat" ).click( function() {
    $( location ).attr( "href","./chat.php" );
  } );
} );
</script>