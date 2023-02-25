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
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.PlantFoundationDate,
			&pm.PlantLiquidationDate,
			&pm.SearchTerm1,
			&pm.SearchTerm2,
			&pm.PlantDeathDate,
			&pm.PlantIsBlocked,
			&pm.GroupPlantName1,
			&pm.GroupPlantName2,
			&pm.AddressID,
			&pm.Country,
			&pm.TimeZone,
			&pm.PlantIDByExtSystem,
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
			CreationDate:         data.CreationDate,
			CreationTime:         data.CreationTime,
			LastChangeDate:       data.LastChangeDate,
			LastChangeTime:       data.LastChangeTime,
			PlantFoundationDate:  data.PlantFoundationDate,
			PlantLiquidationDate: data.PlantLiquidationDate,
			SearchTerm1:          data.SearchTerm1,
			SearchTerm2:          data.SearchTerm2,
			PlantDeathDate:       data.PlantDeathDate,
			PlantIsBlocked:       data.PlantIsBlocked,
			GroupPlantName1:      data.GroupPlantName1,
			GroupPlantName2:      data.GroupPlantName2,
			AddressID:            data.AddressID,
			Country:              data.Country,
			TimeZone:             data.TimeZone,
			PlantIDByExtSystem:   data.PlantIDByExtSystem,
			IsMarkedForDeletion:  data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &generals, nil
	}

	return &generals, nil
}

func ConvertToStorageLocation(sdc *api_input_reader.SDC, rows *sql.Rows) (*StorageLocation, error) {
	defer rows.Close()
	pm := &requests.StorageLocation{}

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.StorageLocationFullName,
			&pm.StorageLocationName,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.SearchTerm1,
			&pm.SearchTerm2,
			&pm.StorageLocationIsBlocked,
			&pm.GroupStorageLocationName1,
			&pm.GroupStorageLocationName2,
			&pm.StorageLocationIDByExtSystem,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &StorageLocation{}, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &StorageLocation{}, nil
	}

	data := pm
	storageLocation := &StorageLocation{
		BusinessPartner:              data.BusinessPartner,
		Plant:                        data.Plant,
		StorageLocation:              data.StorageLocation,
		StorageLocationFullName:      data.StorageLocationFullName,
		StorageLocationName:          data.StorageLocationName,
		CreationDate:                 data.CreationDate,
		CreationTime:                 data.CreationTime,
		LastChangeDate:               data.LastChangeDate,
		LastChangeTime:               data.LastChangeTime,
		SearchTerm1:                  data.SearchTerm1,
		SearchTerm2:                  data.SearchTerm2,
		StorageLocationIsBlocked:     data.StorageLocationIsBlocked,
		GroupStorageLocationName1:    data.GroupStorageLocationName1,
		GroupStorageLocationName2:    data.GroupStorageLocationName2,
		StorageLocationIDByExtSystem: data.StorageLocationIDByExtSystem,
		IsMarkedForDeletion:          data.IsMarkedForDeletion,
	}
	return storageLocation, nil
}
