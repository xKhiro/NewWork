<script>
  import Navbar from '../components/Navbar.svelte';
  import { page } from '$app/stores';
  import Login from '../components/Login.svelte';
  import user from '../lib/user';

  $: isLoggedIn = $user === null ? false : true;
</script>

<div class="container">
  {#if isLoggedIn}
    <Navbar currentPage={$page.route.id} />
    <slot />
  {:else}
    <Login />
  {/if}
</div>
