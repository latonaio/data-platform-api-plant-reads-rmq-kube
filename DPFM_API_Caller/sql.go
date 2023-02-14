package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-plant-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-plant-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
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
	var generals *[]dpfm_api_output_formatter.General
	var storageLocation *dpfm_api_output_formatter.StorageLocation
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		case "Generals":
			func() {
				generals = c.Generals(mtx, input, output, errs, log)
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
		Generals:        generals,
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
) *dpfm_api_output_formatter.General {
	plant := input.General.Plant
	businessPartner := input.General.BusinessPartner

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Plant, PlantFullName, PlantName, Language, CreationDate, CreationTime, 
		LastChangeDate, LastChangeTime, PlantFoundationDate, PlantLiquidationDate, SearchTerm1, SearchTerm2, 
		PlantDeathDate, PlantIsBlocked, GroupPlantName1, GroupPlantName2, AddressID, Country, TimeZone, 
		PlantIDByExtSystem, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_plant_general_data
		WHERE (BusinessPartner, Plant) = (?, ?);`, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

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
	var args []interface{}
	generals := input.Generals

	cnt := 0
	for _, v := range generals {
		args = append(args, v.BusinessPartner, v.Plant)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Plant, PlantFullName, PlantName, Language, CreationDate, CreationTime, 
		LastChangeDate, LastChangeTime, PlantFoundationDate, PlantLiquidationDate, SearchTerm1, SearchTerm2, 
		PlantDeathDate, PlantIsBlocked, GroupPlantName1, GroupPlantName2, AddressID, Country, TimeZone, 
		PlantIDByExtSystem, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_plant_general_data
		WHERE (BusinessPartner, Plant)  IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToGeneral(input, rows)
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
		`SELECT BusinessPartner, Plant, StorageLocation, StorageLocationFullName, StorageLocationName, CreationDate, 
		CreationTime, LastChangeDate, LastChangeTime, SearchTerm1, SearchTerm2, StorageLocationIsBlocked, 
		GroupStorageLocationName1, GroupStorageLocationName2, StorageLocationIDByExtSystem, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_plant_storage_location_data
		WHERE (BusinessPartner, Plant, StorageLocation) = (?, ?, ?);`, businessPartner, plant, storageLocation,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToStorageLocation(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
