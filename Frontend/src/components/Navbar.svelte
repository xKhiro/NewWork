<script>
  export let currentPage;
</script>

<nav class="navbar navbar-expand-lg bg-body-tertiary shadow-sm">
  <div class="container-fluid">
    <a class="navbar-brand fs-4 fw-bold" href="/">NEWWORK</a>
    <button
      class="navbar-toggler"
      type="button"
      data-bs-toggle="collapse"
      data-bs-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent"
      aria-expanded="false"
      aria-label="Toggle navigation">
      <span class="navbar-toggler-icon" />
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <a class="nav-link {currentPage === '/' ? 'active text-decoration-underline' : ''}" href="/"
            >Meine Buchungen</a>
        </li>
        <li class="nav-item">
          <a class="nav-link {currentPage === '/suche' ? 'active text-decoration-underline' : ''}" href="/suche"
            >Suche</a>
        </li>
        <!-- TODO: Hide admin by role -->
        <li class="nav-item">
          <a class="nav-link {currentPage === '/admin' ? 'active text-decoration-underline' : ''}" href="/admin"
            >Admin</a>
        </li>
      </ul>
      <ul class="navbar-nav mb-2 mb-lg-0">
        <li class="nav-item">
          <a class="nav-link" href="/">Abmelden</a>
        </li>
      </ul>
    </div>
  </div>
</nav>
