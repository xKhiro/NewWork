<script>
  import { onMount } from 'svelte';
  import { Datepicker } from 'vanillajs-datepicker';

  onMount(() => {
    const elem = document.querySelector('input[name="datePicker"]');
    const minDate = new Date();
    const maxDate = minDate.setDate(minDate.getDate() + 14);

    const datepicker = new Datepicker(elem, {
      buttonClass: 'btn',
      format: 'dd.mm.yyyy',
      minDate: new Date(),
      maxDate: maxDate,
    });
  });

  export let selectedDate = new Date(Date.now()).toLocaleString().split(',')[0];

  function handleChangeDate(e) {
    selectedDate = e.detail.datepicker.element.value;
  }
</script>

<div>
  <input
    type="text"
    class="form-control"
    name="datePicker"
    on:changeDate={handleChangeDate}
    bind:value={selectedDate} />
</div>
