<script>
  import { showToast } from '../lib/toast';
  import { convertISODateToGerman } from '../lib/dateutils';
  import { goto } from '$app/navigation';
  import DeleteBookingModal from '../components/DeleteBookingModal.svelte';
  import Toast from '../components/Toast.svelte';

  export let data;

  let selectedWorkspace = null;

  function handleRowSelection(ws) {
    selectedWorkspace = selectedWorkspace !== null && selectedWorkspace.workspaceId === ws.workspaceId ? null : ws;
  }

  async function deleteBooking(ws) {
    // TODO: Add dynamic personId, Add error handling
    await fetch('http://localhost:8000/users/1/bookings/' + ws.booking.bookingId, {
      method: 'DELETE',
    });

    // Remove selection
    selectedWorkspace = null;

    // Hack to trigger load function
    goto('/');

    showToast('Stornierung erfolgreich');
  }
</script>

<DeleteBookingModal deleteBooking={() => deleteBooking(selectedWorkspace)} />
<Toast />

<div class="mt-4">
  <table class="table table-hover table-bordered">
    <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">Datum</th>
        <th scope="col">Raum</th>
        <th scope="col">Arbeitsplatz</th>
        <th scope="col">Ausstattung</th>
      </tr>
    </thead>
    <tbody>
      {#each data.filteredWorkspaces as ws, index}
        <tr
          on:click={() => handleRowSelection(ws)}
          class:table-active={selectedWorkspace !== null && selectedWorkspace.workspaceId === ws.workspaceId}>
          <th scope="row">{index + 1}</th>
          <td>{convertISODateToGerman(ws.booking.date)}</td>
          <td>{ws.roomName}</td>
          <td>{ws.name}</td>
          <td>{ws.features}</td>
        </tr>
      {:else}
        <p>Keine Buchungen vorhanden</p>
      {/each}
    </tbody>
  </table>

  <div class="d-flex justify-content-end">
    <button
      type="button"
      data-bs-toggle="modal"
      data-bs-target="#deleteModal"
      class="btn btn-danger"
      class:disabled={selectedWorkspace === null}>Stornieren</button>
  </div>

  <style>
    td {
      cursor: pointer;
    }
  </style>
</div>
