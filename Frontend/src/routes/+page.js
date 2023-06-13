export async function load({ fetch }) {
  //TODO: Replace personId
  const personId = '1';

  const json = await fetch('http://localhost:8000/workspaces?booked=true').then((res) => res.json());

  const filteredWorkspaces = json
    .filter((ws) => {
      return ws.booking !== null && ws.bookings.some((b) => b.personId === personId);
    })
    .map(({ hasTwoScreens, hasDockingStation, hasAdjustableDesk, bookings, ...rest }) => {
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

      const booking = bookings.find((b) => b.personId === personId);

      return {
        ...rest,
        features: features.join(', '),
        booking: booking,
      };
    });

  return { filteredWorkspaces };
}
