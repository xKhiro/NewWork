<script>
  export let isError = false;
</script>

<div class="position-fixed top-0 end-0 p-3" style="z-index: 11">
  <div class="toast align-items-center text-white {isError ? 'bg-danger' : 'bg-primary'}" id="toast" role="alert">
    <div class="d-flex">
      <div class="toast-body" id="toastBody" />
      <button type="button" class="btn-close me-2 m-auto btn-close-white" data-bs-dismiss="toast" />
    </div>
  </div>
</div>
