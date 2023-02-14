package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	General         *General         `json:"General"`
	Generals        *[]General       `json:"Generals"`
	StorageLocation *StorageLocation `json:"StorageLocation"`
}

type Plant struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
}

type General struct {
	BusinessPartner      int             `json:"BusinessPartner"`
	Plant                string          `json:"Plant"`
	PlantFullName        *string         `json:"PlantFullName"`
	PlantName            *string         `json:"PlantName"`
	Language             *string         `json:"Language"`
	CreationDate         *string         `json:"CreationDate"`
	CreationTime         *string         `json:"CreationTime"`
	LastChangeDate       *string         `json:"LastChangeDate"`
	LastChangeTime       *string         `json:"LastChangeTime"`
	PlantFoundationDate  *string         `json:"PlantFoundationDate"`
	PlantLiquidationDate *string         `json:"PlantLiquidationDate"`
	SearchTerm1          *string         `json:"SearchTerm1"`
	SearchTerm2          *string         `json:"SearchTerm2"`
	PlantDeathDate       *string         `json:"PlantDeathDate"`
	PlantIsBlocked       *bool           `json:"PlantIsBlocked"`
	GroupPlantName1      *string         `json:"GroupPlantName1"`
	GroupPlantName2      *string         `json:"GroupPlantName2"`
	AddressID            *int            `json:"AddressID"`
	Country              *string         `json:"Country"`
	TimeZone             *string         `json:"TimeZone"`
	PlantIDByExtSystem   *string         `json:"PlantIDByExtSystem"`
	IsMarkedForDeletion  *bool           `json:"IsMarkedForDeletion"`
	StorageLocation      StorageLocation `json:"StorageLocation"`
}

type StorageLocation struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	StorageLocation              string  `json:"StorageLocation"`
	StorageLocationFullName      *string `json:"StorageLocationFullName"`
	StorageLocationName          *string `json:"StorageLocationName"`
	CreationDate                 *string `json:"CreationDate"`
	CreationTime                 *string `json:"CreationTime"`
	LastChangeDate               *string `json:"LastChangeDate"`
	LastChangeTime               *string `json:"LastChangeTime"`
	SearchTerm1                  *string `json:"SearchTerm1"`
	SearchTerm2                  *string `json:"SearchTerm2"`
	StorageLocationIsBlocked     *bool   `json:"StorageLocationIsBlocked"`
	GroupStorageLocationName1    *string `json:"GroupStorageLocationName1"`
	GroupStorageLocationName2    *string `json:"GroupStorageLocationName2"`
	StorageLocationIDByExtSystem *string `json:"StorageLocationIDByExtSystem"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
}
