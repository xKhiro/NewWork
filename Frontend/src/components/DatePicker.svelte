<script>
  import { onMount } from 'svelte';
  import { Datepicker } from 'vanillajs-datepicker';

  onMount(() => {
    const elem = document.querySelector('input[name="datePicker"]');
    const datepicker = new Datepicker(elem, {
      buttonClass: 'btn',
      allowOneSidedRange: true,
      format: 'dd.mm.yyyy',
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
