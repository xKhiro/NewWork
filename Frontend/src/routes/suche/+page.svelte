<script>
  import DatePicker from '../../components/DatePicker.svelte';
  import { convertGermanDateToISO } from '../../lib/dateutils';
  import Toast from '../../components/Toast.svelte';
  import { showToast } from '../../lib/toast';
  import user from '../../lib/user';

  let filteredWorkspaces = [];
  let selectedDate;
  let selectedWorkspace = null;
  let isError = false;

  function handleRowSelection(ws) {
    selectedWorkspace = selectedWorkspace !== null && selectedWorkspace.workspaceId === ws.workspaceId ? null : ws;
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
    queryParams.append('booked', filter['filterBooked'] === 'booked' ? true : false);

    if (!filter['filterArbeitsplaetze'].includes('Alle')) {
      queryParams.append('workspaceId', parseInt(filter['filterArbeitsplaetze']));
    }

    if (!filter['filterRaeume'].includes('Alle')) {
      queryParams.append('roomId', parseInt(filter['filterRaeume']));
    }

    const json = await fetch('http://localhost:8000/workspaces?' + queryParams).then((res) => res.json());

    filteredWorkspaces = json.map(({ hasTwoScreens, hasDockingStation, hasAdjustableDesk, ...rest }) => {
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
        date: convertGermanDateToISO(selectedDate),
        features: features.join(', '),
      };
    });
  }

  async function createBooking(ws) {
    const res = await fetch(`http://localhost:8000/users/${$user.personId}/bookings`, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        personId: $user.personId,
        workspaceId: ws.workspaceId,
        roomName: ws.roomName,
        date: ws.date,
      }),
    });

    // Remove selection
    selectedWorkspace = null;

    if (res.status === 409) {
      isError = true;
      showToast('Arbeitsplatz bereits gebucht');
      isError = false;
      return;
    }

    filteredWorkspaces = filteredWorkspaces.filter((fw) => fw.workspaceId !== ws.workspaceId);

    showToast('Buchung erfolgreich');
  }
</script>

<Toast isError />

<form on:submit|preventDefault={getWorkspaces} class="mt-4">
  <div class="row row-cols-4">
    <div class="col-2">
      <DatePicker bind:selectedDate />
    </div>
    <div class="col-2">
      <select name="filterArbeitsplaetze" class="form-select">
        <option selected>Alle Arbeitsplätze</option>
        <option value="1">01</option>
        <option value="2">02</option>
        <option value="3">03</option>
        <option value="4">04</option>
        <option value="5">05</option>
        <option value="6">06</option>
        <option value="7">07</option>
        <option value="8">08</option>
        <option value="9">09</option>
        <option value="10">10</option>
      </select>
    </div>
    <div class="col-auto row mt-2 px-4">
      <div class="form-check col-auto">
        <input class="form-check-input" type="radio" name="filterBooked" id="radioBookedAvailable" checked />
        <label class="form-check-label" for="radioBookedAvailable">Nur freie Arbeitsplätze</label>
      </div>
      <div class="form-check col-auto">
        <input class="form-check-input" type="radio" name="filterBooked" id="radioBooked" value="booked" />
        <label class="form-check-label" for="radioBooked">Nur gebuchte Arbeitsplätze</label>
      </div>
    </div>
  </div>
  <div class="row row-cols-3 mt-3">
    <div class="col-2">
      <select name="filterRaeume" class="form-select">
        <option selected>Alle Räume</option>
        <option value="1">01</option>
        <option value="2">02</option>
        <option value="3">03</option>
        <option value="4">04</option>
        <option value="5">05</option>
      </select>
    </div>
    <div class="col-2">
      <select name="filterAusstattung" class="form-select">
        <option selected>Alle Ausstattung</option>
        <option value="1">Schreibtisch höhenverstellbar</option>
        <option value="2">Docking Station</option>
        <option value="3">2 Bildschirme</option>
      </select>
    </div>
    <div>
      <button type="submit" class="btn btn-secondary col-3">Suchen</button>
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
        </tr>
      </thead>
      <tbody>
        {#each filteredWorkspaces as ws, index}
          <tr
            on:click={() => handleRowSelection(ws)}
            class:table-active={selectedWorkspace !== null && selectedWorkspace.workspaceId === ws.workspaceId}>
            <th scope="row">{index + 1}</th>
            <td>{ws.name}</td>
            <td>{ws.roomName}</td>
            <td>{ws.features}</td>
          </tr>
        {:else}
          <p>Keine verfügbaren Arbeitsplätze gefunden</p>
        {/each}
      </tbody>
    </table>

    <div class="d-flex justify-content-end">
      <button
        type="button"
        class="btn btn-primary"
        class:disabled={selectedWorkspace === null}
        on:click={() => createBooking(selectedWorkspace)}>Buchen</button>
    </div>

    <style>
      td {
        cursor: pointer;
      }
    </style>
  </div>
</div>
