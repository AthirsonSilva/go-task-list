package handlers

import (
	"fmt"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/authentication"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/internal/api"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"net/http"
	"os"
	"strings"
)

// ExportToPdf @Summary Export to PDF
//
//	@Tags		users
//	@Produce	json
//	@Success	200				{object}	api.Response
//	@Failure	500				{object}	api.Response
//	@Failure	500				{object}	api.Exception
//	@Failure	400				{object}	api.Exception
//	@Failure	429				{object}	api.Exception
//	@Param		Authorization	header		string	true	"Authorization"
//	@Router		/api/v1/users/export-pdf [get]
func ExportToPdf(res http.ResponseWriter, req *http.Request) {
	logger.Info("ExportToPdf", "Exporting user to PDF")

	user, err := authentication.GetUserFromToken(req)
	if err != nil {
		logger.Error("ExportToPdf", err.Error())
		api.Error(res, req, "Error while getting user from token", err, http.StatusInternalServerError)
		return
	}

	pdfGenerator, err := pdf.NewPDFGenerator()
	if err != nil {
		logger.Error("ExportToPdf", err.Error())
		api.Error(res, req, "Error while creating PDF generator", err, http.StatusInternalServerError)
		return
	}

	htmlStr := fmt.Sprintf(`<html>
				<body>
					<h1>User information</h1>
					<img src="%s" alt="img" height="42" width="42">
					<p>Username: %s</p>		
					<p>Email: %s</p>
					<p>Created At: %s</p>
				</body>
				</html>`, user.PhotoUrl, user.Username, user.Email, user.CreatedAt.String())

	pdfGenerator.Orientation.Set(pdf.OrientationPortrait)
	pdfGenerator.PageSize.Set(pdf.PageSizeA4)
	pdfGenerator.Dpi.Set(300)

	pdfGenerator.AddPage(pdf.NewPageReader(strings.NewReader(htmlStr)))

	err = pdfGenerator.Create()
	if err != nil {
		logger.Error("ExportToPdf", err.Error())
		api.Error(res, req, "Error while creating PDF", err, http.StatusInternalServerError)
		return
	}

	err = pdfGenerator.WriteFile("./user.pdf")
	if err != nil {
		logger.Error("ExportToPdf", err.Error())
		api.Error(res, req, "Error while exporting user information to PDF", err, http.StatusInternalServerError)
		return
	}

	fileBytes, err := os.ReadFile("./user.pdf")
	if err != nil {
		logger.Error("ExportToCsv", err.Error())
		api.Error(res, req, "Error while exporting user information to PDF", err, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/pdf")
	res.Header().Set("Content-Disposition", "attachment; filename=user.pdf")

	_, err = res.Write(fileBytes)
	if err != nil {
		logger.Error("ExportToCsv", err.Error())
		api.Error(res, req, "Error while exporting user information to PDF", err, http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
