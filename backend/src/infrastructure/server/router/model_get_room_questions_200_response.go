/*
 * refactoring-together
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package router

type GetRoomQuestions200Response struct {

	RoomId *interface{} `json:"roomId,omitempty"`

	Questions []Question `json:"questions,omitempty"`
}