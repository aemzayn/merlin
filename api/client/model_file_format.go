/*
 * Merlin
 *
 * API Guide for accessing Merlin's model deployment functionalities
 *
 * API version: 0.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type FileFormat string

// List of FileFormat
const (
	INVALID_FILE_FORMAT_FileFormat FileFormat = "INVALID_FILE_FORMAT"
	CSV_FileFormat                 FileFormat = "CSV"
	PARQUET_FileFormat             FileFormat = "PARQUET"
	AVRO_FileFormat                FileFormat = "AVRO"
	JSON_FileFormat                FileFormat = "JSON"
)
