package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleLambdaEvent)
}

type Event struct {
	HTML64     string `json:"html64"`
	PDFOptions struct {
		IsLandscape    bool `json:"is_landscape"`
		NeedPagination bool `json:"need_pagination"`
	} `json:"pdf_options"`
}

type Response struct {
	PDF64 string `json:"pdf"`
}

func HandleLambdaEvent(event Event) (*Response, error) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, fmt.Errorf("failed to get PDFGenerator: %w", err)
	}

	// Set global options
	pdfg.Dpi.Set(600)
	if event.PDFOptions.IsLandscape { // Portrait is default option
		pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	}

	// Create a new input Page from a HTML64 string
	htmlBytes, err := base64.StdEncoding.DecodeString(event.HTML64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode html64 string: %w", err)
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlBytes))

	if event.PDFOptions.NeedPagination {
		page.FooterRight.Set("[page]")
	}

	pdfg.AddPage(page)

	// Creating PDF document
	var writer bytes.Buffer

	pdfg.SetOutput(&writer)

	err = pdfg.Create()
	if err != nil {
		return nil, fmt.Errorf("failed to create the PDF document: %w", err)
	}

	pdf64 := base64.StdEncoding.EncodeToString(writer.Bytes())

	return &Response{PDF64: pdf64}, nil
}
