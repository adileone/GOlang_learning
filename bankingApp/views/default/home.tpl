<header class="hero-unit">
    <div class="container">
        <div class="row">
            <div class="hero-text">
                <h1>Welcome to the Banking App!</h1>
                Try to log in!

                <br>
                <p style="color:red;">{{.flash.error}}</p>

                <p>          
                    <form id="user" method="POST">
                        name：<input name="username" id="username" type="text" /> </br>
                        password：<input name="password" id="password" type="password" /></br>
                        <input type="submit" value="Sign in" class="btn btn-default" tabindex="4"/> &nbsp;
                    </form>
                </p>
                
                </br>
                if you are not already registered please, do it!
                <a href="/manage/add" title="Sign up">Join us!</a>
      
            </div>
        </div>
    </div>
</header>