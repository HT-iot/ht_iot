DROP KEYSPACE IF EXISTS "manager";
CREATE KEYSPACE manager WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;

CREATE TABLE manager.users (
    email text PRIMARY KEY,
    password text
) WITH bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
    AND compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99PERCENTILE';

CREATE TABLE manager.clients_by_user (
    user text,
    id timeuuid,
    access_key text,
    meta map<text, text>,
    name text,
    type text,
    PRIMARY KEY (user, id)
) WITH CLUSTERING ORDER BY (id ASC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
    AND compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99PERCENTILE';

CREATE TABLE manager.channels_by_user (
    user text,
    id timeuuid,
    connected set<text>,
    name text,
    PRIMARY KEY (user, id)
) WITH CLUSTERING ORDER BY (id ASC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
    AND compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99PERCENTILE';

CREATE MATERIALIZED VIEW manager.clients_by_channel AS
    SELECT id, user, connected
    FROM manager.channels_by_user
    WHERE id IS NOT NULL
    PRIMARY KEY (id, user)
    WITH CLUSTERING ORDER BY (user ASC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
    AND compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99PERCENTILE';

CREATE TABLE manager.device_info (
    hospitalname text,
    deviceid text,
    channelid text,
    availablepowersource int,
    batterylevel int,
    currenttime timestamp,
    devicetype text,
    errorcode int,
    factoryreset boolean,
    firmwareversion text,
    hardwareversion text,
    hospitaldeviceid text,
    manufacturer text,
    memoryfree int,
    modelnumber text,
    powersourcesurrent int,
    powersourcevoltage int,
    reboot boolean,
    reseterrorcode text,
    softwareversion text,
    supportedbindingmodes text,
    timezone text,
    used boolean,
    utcoffset text,
    PRIMARY KEY (hospitalname, deviceid, channelid)
) WITH CLUSTERING ORDER BY (deviceid ASC, channelid ASC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
    AND compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99PERCENTILE';

CREATE TABLE manager.hospital_by_patient (
    hospitalname text,
    deviceid text,
    channelid text,
    hospitalbed text,
    hospitaldeviceid text,
    hospitalid text,
    hospitalzone text,
    id timeuuid,
    inhospital boolean,
    meta map<text, text>,
    patiententrtime timestamp,
    patientexittime timestamp,
    patientid text,
    patientname text,
    patientsex text,
    PRIMARY KEY (hospitalname, deviceid, channelid)
) WITH CLUSTERING ORDER BY (deviceid ASC, channelid ASC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
    AND compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99PERCENTILE';
INSERT INTO "manager"."users" ("email", "password") VALUES ('123@email.com', '$2a$10$sKijwjy4ETAUxWZfmeMX5O9GTHyb1jJAlOHZ1zl0DvLJmoS.MU0sa');
INSERT INTO "manager"."users" ("email", "password") VALUES ('yangtao@email.com', '$2a$10$v7wv7vR.Au7sJImGL6ROleKYS1nrwZzHlIXTB2jXwNB7.Dt7.xm32');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', f26dc0b3-242a-11e8-8c6a-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI3NTAsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiZjI2ZGMwYjMtMjQyYS0xMWU4LThjNmEtMDAwYzI5MzlhODE5In0.EARIZ2AhwI7F4r2sung5Cd66THa5VmcbQNF06awwjmw', '深圳市眼科医院dev1', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 16585c12-242b-11e8-8c6b-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI4MTEsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMTY1ODVjMTItMjQyYi0xMWU4LThjNmItMDAwYzI5MzlhODE5In0.uMCY-a8uL7Myxc1iyYQVSVA3RZid2mEWfM1zuQcqwyQ', '深圳市眼科医院dev2', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 165ba990-242b-11e8-8c6c-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI4MTEsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMTY1YmE5OTAtMjQyYi0xMWU4LThjNmMtMDAwYzI5MzlhODE5In0._x2zU2pHsy8FAzKsNMEoxKgmrPzGJFlD04soT08Yhng', '深圳市眼科医院dev3', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 1661441e-242b-11e8-8c6d-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI4MTEsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMTY2MTQ0MWUtMjQyYi0xMWU4LThjNmQtMDAwYzI5MzlhODE5In0.9JXmYhze9Vbpk1za6ITcpWjcAxW9UJWIzouaNUCslr0', '深圳市眼科医院dev4', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 1665c6cd-242b-11e8-8c6e-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI4MTEsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMTY2NWM2Y2QtMjQyYi0xMWU4LThjNmUtMDAwYzI5MzlhODE5In0.t0inXfEn4bxtRYVWbLwgKrkwbRl9TL9mDTpMUfPGUAM', '深圳市眼科医院dev5', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 166a3e0a-242b-11e8-8c6f-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI4MTEsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMTY2YTNlMGEtMjQyYi0xMWU4LThjNmYtMDAwYzI5MzlhODE5In0.p3clIQUcJtoQp6KwIzTnIwcfmyqnn4wX-YS1BO7pJM8', '深圳市眼科医院dev6', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 166f2912-242b-11e8-8c70-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI4MTEsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMTY2ZjI5MTItMjQyYi0xMWU4LThjNzAtMDAwYzI5MzlhODE5In0.djCeGnPtINEngifOQUGIrkHxcou21id_KuzFhA_w_TI', '深圳市眼科医院dev7', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 1673b62e-242b-11e8-8c71-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI4MTEsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMTY3M2I2MmUtMjQyYi0xMWU4LThjNzEtMDAwYzI5MzlhODE5In0.uDjclaPn9EukbBGKSHpqOEhhDE5dBBPvxEJRkqZA4xU', '深圳市眼科医院dev8', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 1678edf7-242b-11e8-8c72-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI4MTEsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMTY3OGVkZjctMjQyYi0xMWU4LThjNzItMDAwYzI5MzlhODE5In0.Baui0nf7XUKIYMU1uNifeokbDA832n6S-aO5_MSROoA', '深圳市眼科医院dev9', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 4c7d7a99-242b-11e8-8c73-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI5MDIsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNGM3ZDdhOTktMjQyYi0xMWU4LThjNzMtMDAwYzI5MzlhODE5In0.Pev49lE7rUe87AnQpW6nrUv4a8BVRXexY1SUb-9so2k', '深圳市人民医院dev1', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 4c841256-242b-11e8-8c74-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI5MDIsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNGM4NDEyNTYtMjQyYi0xMWU4LThjNzQtMDAwYzI5MzlhODE5In0.znfXvINQPXfDKCgxHfHnsnpiovHFXfdYleyLVP8xz-E', '深圳市人民医院dev2', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 4c87e1d6-242b-11e8-8c75-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI5MDIsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNGM4N2UxZDYtMjQyYi0xMWU4LThjNzUtMDAwYzI5MzlhODE5In0._uCRnCYlUBWkNRBQmy_xO_ZUH19LnsJRwJorWMnaVSg', '深圳市人民医院dev3', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 4c8bf886-242b-11e8-8c76-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI5MDIsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNGM4YmY4ODYtMjQyYi0xMWU4LThjNzYtMDAwYzI5MzlhODE5In0.YThhWROMVBob1FsqZa-olcK90HOk3tTohPO1X-MPo2Y', '深圳市人民医院dev4', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 4c8ed2c7-242b-11e8-8c77-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI5MDIsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNGM4ZWQyYzctMjQyYi0xMWU4LThjNzctMDAwYzI5MzlhODE5In0.YAqxFEhKWmKCFpAl42nGpRBS4ofl2-2kzpegeD-FD_A', '深圳市人民医院dev5', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 4c955b1c-242b-11e8-8c78-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI5MDIsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNGM5NTViMWMtMjQyYi0xMWU4LThjNzgtMDAwYzI5MzlhODE5In0.-g1ff24FXJy26ST_6FtgHYNpZ43LzZTJv7PcoONO9HU', '深圳市人民医院dev6', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('123@email.com', 4d39e5f2-242b-11e8-8c79-000c2939a819, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjA2NjI5MDMsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNGQzOWU1ZjItMjQyYi0xMWU4LThjNzktMDAwYzI5MzlhODE5In0.94B1b99Fuv1yzKtWoI6UOPLqcT6F08BCYpWw5JPtPJk', '深圳市人民医院dev7', 'device');
INSERT INTO "manager"."channels_by_user" ("user", "id", "connected", "name") VALUES ('123@email.com', 8eac4a8b-242d-11e8-8c7a-000c2939a819, {'16585c12-242b-11e8-8c6b-000c2939a819', '165ba990-242b-11e8-8c6c-000c2939a819', '1661441e-242b-11e8-8c6d-000c2939a819', '1665c6cd-242b-11e8-8c6e-000c2939a819', '166a3e0a-242b-11e8-8c6f-000c2939a819', '166f2912-242b-11e8-8c70-000c2939a819', '1673b62e-242b-11e8-8c71-000c2939a819', '1678edf7-242b-11e8-8c72-000c2939a819', 'f26dc0b3-242a-11e8-8c6a-000c2939a819'}, '深圳市眼科医院');
INSERT INTO "manager"."channels_by_user" ("user", "id", "connected", "name") VALUES ('123@email.com', bfb1b279-242d-11e8-8c7b-000c2939a819, {'4c7d7a99-242b-11e8-8c73-000c2939a819', '4c841256-242b-11e8-8c74-000c2939a819', '4c87e1d6-242b-11e8-8c75-000c2939a819', '4c8bf886-242b-11e8-8c76-000c2939a819', '4c8ed2c7-242b-11e8-8c77-000c2939a819', '4c955b1c-242b-11e8-8c78-000c2939a819', '4d39e5f2-242b-11e8-8c79-000c2939a819'}, '深圳市人民医院');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '16585c12-242b-11e8-8c6b-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', 0, 0, 1520668565701, '', 0, False, '', '', '1', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '165ba990-242b-11e8-8c6c-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', 0, 0, 1520668573468, '', 0, False, '', '', '2', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '1661441e-242b-11e8-8c6d-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', 0, 0, 1520668580981, '', 0, False, '', '', '3', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '1665c6cd-242b-11e8-8c6e-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', 0, 0, 1520668588682, '', 0, False, '', '', '4', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '166a3e0a-242b-11e8-8c6f-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', 0, 0, 1520668596659, '', 0, False, '', '', '5', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '166f2912-242b-11e8-8c70-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', 0, 0, 1520668604835, '', 0, False, '', '', '6', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '1673b62e-242b-11e8-8c71-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', 0, 0, 1520668611844, '', 0, False, '', '', '7', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '1678edf7-242b-11e8-8c72-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', 0, 0, 1520668623629, '', 0, False, '', '', '8', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', 'f26dc0b3-242a-11e8-8c6a-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', 0, 0, 1520668631443, '', 0, False, '', '', '9', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '4c7d7a99-242b-11e8-8c73-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', 0, 0, 1520668164615, '', 0, False, '', '', '1', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '4c841256-242b-11e8-8c74-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', 0, 0, 1520668173303, '', 0, False, '', '', '2', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '4c87e1d6-242b-11e8-8c75-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', 0, 0, 1520668181036, '', 0, False, '', '', '3', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '4c8bf886-242b-11e8-8c76-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', 0, 0, 1520668194632, '', 0, False, '', '', '4', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '4c8ed2c7-242b-11e8-8c77-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', 0, 0, 1520668203984, '', 0, False, '', '', '5', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '4c955b1c-242b-11e8-8c78-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', 0, 0, 1520668213390, '', 0, False, '', '', '6', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '4d39e5f2-242b-11e8-8c79-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', 0, 0, 1520668221398, '', 0, False, '', '', '7', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '16585c12-242b-11e8-8c6b-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', '1', '1', '', '外科', b4343c80-2438-11e8-924a-000c2939a819, False, 1520668659489, '', '1', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '165ba990-242b-11e8-8c6c-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', '2', '2', '', '外科', bb5bdf2d-2438-11e8-924b-000c2939a819, False, 1520668671492, '', '2', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '1661441e-242b-11e8-8c6d-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', '3', '3', '', '外科', cc78c73d-2438-11e8-924c-000c2939a819, False, 1520668700203, '', '2', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '1665c6cd-242b-11e8-8c6e-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', '4', '4', '', '外科', d324d342-2438-11e8-924d-000c2939a819, False, 1520668711397, '', '2', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '166a3e0a-242b-11e8-8c6f-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', '5', '5', '', '外科', d87b8f85-2438-11e8-924e-000c2939a819, False, 1520668720354, '', '2', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '166f2912-242b-11e8-8c70-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', '6', '6', '', '外科', e43be249-2438-11e8-924f-000c2939a819, False, 1520668740069, '', '2', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '1673b62e-242b-11e8-8c71-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', '7', '7', '', '外科', ea35a89d-2438-11e8-9250-000c2939a819, False, 1520668750095, '', '2', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '1678edf7-242b-11e8-8c72-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', '8', '8', '', '外科', efbc03f5-2438-11e8-9251-000c2939a819, False, 1520668759364, '', '2', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', 'f26dc0b3-242a-11e8-8c6a-000c2939a819', '8eac4a8b-242d-11e8-8c7a-000c2939a819', '9', '9', '', '外科', f56be69a-2438-11e8-9252-000c2939a819, False, 1520668768905, '', '2', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '4c7d7a99-242b-11e8-8c73-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', '1', '1', '', '内科', bb492c9d-2437-11e8-ae51-000c2939a819, False, 1520668241873, '', '张三', '男');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '4c841256-242b-11e8-8c74-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', '2', '2', '', '内科', 1cb53b36-2438-11e8-9243-000c2939a819, False, 1520668405320, '', '张四', '男');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '4c87e1d6-242b-11e8-8c75-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', '3', '3', '', '外科', 4182b872-2438-11e8-9244-000c2939a819, False, 1520668467065, '', '李四', '女');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '4c8bf886-242b-11e8-8c76-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', '4', '4', '', '外科', 4af8e5be-2438-11e8-9245-000c2939a819, False, 1520668482939, '', '王伟', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '4c8ed2c7-242b-11e8-8c77-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', '5', '5', '', '内科', 51b5e6d4-2438-11e8-9246-000c2939a819, False, 1520668494244, '', '3', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '4c955b1c-242b-11e8-8c78-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', '6', '6', '', '内科', 581b259b-2438-11e8-9247-000c2939a819, False, 1520668504974, '', '3', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalid", "hospitalzone", "id", "inhospital", "patiententrtime", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '4d39e5f2-242b-11e8-8c79-000c2939a819', 'bfb1b279-242d-11e8-8c7b-000c2939a819', '8', '7', '', '内科', 648f9be2-2438-11e8-9249-000c2939a819, False, 1520668525870, '', '3', '');
DROP KEYSPACE IF EXISTS "message_writer";
CREATE KEYSPACE message_writer WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;

CREATE TABLE message_writer.messages_by_channel (
    channel timeuuid,
    id timeuuid,
    l text,
    n text,
    protocol text,
    publisher text,
    s double,
    t double,
    u text,
    ut double,
    v double,
    vb boolean,
    vd text,
    vs text,
    PRIMARY KEY (channel, id)
) WITH CLUSTERING ORDER BY (id DESC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
    AND compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99PERCENTILE';
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d916eb10-2433-11e8-9f65-65e0a0f47d44, '', '4d39e5f2-242b-11e8-8c79-000c2939a819:oxgen', 'http', '4d39e5f2-242b-11e8-8c79-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d9164ed0-2433-11e8-9f65-65e0a0f47d44, '', '4d39e5f2-242b-11e8-8c79-000c2939a819:pressure', 'http', '4d39e5f2-242b-11e8-8c79-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d9158b80-2433-11e8-9f65-65e0a0f47d44, '', '4d39e5f2-242b-11e8-8c79-000c2939a819:puls', 'http', '4d39e5f2-242b-11e8-8c79-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8b0fd50-2433-11e8-9f65-65e0a0f47d44, '', '4c955b1c-242b-11e8-8c78-000c2939a819:oxgen', 'http', '4c955b1c-242b-11e8-8c78-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8af9dc0-2433-11e8-9f65-65e0a0f47d44, '', '4c955b1c-242b-11e8-8c78-000c2939a819:pressure', 'http', '4c955b1c-242b-11e8-8c78-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8ad7ae0-2433-11e8-9f65-65e0a0f47d44, '', '4c955b1c-242b-11e8-8c78-000c2939a819:puls', 'http', '4c955b1c-242b-11e8-8c78-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8a7ae80-2433-11e8-9f65-65e0a0f47d44, '', '4c8ed2c7-242b-11e8-8c77-000c2939a819:oxgen', 'http', '4c8ed2c7-242b-11e8-8c77-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8a73950-2433-11e8-9f65-65e0a0f47d44, '', '4c8ed2c7-242b-11e8-8c77-000c2939a819:pressure', 'http', '4c8ed2c7-242b-11e8-8c77-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8a71240-2433-11e8-9f65-65e0a0f47d44, '', '4c8ed2c7-242b-11e8-8c77-000c2939a819:puls', 'http', '4c8ed2c7-242b-11e8-8c77-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8a368c0-2433-11e8-9f65-65e0a0f47d44, '', '4c8bf886-242b-11e8-8c76-000c2939a819:oxgen', 'http', '4c8bf886-242b-11e8-8c76-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8a16cf0-2433-11e8-9f65-65e0a0f47d44, '', '4c8bf886-242b-11e8-8c76-000c2939a819:pressure', 'http', '4c8bf886-242b-11e8-8c76-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8a11ed0-2433-11e8-9f65-65e0a0f47d44, '', '4c8bf886-242b-11e8-8c76-000c2939a819:puls', 'http', '4c8bf886-242b-11e8-8c76-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d89c3cd0-2433-11e8-9f65-65e0a0f47d44, '', '4c87e1d6-242b-11e8-8c75-000c2939a819:oxgen', 'http', '4c87e1d6-242b-11e8-8c75-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d89beeb0-2433-11e8-9f65-65e0a0f47d44, '', '4c87e1d6-242b-11e8-8c75-000c2939a819:pressure', 'http', '4c87e1d6-242b-11e8-8c75-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d89bc7a0-2433-11e8-9f65-65e0a0f47d44, '', '4c87e1d6-242b-11e8-8c75-000c2939a819:puls', 'http', '4c87e1d6-242b-11e8-8c75-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d895ad20-2433-11e8-9f65-65e0a0f47d44, '', '4c841256-242b-11e8-8c74-000c2939a819:oxgen', 'http', '4c841256-242b-11e8-8c74-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d89510e0-2433-11e8-9f65-65e0a0f47d44, '', '4c841256-242b-11e8-8c74-000c2939a819:pressure', 'http', '4c841256-242b-11e8-8c74-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d8949bb0-2433-11e8-9f65-65e0a0f47d44, '', '4c841256-242b-11e8-8c74-000c2939a819:puls', 'http', '4c841256-242b-11e8-8c74-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d88d96d0-2433-11e8-9f65-65e0a0f47d44, '', '4c7d7a99-242b-11e8-8c73-000c2939a819:oxgen', 'http', '4c7d7a99-242b-11e8-8c73-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d88cd380-2433-11e8-9f65-65e0a0f47d44, '', '4c7d7a99-242b-11e8-8c73-000c2939a819:pressure', 'http', '4c7d7a99-242b-11e8-8c73-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (bfb1b279-242d-11e8-8c7b-000c2939a819, d88c8560-2433-11e8-9f65-65e0a0f47d44, '', '4c7d7a99-242b-11e8-8c73-000c2939a819:puls', 'http', '4c7d7a99-242b-11e8-8c73-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, df2ca590-2432-11e8-9f65-65e0a0f47d44, '', '1678edf7-242b-11e8-8c72-000c2939a819:oxgen', 'http', '1678edf7-242b-11e8-8c72-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, df2c3060-2432-11e8-9f65-65e0a0f47d44, '', '1678edf7-242b-11e8-8c72-000c2939a819:pressure', 'http', '1678edf7-242b-11e8-8c72-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, df2b6d10-2432-11e8-9f65-65e0a0f47d44, '', '1678edf7-242b-11e8-8c72-000c2939a819:puls', 'http', '1678edf7-242b-11e8-8c72-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de9e9660-2432-11e8-9f65-65e0a0f47d44, '', '1673b62e-242b-11e8-8c71-000c2939a819:oxgen', 'http', '1673b62e-242b-11e8-8c71-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de9ce8b0-2432-11e8-9f65-65e0a0f47d44, '', '1673b62e-242b-11e8-8c71-000c2939a819:pressure', 'http', '1673b62e-242b-11e8-8c71-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de9b13f0-2432-11e8-9f65-65e0a0f47d44, '', '1673b62e-242b-11e8-8c71-000c2939a819:puls', 'http', '1673b62e-242b-11e8-8c71-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de960ae0-2432-11e8-9f65-65e0a0f47d44, '', '166f2912-242b-11e8-8c70-000c2939a819:oxgen', 'http', '166f2912-242b-11e8-8c70-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de945d30-2432-11e8-9f65-65e0a0f47d44, '', '166f2912-242b-11e8-8c70-000c2939a819:pressure', 'http', '166f2912-242b-11e8-8c70-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de940f10-2432-11e8-9f65-65e0a0f47d44, '', '166f2912-242b-11e8-8c70-000c2939a819:puls', 'http', '166f2912-242b-11e8-8c70-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de8e90d0-2432-11e8-9f65-65e0a0f47d44, '', '166a3e0a-242b-11e8-8c6f-000c2939a819:oxgen', 'http', '166a3e0a-242b-11e8-8c6f-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de8e69c0-2432-11e8-9f65-65e0a0f47d44, '', '166a3e0a-242b-11e8-8c6f-000c2939a819:pressure', 'http', '166a3e0a-242b-11e8-8c6f-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de8d7f60-2432-11e8-9f65-65e0a0f47d44, '', '166a3e0a-242b-11e8-8c6f-000c2939a819:puls', 'http', '166a3e0a-242b-11e8-8c6f-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de8ae750-2432-11e8-9f65-65e0a0f47d44, '', '1665c6cd-242b-11e8-8c6e-000c2939a819:oxgen', 'http', '1665c6cd-242b-11e8-8c6e-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de8960b0-2432-11e8-9f65-65e0a0f47d44, '', '1665c6cd-242b-11e8-8c6e-000c2939a819:pressure', 'http', '1665c6cd-242b-11e8-8c6e-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, de889d60-2432-11e8-9f65-65e0a0f47d44, '', '1665c6cd-242b-11e8-8c6e-000c2939a819:puls', 'http', '1665c6cd-242b-11e8-8c6e-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, 423ed960-2432-11e8-9f65-65e0a0f47d44, '', '1661441e-242b-11e8-8c6d-000c2939a819:oxgen', 'http', '1661441e-242b-11e8-8c6d-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, 423eb250-2432-11e8-9f65-65e0a0f47d44, '', '1661441e-242b-11e8-8c6d-000c2939a819:pressure', 'http', '1661441e-242b-11e8-8c6d-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, 423d52c0-2432-11e8-9f65-65e0a0f47d44, '', '1661441e-242b-11e8-8c6d-000c2939a819:puls', 'http', '1661441e-242b-11e8-8c6d-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, ec284660-2431-11e8-9f65-65e0a0f47d44, '', '165ba990-242b-11e8-8c6c-000c2939a819:oxgen', 'http', '165ba990-242b-11e8-8c6c-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, ec281f50-2431-11e8-9f65-65e0a0f47d44, '', '165ba990-242b-11e8-8c6c-000c2939a819:pressure', 'http', '165ba990-242b-11e8-8c6c-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, ec26e6d0-2431-11e8-9f65-65e0a0f47d44, '', '165ba990-242b-11e8-8c6c-000c2939a819:puls', 'http', '165ba990-242b-11e8-8c6c-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, 8c474030-2430-11e8-9f65-65e0a0f47d44, '', '16585c12-242b-11e8-8c6b-000c2939a819:oxgen', 'http', '16585c12-242b-11e8-8c6b-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, 8c471920-2430-11e8-9f65-65e0a0f47d44, '', '16585c12-242b-11e8-8c6b-000c2939a819:pressure', 'http', '16585c12-242b-11e8-8c6b-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, 8c4607b0-2430-11e8-9f65-65e0a0f47d44, '', '16585c12-242b-11e8-8c6b-000c2939a819:puls', 'http', '16585c12-242b-11e8-8c6b-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, 2d656510-2430-11e8-9f65-65e0a0f47d44, '', 'f26dc0b3-242a-11e8-8c6a-000c2939a819:oxgen', 'http', 'f26dc0b3-242a-11e8-8c6a-000c2939a819', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, 2d653e00-2430-11e8-9f65-65e0a0f47d44, '', 'f26dc0b3-242a-11e8-8c6a-000c2939a819:pressure', 'http', 'f26dc0b3-242a-11e8-8c6a-000c2939a819', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (8eac4a8b-242d-11e8-8c7a-000c2939a819, 2d64efe0-2430-11e8-9f65-65e0a0f47d44, '', 'f26dc0b3-242a-11e8-8c6a-000c2939a819:puls', 'http', 'f26dc0b3-242a-11e8-8c6a-000c2939a819', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
