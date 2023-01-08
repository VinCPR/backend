package api

// func (server *Server) studentCalendar(ctx *gin.Context) {
// 	academicYearName := ctx.Query("academicYearName")
// 	academicYear, err := server.store.GetAcademicYearByName(ctx, academicYearName)
// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			log.Info().Msgf("cannot find academic year %v", academicYearName)
// 			ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
//
// 	studentID := ctx.Query("studentID")
// 	student, err := server.store.GetStudentByStudentId(ctx, studentID)
// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
//
// 	studentToGroups, err := server.store.GetStudentToGroupByAcademicYearIDAndStudentID(ctx,
// 		db.GetStudentToGroupByAcademicYearIDAndStudentIDParams{
// 			AcademicYearID: academicYear.ID,
// 			StudentID:      student.ID,
// 		})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
//
// 	groupIDs := make([]int64, 0)
// 	for _, studentToGroup := range studentToGroups {
// 		groupIDs = append(groupIDs, studentToGroup.GroupID)
// 	}
// 	clinicalRotationEvents, err := server.store.ListRotationEventsByAcademicYearIDAndGroupID(ctx,
// 		db.ListRotationEventsByAcademicYearIDAndGroupIDParams{
// 			AcademicYearID: academicYear.ID,
// 			GroupIds:       groupIDs,
// 		})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// }
