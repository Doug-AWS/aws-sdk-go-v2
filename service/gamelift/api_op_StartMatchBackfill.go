// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type StartMatchBackfillInput struct {
	_ struct{} `type:"structure"`

	// Name of the matchmaker to use for this request. The name of the matchmaker
	// that was used with the original game session is listed in the GameSession
	// object, MatchmakerData property. This property contains a matchmaking configuration
	// ARN value, which includes the matchmaker name. (In the ARN value "arn:aws:gamelift:us-west-2:111122223333:matchmakingconfiguration/MM-4v4",
	// the matchmaking configuration name is "MM-4v4".) Use only the name for this
	// parameter.
	//
	// ConfigurationName is a required field
	ConfigurationName *string `type:"string" required:"true"`

	// Amazon Resource Name (ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html))
	// that is assigned to a game session and uniquely identifies it.
	//
	// GameSessionArn is a required field
	GameSessionArn *string `min:"1" type:"string" required:"true"`

	// Match information on all players that are currently assigned to the game
	// session. This information is used by the matchmaker to find new players and
	// add them to the existing game.
	//
	//    * PlayerID, PlayerAttributes, Team -\\- This information is maintained
	//    in the GameSession object, MatchmakerData property, for all players who
	//    are currently assigned to the game session. The matchmaker data is in
	//    JSON syntax, formatted as a string. For more details, see Match Data (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-server.html#match-server-data).
	//
	//    * LatencyInMs -\\- If the matchmaker uses player latency, include a latency
	//    value, in milliseconds, for the region that the game session is currently
	//    in. Do not include latency values for any other region.
	//
	// Players is a required field
	Players []Player `type:"list" required:"true"`

	// Unique identifier for a matchmaking ticket. If no ticket ID is specified
	// here, Amazon GameLift will generate one in the form of a UUID. Use this identifier
	// to track the match backfill ticket status and retrieve match results.
	TicketId *string `type:"string"`
}

// String returns the string representation
func (s StartMatchBackfillInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartMatchBackfillInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartMatchBackfillInput"}

	if s.ConfigurationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationName"))
	}

	if s.GameSessionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameSessionArn"))
	}
	if s.GameSessionArn != nil && len(*s.GameSessionArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionArn", 1))
	}

	if s.Players == nil {
		invalidParams.Add(aws.NewErrParamRequired("Players"))
	}
	if s.Players != nil {
		for i, v := range s.Players {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Players", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type StartMatchBackfillOutput struct {
	_ struct{} `type:"structure"`

	// Ticket representing the backfill matchmaking request. This object includes
	// the information in the request, ticket status, and match results as generated
	// during the matchmaking process.
	MatchmakingTicket *MatchmakingTicket `type:"structure"`
}

// String returns the string representation
func (s StartMatchBackfillOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartMatchBackfill = "StartMatchBackfill"

// StartMatchBackfillRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Finds new players to fill open slots in an existing game session. This operation
// can be used to add players to matched games that start with fewer than the
// maximum number of players or to replace players when they drop out. By backfilling
// with the same matchmaker used to create the original match, you ensure that
// new players meet the match criteria and maintain a consistent experience
// throughout the game session. You can backfill a match anytime after a game
// session has been created.
//
// To request a match backfill, specify a unique ticket ID, the existing game
// session's ARN, a matchmaking configuration, and a set of data that describes
// all current players in the game session. If successful, a match backfill
// ticket is created and returned with status set to QUEUED. The ticket is placed
// in the matchmaker's ticket pool and processed. Track the status of the ticket
// to respond as needed.
//
// The process of finding backfill matches is essentially identical to the initial
// matchmaking process. The matchmaker searches the pool and groups tickets
// together to form potential matches, allowing only one backfill ticket per
// potential match. Once the a match is formed, the matchmaker creates player
// sessions for the new players. All tickets in the match are updated with the
// game session's connection information, and the GameSession object is updated
// to include matchmaker data on the new players. For more detail on how match
// backfill requests are processed, see How Amazon GameLift FlexMatch Works
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-match.html).
//
// Learn more
//
//  Backfill Existing Games with FlexMatch (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-backfill.html)
//
//  How GameLift FlexMatch Works (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-match.html)
//
// Related operations
//
//    * StartMatchmaking
//
//    * DescribeMatchmaking
//
//    * StopMatchmaking
//
//    * AcceptMatch
//
//    * StartMatchBackfill
//
//    // Example sending a request using StartMatchBackfillRequest.
//    req := client.StartMatchBackfillRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/StartMatchBackfill
func (c *Client) StartMatchBackfillRequest(input *StartMatchBackfillInput) StartMatchBackfillRequest {
	op := &aws.Operation{
		Name:       opStartMatchBackfill,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartMatchBackfillInput{}
	}

	req := c.newRequest(op, input, &StartMatchBackfillOutput{})
	return StartMatchBackfillRequest{Request: req, Input: input, Copy: c.StartMatchBackfillRequest}
}

// StartMatchBackfillRequest is the request type for the
// StartMatchBackfill API operation.
type StartMatchBackfillRequest struct {
	*aws.Request
	Input *StartMatchBackfillInput
	Copy  func(*StartMatchBackfillInput) StartMatchBackfillRequest
}

// Send marshals and sends the StartMatchBackfill API request.
func (r StartMatchBackfillRequest) Send(ctx context.Context) (*StartMatchBackfillResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartMatchBackfillResponse{
		StartMatchBackfillOutput: r.Request.Data.(*StartMatchBackfillOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartMatchBackfillResponse is the response type for the
// StartMatchBackfill API operation.
type StartMatchBackfillResponse struct {
	*StartMatchBackfillOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartMatchBackfill request.
func (r *StartMatchBackfillResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
