<script>
  import user from '../lib/user.js';
  import { getUser } from '../lib/ad.js';

  let username;
  let password;
  let invalidCredentials = false;

  function login() {
    const userData = getUser(username, password);

    if (!!userData) {
      user.update((u) => (u = { ...userData }));
    } else {
      invalidCredentials = true;
    }
  }
</script>

<div class="form-signin shadow-lg mt-5 rounded-4">
  <div class="content-size pt-5">
    <h1 class="fs-1 mb-4 fw-bold text-center">NEWWORK</h1>
    <form on:submit|preventDefault={login}>
      <div class="form-floating mb-3">
        <input
          type="text"
          class="form-control"
          class:is-invalid={invalidCredentials}
          id="username"
          placeholder="Benutzername"
          bind:value={username} />
        <label for="username">Benutzername</label>
      </div>
      <div class="form-floating mb-3">
        <input
          type="password"
          class="form-control"
          class:is-invalid={invalidCredentials}
          id="password"
          placeholder="Passwort"
          bind:value={password} />
        <label for="password">Passwort</label>
      </div>
      <div>
        <button type="submit" class="btn btn-lg btn-primary w-100">Anmelden</button>
      </div>
    </form>
  </div>
</div>

<style>
  .form-signin {
    max-width: 480px;
    min-height: 400px;
    margin: auto;
  }
  .content-size {
    max-width: 330px;
    margin: auto;
  }
</style>
