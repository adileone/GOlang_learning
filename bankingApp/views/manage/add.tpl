<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Add New User</h1>
     <p class="lead">Add a new User. Once it's successfuly added, you'll be redirected back to login page</p>
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">
    
    <h2>User Details</h2>
    
    <p>          
      <form role="Form" id="Users" method="POST">
        
        <div class="form-group">
          <label for="Username">Username：</label>
          <input name="Username" type="text" value="{{.User.Username}}" class="form-control" tabindex="2" />
        </div>

        <div class="form-group">
          <label for="Password">Password：</label>
          <input name="Password" type="text" value="{{.User.Password}}" class="form-control" tabindex="2" />
        </div>

        <div class="form-group">
          <label for="Email">Email：</label>
          <input name="Email" type="text" value="{{.User.Email}}" class="form-control" tabindex="2" />
        </div>
        
        <input type="submit" value="Create User" class="btn btn-default" tabindex="4"/> &nbsp;
        <a href="/home" title="don't create the user">Cancel</a>
      
      </form>
    </p>
  </div>
</div>