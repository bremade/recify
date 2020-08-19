<script>
  import Card from '@smui/card';
  import Button from '@smui/button';

  function logout() {
    fetch('/api/v1/auth/logout', {
      method: 'POST'
    }).then(result => {
      // This should never fail
      if (result.status === 200) {
        window.location.pathname = '/';
      }
    });
  }
</script>

<div class="login-container">
  <div class="login">
    <Card variant="outlined">
      <div class="title">
        Logout
      </div>
      <p>
        Are you sure?
      </p>
      <div class="input-container">
        <Button class="logout-button" variant="unelevated" on:click={logout}>Logout</Button>
      </div>
    </Card>
  </div>
</div>
