<script>
  import DatePicker from '../../components/DatePicker.svelte';
  import DeleteBookingModal from '../../components/DeleteBookingModal.svelte';
  import Toast from '../../components/Toast.svelte';
  import { showToast } from '../../lib/toast';
  import { convertGermanDateToISO } from '../../lib/dateutils';
  import { getName } from '../../lib/ad';
  import { onMount } from 'svelte';
  import user from '../../lib/user';
  import { goto } from '$app/navigation';

  let selectedDate;
  let selectedWorkspace = null;
  let filteredWorkspaces = [];

  onMount(() => {
    if ($user === null || $user.role !== 'admin') {
      goto('/');
    }
  });

  function handleRowSelection(b) {
    selectedWorkspace = selectedWorkspace !== null && selectedWorkspace.bookingId === b.bookingId ? null : b;
  }

  async function getWorkspaces(e) {
    const formData = new FormData(e.target);

    const filter = {};
    for (let field of formData) {
      const [key, value] = field;
      filter[key] = value;
    }

    const queryParams = new URLSearchParams();
    queryParams.append('date', convertGermanDateToISO(filter['datePicker']));
    queryParams.append('hasAdjustableDesk', filter['filterAusstattung'] === '1');
    queryParams.append('hasDockingStation', filter['filterAusstattung'] === '2');
    queryParams.append('hasTwoScreens', filter['filterAusstattung'] === '3');

    if (!filter['filterArbeitsplaetze'].includes('Alle')) {
      queryParams.append('workspaceId', parseInt(filter['filterArbeitsplaetze']));
    }

    if (!filter['filterRaeume'].includes('Alle')) {
      queryParams.append('roomId', parseInt(filter['filterRaeume']));
    }

    const json = await fetch('http://localhost:8000/workspaces?booked=true&' + queryParams).then((res) => res.json());

    filteredWorkspaces = json
      .filter((ws) => {
        return ws.booking !== null;
      })
      .map(({ hasTwoScreens, hasDockingStation, hasAdjustableDesk, ...rest }) => {
        const features = [];

        if (hasTwoScreens) {
          features.push('2 Bildschirme');
        }
        if (hasDockingStation) {
          features.push('Docking Station');
        }
        if (hasAdjustableDesk) {
          features.push('Schreibtisch höhenverstellbar');
        }

        return {
          ...rest,
          features: features.join(', '),
        };
      });
  }

  async function deleteBooking(b) {
    await fetch(`http://localhost:8000/users/${b.personId}/bookings/${b.bookingId}`, {
      method: 'DELETE',
    });

    // Remove selection
    selectedWorkspace = null;

    filteredWorkspaces = filteredWorkspaces.filter((fw) => fw.bookingId !== b.bookingId);

    showToast('Stornierung erfolgreich');
  }
</script>

<DeleteBookingModal deleteBooking={() => deleteBooking(selectedWorkspace)} />
<Toast />

<form on:submit|preventDefault={getWorkspaces} class="mt-4">
  <div class="row row-cols-3">
    <div class="col-2" style="min-width: 13rem;">
      <DatePicker bind:selectedDate />
    </div>
    <div class="col-2" style="min-width: 13rem;">
      <select name="filterArbeitsplaetze" class="form-select">
        <option selected>Alle Arbeitsplätze</option>
        <option value="1">Arbeitsplatz 01</option>
        <option value="2">Arbeitsplatz 02</option>
        <option value="3">Arbeitsplatz 03</option>
        <option value="4">Arbeitsplatz 04</option>
        <option value="5">Arbeitsplatz 05</option>
        <option value="6">Arbeitsplatz 06</option>
        <option value="7">Arbeitsplatz 07</option>
        <option value="8">Arbeitsplatz 08</option>
        <option value="9">Arbeitsplatz 09</option>
        <option value="10">Arbeitsplatz 10</option>
      </select>
    </div>
    <div class="col-2" style="min-width: 13rem;">
      <select name="filterMitarbeiter" class="form-select">
        <option selected>Alle Mitarbeiter</option>
        <option value="1">admin</option>
        <option value="2">user</option>
        <option value="3">user2</option>
        <option value="4">user3</option>
        <option value="5">user4</option>
      </select>
    </div>
  </div>
  <div class="row row-cols-3 mt-3">
    <div class="col-2" style="min-width: 13rem;">
      <select name="filterRaeume" class="form-select">
        <option selected>Alle Räume</option>
        <option value="1">Raum 01</option>
        <option value="2">Raum 02</option>
        <option value="3">Raum 03</option>
        <option value="4">Raum 04</option>
        <option value="5">Raum 05</option>
      </select>
    </div>
    <div class="col-2" style="min-width: 13rem;">
      <select name="filterAusstattung" class="form-select">
        <option selected>Alle Ausstattung</option>
        <option value="1">Schreibtisch höhenverstellbar</option>
        <option value="2">Docking Station</option>
        <option value="3">2 Bildschirme</option>
      </select>
    </div>
    <div class="col-2">
      <button type="submit" class="btn btn-secondary col-3" style="min-width: 6rem;">Suchen</button>
    </div>
  </div>
</form>

<div class="mt-4">
  <div class="mt-4">
    <table class="table table-hover table-bordered">
      <thead>
        <tr>
          <th scope="col">#</th>
          <th scope="col">Raum</th>
          <th scope="col">Arbeitsplatz</th>
          <th scope="col">Ausstattung</th>
          <th scope="col">Gebucht von</th>
        </tr>
      </thead>
      <tbody>
        {#each filteredWorkspaces as ws, index}
          {#if ws.bookings !== null}
            {#each ws.bookings as b}
              <tr
                on:click={() => handleRowSelection(b)}
                class:table-active={selectedWorkspace !== null && selectedWorkspace.bookingId === b.bookingId}>
                <th scope="row">{index + 1}</th>
                <td>{ws.roomName}</td>
                <td>{ws.name}</td>
                <td>{ws.features}</td>
                <td>{getName(b.personId)}</td>
              </tr>
            {/each}
          {/if}
        {:else}
          <p>Keine gebuchten Arbeitsplätze gefunden</p>
        {/each}
      </tbody>
    </table>

    <div class="d-flex justify-content-end mb-5">
      <button
        type="button"
        style="min-width: 6rem;"
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
</div>
