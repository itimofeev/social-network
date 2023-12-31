create user replicator with replication encrypted password 'replicator_password';
select pg_create_physical_replication_slot('replication_slot_sync');
select pg_create_physical_replication_slot('replication_slot_async');