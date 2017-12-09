package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/gocql/gocql"
)

const (
	sep         string = ","
	defCluster  string = "127.0.0.1"
	mgrKeyspace string = "manager"
	mgrSecret   string = "manager"
	msgKeyspace string = "message_writer"
	msgSecret   string = "message_writer"
)

/*SessionMgr cassandra db manager keyspace session */
var SessionMgr *gocql.Session

/*SessionMsg cassandra db message keyspace session */
var SessionMsg *gocql.Session

/*InitDB init the Cassandra db session*/
func InitDB() error {

	log := logs.GetBeeLogger()
	log.Info("init DB ...")

	mgrCluster := gocql.NewCluster(defCluster)
	mgrCluster.Keyspace = mgrKeyspace
	mgrCluster.Consistency = gocql.Quorum
	Session, err := mgrCluster.CreateSession()

	if err != nil {
		log.Info("init cassandra manager session failure ")
		return err
	}
	for _, table := range tables {
		if err := Session.Query(table).Exec(); err != nil {
			log.Info("create table in cassandra failure ")
			//		return err
		}
	}

	SessionMgr = Session

	msgCluster := gocql.NewCluster(defCluster)
	msgCluster.Keyspace = msgKeyspace
	msgCluster.Consistency = gocql.Quorum
	Session, err = msgCluster.CreateSession()
	if err != nil {
		log.Info("init cassandra message session failure ")
		return err
	}
	SessionMsg = Session

	log.Info("connected to cassandra ...")
	return nil
}
