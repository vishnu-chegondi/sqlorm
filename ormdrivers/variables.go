package ormdrivers

import (
	"time"
)

var (
	DriverName       string            = "mysql"
	User             string            = "root"           // Username
	Passwd           string            = "MYSQLpassword1" // Password (requires User)
	Net              string            = "tcp"            // Network type
	Addr             string            = "localhost:3306" // Network address (requires Net)
	DBName           string            = "sqlorm"         // Database name
	Params           map[string]string                    // Connection parameters
	Collation        string                               // Connection collation
	Loc              *time.Location                       // Location for time.Time values
	MaxAllowedPacket int                                  // Max packet size allowed
	ServerPubKey     string                               // Server public key name

	TLSConfig string // TLS configuration name

	Timeout      time.Duration // Dial timeout
	ReadTimeout  time.Duration // I/O read timeout
	WriteTimeout time.Duration // I/O write timeout

	AllowAllFiles           bool        // Allow all files to be used with LOAD DATA LOCAL INFILE
	AllowCleartextPasswords bool        // Allows the cleartext client side plugin
	AllowNativePasswords    bool = true // Allows the native password authentication method
	AllowOldPasswords       bool        // Allows the old insecure password method
	CheckConnLiveness       bool = true // Check connections for liveness before using them
	ClientFoundRows         bool        // Return number of matching rows instead of rows changed
	ColumnsWithAlias        bool        // Prepend table alias to column names
	InterpolateParams       bool        // Interpolate placeholders into query string
	MultiStatements         bool        // Allow multiple statements in one query
	ParseTime               bool        // Parse time values to time.Time
	RejectReadOnly          bool        // Reject read-only connections
	// contains filtered or unexported fields

)
