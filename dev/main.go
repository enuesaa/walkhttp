package main

import (
	"os"

	"github.com/xitongsys/parquet-go/writer"
)

type HTTPResponse struct {
	URL        string `parquet:"name=url, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	Method     string `parquet:"name=method, type=BYTE_ARRAY, encoding=PLAIN_DICTIONARY"`
	StatusCode int    `parquet:"name=status_code, type=INT32"`
	Headers    string `parquet:"name=headers, type=BYTE_ARRAY"`
	Body       string `parquet:"name=body, type=BYTE_ARRAY"`
}

func main() {
	response := HTTPResponse{
		URL:        "https://example.com",
		Method:     "GET",
		StatusCode: 200,
		Headers:    "",
		Body:       "",
	}

	f, err := os.Create("data.parquet")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pw, err := writer.NewParquetWriterFromWriter(f, new(HTTPResponse), 4)
	if err != nil {
		panic(err)
	}

	if err = pw.Write(response); err != nil {
		panic(err)
	}
	if err = pw.WriteStop(); err != nil {
		panic(err)
	}
}
