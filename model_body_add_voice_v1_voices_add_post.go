/*
 * ElevenLabs API Documentation
 *
 * This is the documentation for the ElevenLabs API. You can use this API to use our service programmatically, this is done by using your xi-api-key. <br/> You can view your xi-api-key using the 'Profile' tab on https://beta.elevenlabs.io. Our API is experimental so all endpoints are subject to change.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BodyAddVoiceV1VoicesAddPost struct {
	// The name that identifies this voice. This will be displayed in the dropdown of the website.
	Name string `json:"name"`
	// One or more audio files to clone the voice from
	Files []*os.File `json:"files"`
	// How would you describe the voice?
	Description string `json:"description,omitempty"`
	// Serialized labels dictionary for the voice.
	Labels string `json:"labels,omitempty"`
}
