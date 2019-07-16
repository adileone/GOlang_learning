<html>

<head>
<script src="http://cdnjs.cloudflare.com/ajax/libs/jquery/2.0.3/jquery.min.js"></script>
<script src="http://www.jqueryscript.net/demo/jQuery-Plugin-To-Convert-HTML-Table-To-CSV-tabletoCSV/jquery.tabletoCSV.js"></script>
</head>

<body>

<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>All users</h1>
     <p class="lead">Here you'll find all the available users</p>
    </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">
    <table class="table">
      <thead>
        <tr>
          <th>Id</th>
          <th>Username</th>
          <th>Email</th>
          <th>Balance</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td>{{$record.Id}}</td>
          <td>{{$record.Username}}</td>
          <td>{{$record.Email}}</td>
          <td>{{$record.Balance}}</td>
          </tr>
        {{end}}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="4"><a href="{{urlfor "ManageController.Add"}}" title="add new user">add new user</a></td>
        </tr>
      </tfoot>
       <tr>
           <p>          
              <form id="del" method="POST">
                  Id：<input name="token" id="token" type="text" /> 
                  <input type="submit" value="Delete user" class="btn btn-default" tabindex="4" />
              </form>
            </p>
          </tr>
    </table>
    <a href="{{urlfor "ManageController.Home"}}" title="back">back</a>
  </div>
</div>

<button id="export" data-export="export">Export</button>
    
<script>
$("#export").click(function(){
      $("table").tableToCSV();
    });
</script>
</body>

</html>