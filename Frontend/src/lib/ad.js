const persons = [
  {
    personId: '1',
    name: 'admin',
    password: 'admin12',
    role: 'admin',
  },
  {
    personId: '2',
    name: 'user',
    password: 'user12',
    role: 'standarduser',
  },
];

export function getPerson(personId) {
  return persons.find((p) => p.personId === personId);
}
