package script

import (
	"testing"
)

func TestSeedStudentToGroup(t *testing.T) {
	// type studentToGroupData struct {
	// 	AcademicYearName string `json:"academic_year_name"`
	// 	StudentID        string `json:"student_id"`
	// 	GroupName        string `json:"group_name"`
	// }
	// var relations []studentToGroupData
	// {
	// 	// 48 Medical students
	// 	// ID from 1 -> 48
	// 	// 8 groups, each group from 4-5 students
	// 	noOfMDStudents := 48
	// 	noOfMDGroups := 8
	// 	for i := 1; i <= noOfMDStudents; i++ {
	// 		studentID := strconv.Itoa(i)
	// 		for len(studentID) < 3 {
	// 			studentID = "0" + studentID
	// 		}
	// 		studentID = "V" + studentID
	// 		relations = append(relations, studentToGroupData{
	// 			AcademicYearName: "2023-2024 MD Program",
	// 			StudentID:        studentID,
	// 			GroupName:        fmt.Sprintf("Medical Doctor Group %v", i%noOfMDGroups+1),
	// 		})
	// 	}
	// }
	// {
	// 	// 15 Nursing students
	// 	// ID from 51 -> 65
	// 	noOfNursingStudents := 15
	// 	noOfNursingGroups := 5
	// 	for i := 1; i <= noOfNursingStudents; i++ {
	// 		studentID := strconv.Itoa(i + 50)
	// 		for len(studentID) < 3 {
	// 			studentID = "0" + studentID
	// 		}
	// 		studentID = "V" + studentID
	// 		relations = append(relations, studentToGroupData{
	// 			AcademicYearName: "2023-2024 Nursing Program",
	// 			StudentID:        studentID,
	// 			GroupName:        fmt.Sprintf("Nursing Group %v", i%noOfNursingGroups+1),
	// 		})
	// 	}
	// }
	// res, err := json.Marshal(relations)
	// assert.Nil(t, err)
	// fmt.Println(string(res))
	ClearDataDBMigration("file://../../db/migration", config.DBUrl)
	SeedStudentData(config.BasePath, server)
	SeedAcademicYearData(config.BasePath, server)
	SeedGroupData(config.BasePath, server)
	SeedStudentToGroupData(config.BasePath, server)
}
