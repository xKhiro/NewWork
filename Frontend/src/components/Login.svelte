<script>
  import user from '../lib/user.js';
  import { getUser } from '../lib/ad.js';

  let username;
  let password;

  function login() {
    const userData = getUser(username, password);

    if (!!userData) {
      user.update((u) => (u = { ...userData }));
    }
  }
</script>

<div class="form-signin">
  <form on:submit|preventDefault={login}>
    <div class="form-floating">
      <input type="text" class="form-control" id="username" placeholder="Benutzername" bind:value={username} />
      <label for="username">Benutzername</label>
    </div>
    <div class="form-floating">
      <input type="password" class="form-control" id="password" placeholder="Passwort" bind:value={password} />
      <label for="password">Passwort</label>
    </div>
    <button type="submit" class="btn btn-primary mb-4">Einloggen</button>
  </form>
</div>

<style>
  .form-signin {
    width: 100%;
    max-width: 330px;
    padding: 15px;
    margin: auto;
  }

  .form-signin .form-floating:focus-within {
    z-index: 2;
  }

  .form-signin input[type='text'] {
    margin-bottom: -1px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }

  .form-signin input[type='password'] {
    margin-bottom: 10px;
    border-top-left-radius: 0;
    border-top-right-radius: 0;
  }
</style>
