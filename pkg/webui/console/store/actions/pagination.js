// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// utils
export const createGetBaseActionType = name => (
  `GET_${name}_LIST`
)

// pagination
export const createGetPaginationActionType = name => (
  `${createGetBaseActionType(name)}_REQUEST`
)
export const createGetPaginationSuccessActionType = name => (
  `${createGetBaseActionType(name)}_SUCCESS`
)
export const createGetPaginationFailureActionType = name => (
  `${createGetBaseActionType(name)}_FAILURE`
)

export const getPagination = name => (params, entityId) => (
  { type: createGetPaginationActionType(name), params, entityId }
)
export const getPaginationSuccess = name => (entities, totalCount, entityId) => (
  { type: createGetPaginationSuccessActionType(name), entities, totalCount, entityId }
)
export const getPaginationFailure = name => error => (
  { type: createGetPaginationFailureActionType(name), error }
)