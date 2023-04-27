package config

type Config struct {
	Version string
	Server  struct {
		Host string
		Port int
	}
	DB_PGSQL  string
	BotParams struct {
		Token string
	}
	Sql struct {
		Select struct {
			GetNumberOfUsers    string
			GetUsers            string
			GetPhone            string
			GetListenCallTest   string
			GetWrongMessageTest string
			GetUid              string
			GetCalls            string
			GetGetPhoto         string
		}
		Insert struct {
			AddUser string
		}
		Update struct {
			TimeBegin      string
			Forward        string
			TimeEnd        string
			ErrorUpdate    string
			UpdateGetPhoto string
			UpdateStatus   string
		}
	}
	ErrorRu struct {
		GetNumberOfUsers string
		AddUser          string
	}
}

type User struct {
	Name     string
	JobTitle string
	ChatID   int64
}
