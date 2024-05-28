# frog blossom

## Database for frog blossom CMS

Database repository for a web-based CMS. Herein are the scripts, configurations, and documentation related to the database architecture and management of the CMS application.

## About the CMS System

The CMS system is a web-based platform designed to empower users to create, manage, and publish digital content without requiring extensive technical knowledge. Key features of the CMS system include:

- **Website Creation**: Users can create websites with ease using the intuitive interface. They can choose from a variety of templates, customize layouts, and manage content effortlessly.

- **Content Management**: The CMS system enables users to create, edit, organize, and publish various types of digital content such as articles, blog posts, images, and videos. Content can be saved as drafts, published, or archived as needed.

## Technologies used

**db**: postgresql
open-source relational database management system used for storing and managing structured data within the application.

**migration tool**: golang-migrate
database migrations written in Go.

## Getting Started

<<<<<<< HEAD
### Run it with docker

There is a makefile that has the scripts for running an instance of the frog_blossom_db, postgreSQL

***start the Postgres docker container***

```bash
make postgres

```

***create frog_blossom_db database***

```bash
make cratedb
```

***drop frog_blossom_db database***

```bash
make dropdb
```

### schema migration

***migrate up frog_blossom_db database***

```bash
make migrateup
```

***migrate down frog_blossom_db database***

```bash
make migratedown
```


***warning: dirty migration error!***

You clean your migration, and then force clean and then attempt to migrate up/down again

```bash
make migrateforce
```
=======
### Basic usage

### Docker usage
>>>>>>> 75c7508e84c838784a1d79412c17aef70377e4c4

### Revision History

| Date       | Version | Description of Changes  | Author |
|------------|---------|-------------------------|--------|
| 2024-05-02 | 1.0     | Initial commit          | @minierparedes    |
| 2024-05-02 | 1.1     | migration files up/down | @minierparedes    |
| 2024-05-02 | 1.2     | makefile                | @minierparedes    |
| 2024-05-02 | 1.3     | readme                  | @minierparedes    |
| 2024-05-10 | 1.4     | db schema v5            | @minierparedes    |
<<<<<<< HEAD
| 2024-05-20 | 1.5     | readme: run docker      | @minierparedes    |
| 2024-05-27 | 1.6     | github actions      | @minierparedes    |
=======
>>>>>>> 75c7508e84c838784a1d79412c17aef70377e4c4
