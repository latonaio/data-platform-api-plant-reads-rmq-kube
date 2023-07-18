package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-plant-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-plant-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *dpfm_api_output_formatter.General
	var storageLocation *dpfm_api_output_formatter.StorageLocation
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		case "Generals":
			func() {
				general = c.Generals(mtx, input, output, errs, log)
			}()
		case "GeneralsByBusinessPartners":
			func() {
				general = c.GeneralsByBusinessPartners(mtx, input, output, errs, log)
			}()
		case "StorageLocation":
			func() {
				storageLocation = c.StorageLocation(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:         general,
		StorageLocation: storageLocation,
	}

	return data
}

func (c *DPFMAPICaller) General(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	where := "WHERE 1 = 1"

	if input.General.BusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND BusinessPartner = %v", where, *input.General.BusinessPartner)
	}

	if input.General.Plant != nil {
		where = fmt.Sprintf("%s\nAND Plant = \"%s\"", where, *input.General.Plant)
	}

	if input.General.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.General.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_plant_general_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return &((*data)[0])
}

func (c *DPFMAPICaller) Generals(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	where := "WHERE 1 = 1"

	if input.General.BusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND BusinessPartner = %v", where, *input.General.BusinessPartner)
	}

	if input.General.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.General.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_plant_general_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return &((*data)[0])
}

func (c *DPFMAPICaller) GeneralsByBusinessPartners(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	log.Info("GeneralsByBusinessPartners")
	in := ""

	for iGeneral, vGeneral := range input.Generals {
		businessPartner := vGeneral.BusinessPartner
		if iGeneral == 0 {
			in = fmt.Sprintf(
				"( '%d' )",
				businessPartner,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%d' )",
			in,
			businessPartner,
		)
	}

	where := fmt.Sprintf(" WHERE ( BusinessPartner ) IN ( %s ) ", in)
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_plant_general_data
		` + where + ` ;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) StorageLocation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.StorageLocation {
	plant := input.General.Plant
	businessPartner := input.General.BusinessPartner
	storageLocation := input.General.StorageLocation.StorageLocation

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Plant, StorageLocation,
		StorageLocationFullName, StorageLocationName,
		StorageLocationIDByExtSystem, 
		CreationDate, LastChangeDate,
		IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_plant_storage_location_data
		WHERE (BusinessPartner, Plant, StorageLocation) = (?, ?, ?);`, businessPartner, plant, storageLocation,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToStorageLocation(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
