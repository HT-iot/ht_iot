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
INSERT INTO "manager"."users" ("email", "password") VALUES ('123@test.com', '123');
INSERT INTO "manager"."users" ("email", "password") VALUES ('yangtao@email.com', '$2a$10$3MSUE4ztvnrcm7mjEsZecOG9.oJ2xFmtXdyo9llnd/UzYbNHSZSOm');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', 01639ab2-f835-11e7-992d-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4MjkyMTksImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMDE2MzlhYjItZjgzNS0xMWU3LTk5MmQtMDAxNjNlMDg0NTk4In0.yJB8B4ulHQXmJCa6sfd97t0D_vpdwCV3Tbtckb7q2ec', '深圳市眼科医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', 0a1de8de-f835-11e7-992e-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4MjkyMzQsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMGExZGU4ZGUtZjgzNS0xMWU3LTk5MmUtMDAxNjNlMDg0NTk4In0.EwxwV6wOSqbjbGyRG22UkNie03sIIQ47Y4E7skIk9l4', '深圳市眼科医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', 32952fe1-f835-11e7-992f-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4MjkzMDIsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiMzI5NTJmZTEtZjgzNS0xMWU3LTk5MmYtMDAxNjNlMDg0NTk4In0.cyDSy7J2sQCRb96mZdziN6g7u1Vl1-Wl0zOPz_z_HpA', '深圳市眼科医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', 4ddf6970-f835-11e7-9930-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4MjkzNDgsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNGRkZjY5NzAtZjgzNS0xMWU3LTk5MzAtMDAxNjNlMDg0NTk4In0.j-VdssyjetsbbjioNRW0BUdePHTktD2cXHr4aYS5_ic', '深圳市人民医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', 54931a27-f835-11e7-9931-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4MjkzNTksImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNTQ5MzFhMjctZjgzNS0xMWU3LTk5MzEtMDAxNjNlMDg0NTk4In0.3A9G0uzCSWS1ln9aDxTapVsdoIo1zG5CDwL6XBevOxU', '深圳市人民医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', 57dbfd30-f835-11e7-9932-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4MjkzNjQsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiNTdkYmZkMzAtZjgzNS0xMWU3LTk5MzItMDAxNjNlMDg0NTk4In0.4uKrQB3jHuWzYvNeiQso0HtQDEh2T7KECiulFF0Rb6Q', '深圳市人民医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', 9a95d806-f835-11e7-aab7-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4Mjk0NzYsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiOWE5NWQ4MDYtZjgzNS0xMWU3LWFhYjctMDAxNjNlMDg0NTk4In0.LYozFasLPfqpLYKy-mWDi3Xuojb9aILedqiKAjGyVZU', '深圳市人民医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', 9b53cc9e-f835-11e7-aab8-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4Mjk0NzgsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiOWI1M2NjOWUtZjgzNS0xMWU3LWFhYjgtMDAxNjNlMDg0NTk4In0.VZixkWmAnEE9PJ4eZzaofdIDSDoTaMCjGR6fUkZHkvw', '深圳市人民医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', 9c0d7645-f835-11e7-aab9-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4Mjk0NzksImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiOWMwZDc2NDUtZjgzNS0xMWU3LWFhYjktMDAxNjNlMDg0NTk4In0.xu62jl0iI2oEDjF_LasCx6UKIFjqSXbec1B0I4aE-rE', '深圳市人民医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', a414b784-f835-11e7-aaba-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4Mjk0OTIsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiYTQxNGI3ODQtZjgzNS0xMWU3LWFhYmEtMDAxNjNlMDg0NTk4In0.97HttqdtxAdwdJ-dx2_iTfcpeOoOR4hD0ufDRD2PqRk', '深圳市眼科医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', a497b3c8-f835-11e7-aabb-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4Mjk0OTMsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiYTQ5N2IzYzgtZjgzNS0xMWU3LWFhYmItMDAxNjNlMDg0NTk4In0.QW-1jOmtPDZrJL_3XajikVBscHaxpLloVN75McXN_Gs', '深圳市眼科医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', a502e278-f835-11e7-aabc-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4Mjk0OTQsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiYTUwMmUyNzgtZjgzNS0xMWU3LWFhYmMtMDAxNjNlMDg0NTk4In0.jk0mkjV3ipNtE6a3jYxBigzVeU8qTA6bvMmvlCZn-G8', '深圳市眼科医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', a5628c8c-f835-11e7-aabd-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4Mjk0OTQsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiYTU2MjhjOGMtZjgzNS0xMWU3LWFhYmQtMDAxNjNlMDg0NTk4In0.p3c4Ezpesf0U3d0NcFHcSrGr9TMFgqFb2PsU8C1vlgA', '深圳市眼科医院', 'device');
INSERT INTO "manager"."clients_by_user" ("user", "id", "access_key", "name", "type") VALUES ('yangtao@email.com', a5bf1806-f835-11e7-aabe-00163e084598, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTU4Mjk0OTUsImlzcyI6Im1haW5mbHV4Iiwic3ViIjoiYTViZjE4MDYtZjgzNS0xMWU3LWFhYmUtMDAxNjNlMDg0NTk4In0.7Hb3BXaqEjFc_GiNRTjxdYH7IsXQaCX_WYB1_MEAH6A', '深圳市眼科医院', 'device');
INSERT INTO "manager"."channels_by_user" ("user", "id", "connected", "name") VALUES ('yangtao@email.com', 432c1105-f83b-11e7-aabf-00163e084598, {'01639ab2-f835-11e7-992d-00163e084598', '0a1de8de-f835-11e7-992e-00163e084598', '32952fe1-f835-11e7-992f-00163e084598', 'a414b784-f835-11e7-aaba-00163e084598', 'a497b3c8-f835-11e7-aabb-00163e084598', 'a502e278-f835-11e7-aabc-00163e084598', 'a5628c8c-f835-11e7-aabd-00163e084598', 'a5bf1806-f835-11e7-aabe-00163e084598'}, 'chan1');
INSERT INTO "manager"."channels_by_user" ("user", "id", "connected", "name") VALUES ('yangtao@email.com', 0901534b-f83c-11e7-aac0-00163e084598, {'4ddf6970-f835-11e7-9930-00163e084598', '54931a27-f835-11e7-9931-00163e084598', '57dbfd30-f835-11e7-9932-00163e084598', '9a95d806-f835-11e7-aab7-00163e084598', '9b53cc9e-f835-11e7-aab8-00163e084598', '9c0d7645-f835-11e7-aab9-00163e084598'}, 'chan2');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '01639ab2-f835-11e7-992d-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', 0, 0, 1515847239392, '', 0, False, '', '', '1', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '0a1de8de-f835-11e7-992e-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', 0, 0, 1515847245002, '', 0, False, '', '', '2', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', '32952fe1-f835-11e7-992f-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', 0, 0, 1515847251842, '', 0, False, '', '', '3', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', 'a414b784-f835-11e7-aaba-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', 0, 0, 1515847257450, '', 0, False, '', '', '4', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', 'a497b3c8-f835-11e7-aabb-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', 0, 0, 1515847264260, '', 0, False, '', '', '5', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', 'a502e278-f835-11e7-aabc-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', 0, 0, 1515847269958, '', 0, False, '', '', '6', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', 'a5628c8c-f835-11e7-aabd-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', 0, 0, 1515847275969, '', 0, False, '', '', '7', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市眼科医院', 'a5bf1806-f835-11e7-aabe-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', 0, 0, 1515847283041, '', 0, False, '', '', '8', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '4ddf6970-f835-11e7-9930-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', 0, 0, 1515847170482, '', 0, False, '', '', '1', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '54931a27-f835-11e7-9931-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', 0, 0, 1515847176851, '', 0, False, '', '', '2', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '57dbfd30-f835-11e7-9932-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', 0, 0, 1515847184104, '', 0, False, '', '', '3', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '9a95d806-f835-11e7-aab7-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', 0, 0, 1515847200420, '', 0, False, '', '', '4', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '9b53cc9e-f835-11e7-aab8-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', 0, 0, 1515847207540, '', 0, False, '', '', '5', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."device_info" ("hospitalname", "deviceid", "channelid", "availablepowersource", "batterylevel", "currenttime", "devicetype", "errorcode", "factoryreset", "firmwareversion", "hardwareversion", "hospitaldeviceid", "manufacturer", "memoryfree", "modelnumber", "powersourcesurrent", "powersourcevoltage", "reboot", "reseterrorcode", "softwareversion", "supportedbindingmodes", "timezone", "utcoffset") VALUES ('深圳市人民医院', '9c0d7645-f835-11e7-aab9-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', 0, 0, 1515847216230, '', 0, False, '', '', '6', '', 0, '', 0, 0, False, '', '', '', '', '');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '01639ab2-f835-11e7-992d-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', '1', '1', '眼科', '', '张三', '男');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '0a1de8de-f835-11e7-992e-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', '2', '2', '眼科', '', '张士', '男');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', '32952fe1-f835-11e7-992f-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', '1', '1', '口腔', '', '张三', '女');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市眼科医院', 'a414b784-f835-11e7-aaba-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598', '2', '2', '口腔', '', '张国光', '女');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid") VALUES ('深圳市眼科医院', 'a497b3c8-f835-11e7-aabb-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid") VALUES ('深圳市眼科医院', 'a502e278-f835-11e7-aabc-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid") VALUES ('深圳市眼科医院', 'a5628c8c-f835-11e7-aabd-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid") VALUES ('深圳市眼科医院', 'a5bf1806-f835-11e7-aabe-00163e084598', '432c1105-f83b-11e7-aabf-00163e084598');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '4ddf6970-f835-11e7-9930-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', '1', '1', '内科', '', '王一', '男');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '54931a27-f835-11e7-9931-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', '2', '2', '内科', '', '李四', '女');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '57dbfd30-f835-11e7-9932-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', '1', '1', '外科', '', '可口', '女');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '9a95d806-f835-11e7-aab7-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', '2', '2', '外科', '', '好好', '女');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '9b53cc9e-f835-11e7-aab8-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', '3', '3', '内科', '', '哈皮哈', '男');
INSERT INTO "manager"."hospital_by_patient" ("hospitalname", "deviceid", "channelid", "hospitalbed", "hospitaldeviceid", "hospitalzone", "patientid", "patientname", "patientsex") VALUES ('深圳市人民医院', '9c0d7645-f835-11e7-aab9-00163e084598', '0901534b-f83c-11e7-aac0-00163e084598', '3', '3', '外科', '', '娄娄', '女');
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
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (0901534b-f83c-11e7-aac0-00163e084598, 7d6bb7d0-f86b-11e7-bae4-8f77f3976232, '', '54931a27-f835-11e7-9931-00163e084598:oxgen', 'http', '0901534b-f83c-11e7-aac0-00163e084598', 0.0, 1259090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (0901534b-f83c-11e7-aac0-00163e084598, 7d6b69b0-f86b-11e7-bae4-8f77f3976232, '', '54931a27-f835-11e7-9931-00163e084598:pressure', 'http', '0901534b-f83c-11e7-aac0-00163e084598', 0.0, 1259090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (0901534b-f83c-11e7-aac0-00163e084598, 7d6b1b90-f86b-11e7-bae4-8f77f3976232, '', '54931a27-f835-11e7-9931-00163e084598:puls', 'http', '0901534b-f83c-11e7-aac0-00163e084598', 0.0, 1259090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (0901534b-f83c-11e7-aac0-00163e084598, 64419c21-f86b-11e7-bae4-8f77f3976232, '', '57dbfd30-f835-11e7-9932-00163e084598:oxgen', 'http', '0901534b-f83c-11e7-aac0-00163e084598', 0.0, 1258090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (0901534b-f83c-11e7-aac0-00163e084598, 64419c20-f86b-11e7-bae4-8f77f3976232, '', '57dbfd30-f835-11e7-9932-00163e084598:pressure', 'http', '0901534b-f83c-11e7-aac0-00163e084598', 0.0, 1258090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (0901534b-f83c-11e7-aac0-00163e084598, 64414e00-f86b-11e7-bae4-8f77f3976232, '', '57dbfd30-f835-11e7-9932-00163e084598:puls', 'http', '0901534b-f83c-11e7-aac0-00163e084598', 0.0, 1258090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (0901534b-f83c-11e7-aac0-00163e084598, 3a2630e0-f86b-11e7-bae4-8f77f3976232, '', '4ddf6970-f835-11e7-9930-00163e084598:oxgen', 'http', '0901534b-f83c-11e7-aac0-00163e084598', 0.0, 1258090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (0901534b-f83c-11e7-aac0-00163e084598, 3a2609d0-f86b-11e7-bae4-8f77f3976232, '', '4ddf6970-f835-11e7-9930-00163e084598:pressure', 'http', '0901534b-f83c-11e7-aac0-00163e084598', 0.0, 1258090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (0901534b-f83c-11e7-aac0-00163e084598, 3a25e2c0-f86b-11e7-bae4-8f77f3976232, '', '4ddf6970-f835-11e7-9930-00163e084598:puls', 'http', '0901534b-f83c-11e7-aac0-00163e084598', 0.0, 1258090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a2100c1-f86b-11e7-bae4-8f77f3976232, '', 'a414b784-f835-11e7-aaba-00163e084598:oxgen', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1258090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a2100c0-f86b-11e7-bae4-8f77f3976232, '', 'a414b784-f835-11e7-aaba-00163e084598:pressure', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1258090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a20d9b0-f86b-11e7-bae4-8f77f3976232, '', 'a414b784-f835-11e7-aaba-00163e084598:puls', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1258090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a1c1ec0-f86b-11e7-bae4-8f77f3976232, '', 'a414b784-f835-11e7-aaba-00163e084598:oxgen', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1269090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a1bf7b1-f86b-11e7-bae4-8f77f3976232, '', 'a414b784-f835-11e7-aaba-00163e084598:pressure', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1269090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a1bf7b0-f86b-11e7-bae4-8f77f3976232, '', 'a414b784-f835-11e7-aaba-00163e084598:puls', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1269090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a16eea1-f86b-11e7-bae4-8f77f3976232, '', 'a414b784-f835-11e7-aaba-00163e084598:oxgen', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1259090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a16eea0-f86b-11e7-bae4-8f77f3976232, '', 'a414b784-f835-11e7-aaba-00163e084598:pressure', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1259090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a16c790-f86b-11e7-bae4-8f77f3976232, '', 'a414b784-f835-11e7-aaba-00163e084598:puls', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1259090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a11e591-f86b-11e7-bae4-8f77f3976232, '', '0a1de8de-f835-11e7-992e-00163e084598:oxgen', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1289090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a11e590-f86b-11e7-bae4-8f77f3976232, '', '0a1de8de-f835-11e7-992e-00163e084598:pressure', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1289090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a117060-f86b-11e7-bae4-8f77f3976232, '', '0a1de8de-f835-11e7-992e-00163e084598:puls', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1289090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a0cb571-f86b-11e7-bae4-8f77f3976232, '', '0a1de8de-f835-11e7-992e-00163e084598:oxgen', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1288090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a0cb570-f86b-11e7-bae4-8f77f3976232, '', '0a1de8de-f835-11e7-992e-00163e084598:pressure', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1288090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a0c8e60-f86b-11e7-bae4-8f77f3976232, '', '0a1de8de-f835-11e7-992e-00163e084598:puls', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1288090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a07ac60-f86b-11e7-bae4-8f77f3976232, '', '0a1de8de-f835-11e7-992e-00163e084598:oxgen', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1285090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a078550-f86b-11e7-bae4-8f77f3976232, '', '0a1de8de-f835-11e7-992e-00163e084598:pressure', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1285090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 3a05d7a0-f86b-11e7-bae4-8f77f3976232, '', '0a1de8de-f835-11e7-992e-00163e084598:puls', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1285090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 39fed2c0-f86b-11e7-bae4-8f77f3976232, '', '01639ab2-f835-11e7-992d-00163e084598:oxgen', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1287090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 39feabb0-f86b-11e7-bae4-8f77f3976232, '', '01639ab2-f835-11e7-992d-00163e084598:pressure', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1287090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 39fe84a0-f86b-11e7-bae4-8f77f3976232, '', '01639ab2-f835-11e7-992d-00163e084598:puls', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1287090086.001, 'V', 0.0, 120.1, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 39f758b0-f86b-11e7-bae4-8f77f3976232, '', '01639ab2-f835-11e7-992d-00163e084598:oxgen', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1286090082.001, 'A', 0.0, 1.3, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 39f731a0-f86b-11e7-bae4-8f77f3976232, '', '01639ab2-f835-11e7-992d-00163e084598:pressure', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1286090081.001, 'A', 0.0, 1.2, False, '', '');
INSERT INTO "message_writer"."messages_by_channel" ("channel", "id", "l", "n", "protocol", "publisher", "s", "t", "u", "ut", "v", "vb", "vd", "vs") VALUES (432c1105-f83b-11e7-aabf-00163e084598, 39f70a90-f86b-11e7-bae4-8f77f3976232, '', '01639ab2-f835-11e7-992d-00163e084598:puls', 'http', '432c1105-f83b-11e7-aabf-00163e084598', 0.0, 1286090086.001, 'V', 0.0, 120.1, False, '', '');
