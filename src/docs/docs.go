// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/dashu-baba/cricket-scoreboard-api",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/dashu-baba/cricket-scoreboard-api/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/series": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Create a series item",
                "parameters": [
                    {
                        "description": "Create Series",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.SeriesCreateModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.Series"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/series/:id": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Get singe item of series",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responsemodels.Team"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Update series status",
                "parameters": [
                    {
                        "description": "Update Series Status Model",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.UpdateSeriesStatusModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/series/:id/matches": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Create list of matches under a series",
                "parameters": [
                    {
                        "description": "Create Matches",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.MatchCreateModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/series/:id/matches/:matchid": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Update starting players of the match",
                "parameters": [
                    {
                        "description": "Match Playing Squad",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.MatchPlayingSquadModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Match ID",
                        "name": "matchid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Update match status",
                "parameters": [
                    {
                        "description": "Update Match Status Model",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.UpdateSeriesStatusModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Match ID",
                        "name": "matchid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/series/:id/matches/:matchid/innings": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Created an innings",
                "parameters": [
                    {
                        "description": "Create Innings Model",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.CreateInningsModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Match ID",
                        "name": "matchid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/series/:id/matches/:matchid/innings/:inningsid/start": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Created an innings",
                "parameters": [
                    {
                        "description": "Start Innings Model",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.StartInningsModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Match ID",
                        "name": "matchid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Innings ID",
                        "name": "inningsid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/series/:id/teams": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Add list of teams to the series item",
                "parameters": [
                    {
                        "description": "Add Teams",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.TeamsAddModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.Series"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Remove list of teams to the series item",
                "parameters": [
                    {
                        "description": "Remove Teams",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.TeamsRemoveModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/series/:id/teams/:teamid": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "description": "Update Squad model",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.UpdateSquadModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Team ID",
                        "name": "teamid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/teams": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Get list of teams",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responsemodels.Team"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Create a team item",
                "parameters": [
                    {
                        "description": "Create Team",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.TeamCreateModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.Team"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/teams/:id": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Get singe item of team",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.Team"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Update a team item",
                "parameters": [
                    {
                        "description": "Update Team",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.TeamUpdateModel"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Team ID",
                        "name": "string",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/teams/:id/players": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Add a player to the team item",
                "parameters": [
                    {
                        "description": "Add Team",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.PlayerCreateModel"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Team ID",
                        "name": "string",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.Player"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            }
        },
        "/teams/:id/players/:playerid": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Update a player item",
                "parameters": [
                    {
                        "description": "Update Team",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodels.PlayerUpdateModel"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Team ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Player ID",
                        "name": "playerid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsemodels.ErrorModel"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Remove a player from the team item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Player ID",
                        "name": "playerid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {}
                }
            }
        }
    },
    "definitions": {
        "gin.H": {
            "type": "object",
            "additionalProperties": true
        },
        "requestmodels.CreateInningsModel": {
            "type": "object",
            "required": [
                "battingTeamId",
                "bowlingTeamId",
                "tossWinningTeamId"
            ],
            "properties": {
                "battingTeamId": {
                    "type": "string"
                },
                "bowlingTeamId": {
                    "type": "string"
                },
                "tossWinningTeamId": {
                    "type": "string"
                }
            }
        },
        "requestmodels.Match": {
            "type": "object",
            "required": [
                "matchType",
                "overLimit",
                "participants"
            ],
            "properties": {
                "matchType": {
                    "type": "integer"
                },
                "overLimit": {
                    "type": "integer"
                },
                "participants": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "requestmodels.MatchCreateModel": {
            "type": "object",
            "required": [
                "matches"
            ],
            "properties": {
                "matches": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/requestmodels.Match"
                    }
                }
            }
        },
        "requestmodels.MatchPlayingSquadModel": {
            "type": "object",
            "required": [
                "teamId"
            ],
            "properties": {
                "extras": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "players": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "teamId": {
                    "type": "string"
                }
            }
        },
        "requestmodels.PlayerCreateModel": {
            "type": "object",
            "required": [
                "name",
                "playerType"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "playerType": {
                    "type": "integer"
                }
            }
        },
        "requestmodels.PlayerModel": {
            "type": "object",
            "required": [
                "name",
                "playerType"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "playerType": {
                    "type": "integer"
                }
            }
        },
        "requestmodels.PlayerUpdateModel": {
            "type": "object",
            "required": [
                "name",
                "playerType"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "playerType": {
                    "type": "integer"
                }
            }
        },
        "requestmodels.SeriesCreateModel": {
            "type": "object",
            "required": [
                "gameType",
                "name",
                "teams"
            ],
            "properties": {
                "gameType": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "teams": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/requestmodels.SeriesParticipantModel"
                    }
                }
            }
        },
        "requestmodels.SeriesParticipantModel": {
            "type": "object",
            "required": [
                "squadPlayers",
                "teamId"
            ],
            "properties": {
                "squadPlayers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "teamId": {
                    "type": "string"
                }
            }
        },
        "requestmodels.StartInningsModel": {
            "type": "object",
            "required": [
                "bowlerID",
                "nonStrikeBatsmanId",
                "strikeBatsmanId"
            ],
            "properties": {
                "bowlerID": {
                    "type": "string"
                },
                "nonStrikeBatsmanId": {
                    "type": "string"
                },
                "strikeBatsmanId": {
                    "type": "string"
                }
            }
        },
        "requestmodels.TeamCreateModel": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "logo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/requestmodels.PlayerModel"
                    }
                }
            }
        },
        "requestmodels.TeamUpdateModel": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "requestmodels.TeamsAddModel": {
            "type": "object",
            "required": [
                "teams"
            ],
            "properties": {
                "teams": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/requestmodels.SeriesParticipantModel"
                    }
                }
            }
        },
        "requestmodels.TeamsRemoveModel": {
            "type": "object",
            "required": [
                "teams"
            ],
            "properties": {
                "teams": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "requestmodels.UpdateSeriesStatusModel": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "type": "integer"
                }
            }
        },
        "requestmodels.UpdateSquadModel": {
            "type": "object",
            "properties": {
                "addedPlayer": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "removedPlayer": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "responsemodels.ErrorModel": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "responsemodels.Player": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "playerType": {
                    "type": "integer"
                },
                "teamID": {
                    "type": "string"
                }
            }
        },
        "responsemodels.Series": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "teams": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responsemodels.Team"
                    }
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "responsemodels.Team": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responsemodels.Player"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Cricket Scoreboard API",
	Description: "This contains the commn REST api to support a modern cricket scoreboard project.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
