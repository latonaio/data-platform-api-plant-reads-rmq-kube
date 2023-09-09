package dpfm_api_output_formatter

import (
	"data-platform-api-plant-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-plant-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	generals := make([]General, 0, len(sdc.Generals))

	i := 0
	for rows.Next() {
		i++
		pm := &requests.General{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.PlantFullName,
			&pm.PlantName,
			&pm.Language,
			&pm.PlantFoundationDate,
			&pm.PlantLiquidationDate,
			&pm.PlantDeathDate,
			&pm.AddressID,
			&pm.Country,
			&pm.TimeZone,
			&pm.PlantIDByExtSystem,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &generals, err
		}

		data := pm

		generals = append(generals, General{
			BusinessPartner:      data.BusinessPartner,
			Plant:                data.Plant,
			PlantFullName:        data.PlantFullName,
			PlantName:            data.PlantName,
			Language:             data.Language,
			PlantFoundationDate:  data.PlantFoundationDate,
			PlantLiquidationDate: data.PlantLiquidationDate,
			PlantDeathDate:       data.PlantDeathDate,
			AddressID:            data.AddressID,
			Country:              data.Country,
			TimeZone:             data.TimeZone,
			PlantIDByExtSystem:   data.PlantIDByExtSystem,
			CreationDate:         data.CreationDate,
			LastChangeDate:       data.LastChangeDate,
			IsMarkedForDeletion:  data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &generals, nil
	}

	return &generals, nil
}

func ConvertToStorageLocation(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]StorageLocation, error) {
	defer rows.Close()
	pm := &requests.StorageLocation{}
	storageLocations := make([]StorageLocation, 0, len(sdc.Generals))

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.StorageLocationFullName,
			&pm.StorageLocationName,
			&pm.StorageLocationIDByExtSystem,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &storageLocations, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &storageLocations, nil
	}

	data := pm

	storageLocations = append(storageLocations, StorageLocation{
			BusinessPartner:              data.BusinessPartner,
			Plant:                        data.Plant,
			StorageLocation:              data.StorageLocation,
			StorageLocationFullName:      data.StorageLocationFullName,
			StorageLocationName:          data.StorageLocationName,
			StorageLocationIDByExtSystem: data.StorageLocationIDByExtSystem,
			CreationDate:                 data.CreationDate,
			LastChangeDate:               data.LastChangeDate,
			IsMarkedForDeletion:          data.IsMarkedForDeletion,
	})

	return &storageLocations, nil
}
