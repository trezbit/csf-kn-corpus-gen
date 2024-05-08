# CSF Corpus Builder ChatGPT Client



**Description**:  

Go command-line utility leveraging OpenAI - ChatGPT4 to generate a corpus of: assessment questionnaires, guideline and implementation documentation covering NIST CSF 2.0 and related standards and controls.

The [Generated Corpus](./corpus/README.md) is used in other CSF KB - KNN utilities for building and training models addressing a variety of Cybersecurity compliance use-cases including:

1. Training domain specific models and KBs that would allow mapping between different standards
2. Generating questionnaires for compliance assessment for all applicable standards.
3. To build knowledge-bases, supporting graph-based inferences for tracing continuity and coverage between published CSF versions and mapped standards.

**Status**:  Baseline version 1.0

![**Flows**](./docs/Flows.drawio.png)

## Runtime Env and Dependencies
*Tested on*: Ubuntu 22.04 - *Golang Version*: 1.22.1 

*Requires:* Go version 1.18 or greater

Active [OpenAI API](https://platform.openai.com/) subscription allowing access to GPT 3.5+ (More on creating API keys at: [Go-OpenAI](https://github.com/sashabaranov/go-openai?tab=readme-ov-file#getting-an-openai-api-key) https://github.com/sashabaranov/go-openai?tab=readme-ov-file#getting-an-openai-api-key

## Configuration

###### ChatGPT API Connectivity

The utility expects the following environment variables to be set:

```bash
export OPEN_API_SECRET=<YOUR-OPEN-API-SECRET>
export OPEN_API_ORGID=<YOUR-OPEN-API-ORGID>
```

###### Corpus Generation Settings

Sample configuration is available at: `./config/.env`

```config
# Prompt generation per control
INPATH_QUERY_ASSESS=model/controls/csf-core-v2.0.controls.txt
INPATH_QUERY_CSF=model/maps/csf-core-v2-map.txt
INPATH_QUERY_MAP=model/maps/map.csf-v2.0.txt
# Corpus Out locations
OUTPATH_ASSESS=corpus/assess
OUTPATH_CSF=corpus/csf-core
OUTPATH_MAP=corpus/map-v2.0
# Go routine Workgroup settings to work around ChatGPT API access limitations
API_BATCH_SIZE=10
API_WAIT_SECONDS=60
```

## Usage

Having all corpus generation options are captured in the `./config/.env` as explained above:

+ Generate assessment queries for Standard Controls:        `go run main.go -qtype=assess`
+ Generate general guidance and implementation documentation for CSF  (v 1.1 and 2.0): `go run main.go -qtype=csf`
+ Generate general guidance and implementation documentation for explicitly mapped standards:  `go run main.go -qtype=map`

## Testing

Rudimentary testing provided with Golang `go test`

## Features to Implement & Known issues

- Extending configuration to specify a GPT version
- Dockerized version for the CLI

## Getting help

If you have questions, concerns, bug reports, etc, please file an issue in this repository's Issue Tracker.

## Open source licensing info
1. [LICENSE](LICENSE)

----
## Credits and references

#### [Go OpenAI](https://github.com/sashabaranov/go-openai)

[![Go Reference](https://camo.githubusercontent.com/d8dac80151ec3132e6c45fbfd75d6b3ac13e7440c810a36f872b1613209a953b/68747470733a2f2f706b672e676f2e6465762f62616467652f6769746875622e636f6d2f7361736861626172616e6f762f676f2d6f70656e61692e737667)](https://pkg.go.dev/github.com/sashabaranov/go-openai) [![Go Report Card](https://camo.githubusercontent.com/500b6977e8288774a51ab9d78c84ed9f60cbf63bbcb8fe60f9acd21a50451304/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f7361736861626172616e6f762f676f2d6f70656e6169)](https://goreportcard.com/report/github.com/sashabaranov/go-openai) [![codecov](https://camo.githubusercontent.com/62ae8603d7891a6e63dc56759d5145120bf187ba09ebb9583b22fc59fc9fdff5/68747470733a2f2f636f6465636f762e696f2f67682f7361736861626172616e6f762f676f2d6f70656e61692f6272616e63682f6d61737465722f67726170682f62616467652e7376673f746f6b656e3d6243624966484c497357)](https://codecov.io/gh/sashabaranov/go-openai)

Unofficial Go clients library for [OpenAI API](https://platform.openai.com/). 

#### [CSF Tools](https://csf.tools/) 

Exploration and Visualization Tools by NIST Cybersecurity Framework (CSF) and Privacy Framework (PF)

