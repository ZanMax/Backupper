# Backupper

Application to backup all your files and databases and store them securely.

### Configuration

Edit config.json

- Add files, databases to backup
- add destination where you want to save your file: s3, sftp etc.

### Config example

```
{
  "files": [
    "path/to/file1",
    "path/to/file2",
    "path/to/file3"
  ],
  "DBs": [
    {
      "type": "mysql",
      "connString": "mysql://user:pass@host:port/dbname",
      "dbName": [
        "dbname1",
        "dbname2",
        "dbname3"
      ]
    },
    {
      "type": "mysql",
      "connString": "mysql://user:pass@host:port/dbname",
      "dbName": [
        "dbname1",
        "dbname2",
        "dbname3"
      ]
    }
  ]
}
```

### DBs support

- MySQL
- PostgreSQL

### Destination support

- S3
- sFTP/FTP
- Samba

### Compression support

- gzip
