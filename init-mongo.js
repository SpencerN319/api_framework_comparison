db.createUser({
  user: 'root',
  pwd: 'root',
  roles: [
    {
      role: 'readWrite',
      db: 'test',
    },
  ],
});

db = new Mongo().getDB("test");

db.createCollection('vulnerabilities', { capped: false });
