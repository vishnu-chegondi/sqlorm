package ormdrivers

import (
	"database/sql/driver"
	"time"

	mysql "github.com/go-sql-driver/mysql"
)

func mysqlConfig() *mysql.Config {
	// creates a new Config with default Values
	userConfig := mysql.NewConfig()

	// Assign Mandatory Values Attached in settings package
	if User != "" {
		userConfig.User = User
	} else {
		panic("User can not be empty to create db connection")
	}
	if Passwd != "" {
		userConfig.Passwd = Passwd
	} else {
		panic("Passwd can not be empty to create db connection")
	}
	if Addr != "" {
		userConfig.Addr = Addr
	} else {
		panic("Addr can not be empty to create db connection")
	}
	if DBName != "" {
		userConfig.DBName = DBName
	} else {
		panic("DBName can not be empty to create db connection")
	}
	if Net != "" {
		userConfig.Net = Net
	} else {
		panic("Net can not be empty to create db connection")
	}

	// Assign Values if specified in settings package else use default
	// if Params != map[] {
	// 	userConfig.Params = setParams
	// }
	if Collation != "" {
		userConfig.Collation = Collation
	}
	if Loc != time.UTC {
		userConfig.Loc = Loc
	}
	if MaxAllowedPacket != 0 {
		userConfig.MaxAllowedPacket = MaxAllowedPacket
	}
	if ServerPubKey != "" {
		userConfig.ServerPubKey = ServerPubKey
	}
	if TLSConfig != "" {
		userConfig.TLSConfig = TLSConfig
	}
	if Timeout != 0*time.Second {
		userConfig.Timeout = Timeout
	}
	if ReadTimeout != 0*time.Second {
		userConfig.ReadTimeout = ReadTimeout
	}
	if WriteTimeout != 0*time.Second {
		userConfig.WriteTimeout = WriteTimeout
	}
	if AllowAllFiles != false {
		userConfig.AllowAllFiles = AllowAllFiles
	}
	if AllowCleartextPasswords != false {
		userConfig.AllowCleartextPasswords = AllowCleartextPasswords
	}
	if AllowNativePasswords != true {
		userConfig.AllowNativePasswords = AllowNativePasswords
	}
	if AllowOldPasswords != false {
		userConfig.AllowOldPasswords = AllowOldPasswords
	}
	if CheckConnLiveness != true {
		userConfig.CheckConnLiveness = CheckConnLiveness
	}
	if ClientFoundRows != false {
		userConfig.ClientFoundRows = ClientFoundRows
	}
	if ColumnsWithAlias != false {
		userConfig.ColumnsWithAlias = ColumnsWithAlias
	}
	if InterpolateParams != false {
		userConfig.InterpolateParams = InterpolateParams
	}
	if MultiStatements != false {
		userConfig.MultiStatements = MultiStatements
	}
	if ParseTime != false {
		userConfig.ParseTime = ParseTime
	}
	if RejectReadOnly != false {
		userConfig.RejectReadOnly = RejectReadOnly
	}
	return userConfig
}

func mysqlConnector() (driver.Connector, error) {
	config := mysqlConfig()
	connector, err := mysql.NewConnector(config)
	return connector, err
}
