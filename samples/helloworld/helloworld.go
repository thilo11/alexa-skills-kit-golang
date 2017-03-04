package main

//TODO Need to be main?  Breaks as helloworld

import "C"

import (
	"errors"
	"log"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	alexa "github.com/ericdaugherty/alexa-skills-kit-golang"
)

var a = &alexa.Alexa{ApplicationID: "amzn1.ask.skill.08857461-c080-49d6-8646-2f0ca2c14914", RequestHandler: &HelloWorld{}, IgnoreTimestamp: true}

const cardTitle = "HelloWorld"

// HelloWorld handles reqeusts from the HelloWorld skill.
type HelloWorld struct{}

// Handle processes calls from Lambda
func Handle(requestEnv *alexa.RequestEnvelope, ctx *runtime.Context) (interface{}, error) {
	return a.ProcessRequest(requestEnv)
}

// OnSessionStarted called when a new session is created.
func (h *HelloWorld) OnSessionStarted(request *alexa.Request, session *alexa.Session, response *alexa.Response) error {

	log.Printf("OnSessionStarted requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	return nil
}

// OnLaunch called with a reqeust is received of type LaunchRequest
func (h *HelloWorld) OnLaunch(request *alexa.Request, session *alexa.Session, response *alexa.Response) error {
	speechText := "Welcome to the Alexa Skills Kit, you can say hello"

	log.Printf("OnLaunch requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	response.SetSimpleCard(cardTitle, speechText)
	response.SetOutputText(speechText)
	response.SetRepromptText(speechText)

	response.ShouldSessionEnd = true

	return nil
}

// OnIntent called with a reqeust is received of type IntentRequest
func (h *HelloWorld) OnIntent(request *alexa.Request, session *alexa.Session, response *alexa.Response) error {

	log.Printf("OnIntent requestId=%s, sessionId=%s, intent=%s", request.RequestID, session.SessionID, request.Intent.Name)

	switch request.Intent.Name {
	case "HelloWorldIntent":
		log.Println("HelloWorldIntent triggered")
		speechText := "Hello World"

		response.SetSimpleCard(cardTitle, speechText)
		response.SetOutputText(speechText)

		log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
	case "AMAZON.HelpIntent":
		log.Println("AMAZON.HelpIntent triggered")
		speechText := "You can say hello to me!"

		response.SetSimpleCard("HelloWorld", speechText)
		response.SetOutputText(speechText)
		response.SetRepromptText(speechText)
	default:
		return errors.New("Invalid Intent")
	}

	return nil
}

// OnSessionEnded called with a reqeust is received of type SessionEndedRequest
func (h *HelloWorld) OnSessionEnded(request *alexa.Request, session *alexa.Session, response *alexa.Response) error {

	log.Printf("OnSessionEnded requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	return nil
}