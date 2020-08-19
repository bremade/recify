<script>
  import Card from '@smui/card';
  import Textfield from '@smui/textfield'
  import Button from '@smui/button';
  import Snackbar, { Label } from '@smui/snackbar';

  var username = "";
  var password = "";
  var repeatedPassword = "";
  var register = false;

  var registerErrorSnackbar;
  var registerErrorSnackbarText;

  function doRegister() {
    if (!register) {
      register = true;
    } else {
      if (password !== repeatedPassword) {
        snackbar('The passwords do not match.');
      } else {
        registerRequest();
      }
    }
  }

  function login() {
    fetch('/api/v1/auth/login', {
      method: 'POST',
      body: JSON.stringify({
        username, password
      })
    }).then(result => {
      if (result.status === 200) {
        // We have our session cookie!
        // Redirect to homepage
        window.location.pathname = '/';
      } else {
        snackbar('Wrong username or password.');
      }
    });
  }

  function registerRequest() {
    fetch('/api/v1/auth/register', {
      method: 'POST',
      body: JSON.stringify({
        username, password
      })
    }).then(result => {
        if (result.status === 200) {
          register = false;
          snackbar('Registration successful.');
        } else if (result.status === 409) {
          username = '';
          snackbar('Username was already taken.');
        } else {
          snackbar('Unknown error.');
        }
      });
  }

  function snackbar(message) {
    registerErrorSnackbarText = message;
    registerErrorSnackbar.open();
  }
</script>

<div class="login-container">
  <div class="login">
    <Card variant="outlined">
      <div class="title">
        Login
      </div>
      <div class="input-container">
        <Textfield class="textfield" label="Username" bind:value={username}></Textfield>
        <Textfield type="password" class="textfield" label="Password" bind:value={password}></Textfield>
        {#if register}
          <Textfield type="password" class="textfield" label="Repeat password" bind:value={repeatedPassword} invalid={password !== repeatedPassword}></Textfield>
         {/if}
        <Button class="login-button" on:click={login} variant="unelevated">Login</Button>
        <Button class="register-button" on:click={doRegister}>Register</Button>
      </div>
    </Card>
  </div>
  <Snackbar bind:this={registerErrorSnackbar} labelText={registerErrorSnackbarText}>
    <Label></Label>
  </Snackbar>
</div>
