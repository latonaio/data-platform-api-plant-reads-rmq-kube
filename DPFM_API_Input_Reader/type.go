package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	General           General  `json:"Plant"`
	Generals          Generals `json:"Plants"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
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

type Generals []struct {
	BusinessPartner *int    `json:"BusinessPartner"`
	Plant           *string `json:"Plant"`
	Language        *string `json:"Language"`
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
