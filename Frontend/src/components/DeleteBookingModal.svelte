<script>
  export let deleteBooking;
</script>

<div class="modal fade" id="deleteModal" tabindex="-1">
  <div class="modal-dialog modal-dialog-centered">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Stornierung bestätigen</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" />
      </div>
      <form on:submit|preventDefault={deleteBooking}>
        <div class="modal-body">
          <input type="hidden" name="id" id="id" />
          Bist du dir sicher, dass du die Buchung stornieren möchtest?
        </div>
        <div class="modal-footer">
          <input type="button" class="btn fw-bold btn-secondary" value="Abbrechen" data-bs-dismiss="modal" />
          <input
            type="submit"
            class="btn fw-bold btn-danger"
            value="Bestätigen"
            name="btnDelete"
            data-bs-dismiss="modal" />
        </div>
      </form>
    </div>
  </div>
</div>
