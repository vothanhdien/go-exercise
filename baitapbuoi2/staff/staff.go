package staff

import (
	"errors"
	"sort"
)

type Staff struct {
	Name               string
	Coefficientssalary float64
	Subsidy            float64
}

var SalaryBasic float64 = 1500000

func SortName(s []Staff) ([]Staff, error) {
	if s != nil {
		sort.SliceStable(s, func(i, j int) bool { return s[i].Name < s[j].Name })
		return s, nil
	}
	return []Staff{}, errors.New("empty slice")
}

func Salary(s []Staff) ([]Staff, error) {
	if s != nil {
		sort.SliceStable(s, func(i, j int) bool {
			return (s[i].Coefficientssalary*SalaryBasic + s[i].Subsidy) < (s[j].Coefficientssalary*SalaryBasic + s[j].Subsidy)
		})
		return s, nil
	}
	return []Staff{}, errors.New("empty slice")
}

func Max2Salary(s []Staff) ([]Staff, error) {
	if s != nil {
		sliceSalarySort := []float64{}
		keyListbool := make(map[float64]bool)
		listRemoveDupSalary := []float64{}
		sliceMax2Salary := []Staff{}
		for _, item := range s {
			sliceSalarySort = append(sliceSalarySort, (item.Coefficientssalary*SalaryBasic + item.Subsidy))
		}
		for _, v := range sliceSalarySort {
			if value := keyListbool[v]; !value {
				keyListbool[v] = true
				listRemoveDupSalary = append(listRemoveDupSalary, v)
			}
		}
		secondMaxSalary := listRemoveDupSalary[len(listRemoveDupSalary)-2]
		for _, item := range s {
			if salaryItem := item.Coefficientssalary*SalaryBasic + item.Subsidy; salaryItem == secondMaxSalary {
				sliceMax2Salary = append(sliceMax2Salary, item)
			}
		}
		return sliceMax2Salary, nil

	}
	return []Staff{}, errors.New("empty slice")
}
