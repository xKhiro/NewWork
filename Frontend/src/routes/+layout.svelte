<script>
  import Navbar from '../components/Navbar.svelte';
  import { page } from '$app/stores';
</script>

<div class="container">
  <Navbar currentPage={$page.route.id} />
  <slot />
</div>
