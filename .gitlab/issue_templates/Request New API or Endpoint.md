<!-- 🚧 Please make sure to add a meaningful issue title above -->

<!-- 🚧 Please change the first heading to either `New API Support` or `New API Endpoint Support` -->

/label ~"group::environments" ~"type::feature" ~"feature::addition"

## New API / Endpoint

<!-- 🚧 Please add the related API documentation link below -->
<!-- 🚧 Please note we do not typically add support for Beta/Experimental APIs -->

API documentation:

## Additional Details

<!-- 🚧 Please tick the boxes which apply: -->

- [ ] I'd like to contribute it myself

## Implementation Guide

The code in `client-go` has a strong pattern that is repeatable when adding support for new APIs.

- Read the instructions in `CONTRIBUTING.md` to get your local development environment set up.
- Follow the instructions in `AddingAPISupport.md`.
  - This file includes instructions for adding all endpoints of an API.
  - It also describes how to write functions for each type of endpoint, so can be used for adding new endpoint support as well.
