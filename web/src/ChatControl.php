<?php
namespace Src\ChatControl;
class Chat
{

  private $content;
  function __construct()
  {
    $this->content = $_POST['chatContent'];

    echo $this->content;
    echo 'Hello World';
  }

}
new Chat();

