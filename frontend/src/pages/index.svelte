<script>
  import Filler from '../components/Filler.svelte';
  import Navbar from '../components/Navbar.svelte';
  import session from '../stores/session.js';

  session.updateSessionStatus();
</script>

<div>
  <Navbar />
  <div style="background-color: #eeeeee; padding: 100px;">
    <div>Logged in {$session.logged_in} as user: {$session.username}</div>
    <a href="/login">Login</a><br />
    <a href="/logout">Logout</a>
  </div>
  <Filler/>
</div>
