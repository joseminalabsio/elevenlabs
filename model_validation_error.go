/*
 * ElevenLabs API Documentation
 *
 * This is the documentation for the ElevenLabs API. You can use this API to use our service programmatically, this is done by using your xi-api-key. <br/> You can view your xi-api-key using the 'Profile' tab on https://beta.elevenlabs.io. Our API is experimental so all endpoints are subject to change.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ValidationError struct {
	Loc []AnyOfValidationErrorLocItems `json:"loc"`
	Msg string `json:"msg"`
	Type_ string `json:"type"`
}
