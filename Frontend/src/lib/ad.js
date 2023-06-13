const users = [
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

export function getName(personId) {
  return users.find((u) => u.personId == personId).name;
}

export function getUser(username, password) {
  return users.find((u) => u.name == username && u.password == password);
}
